package middleware

import (
	"strings"
	"time"

	"github.com/adminboard/adminboard/pkg/adminboard/db"
	"github.com/adminboard/adminboard/pkg/adminboard/routes"
	"github.com/eqto/api-server"
	"github.com/eqto/dbm"
)

func AuthMiddleware(ctx api.Context) error {
	sessionID, e := routes.GetSessionID(ctx)
	if e != nil {
		return e
	}

	rs, e := db.CN().Get(db.QueryWithPrefix(`SELECT us.user_id, ug.group_id FROM {prefix}user_session us 
		INNER JOIN {prefix}user_group ug ON us.user_id = ug.user_id WHERE us.id = ?`), sessionID)
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	if rs == nil {
		return ctx.StatusUnauthorized(`session not found or user have no group`)
	}

	ctx.Session().Put(`userID`, rs.Int(`user_id`))
	ctx.Session().Put(`groupID`, rs.Int(`group_id`))
	ctx.Session().Put(`sessionID`, sessionID)

	session := ctx.Request().Header().Cookie(`SESS`)
	ctx.Response().Header().SetCookie(`SESS`, session, 15*time.Minute)
	return nil
}

func AuthAPI(ctx api.Context) error {
	groupID := ctx.Session().GetInt(`groupID`)
	if groupID == 0 {
		return nil
	}
	path := ctx.URL().Path
	if strings.HasPrefix(path, `/api`) {
		path = path[4:]
	}
	cn, e := ctx.Database()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}

	stmt := dbm.Select(`*`).From(db.Prefix(`api a`)).InnerJoin(db.Prefix(`group_api ga`), `a.id = ga.api_id`).Where(`a.url = ?`, `ga.group_id = ?`)

	rs, e := cn.Select(cn.SQL(stmt), path, groupID)

	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	if rs == nil {
		return ctx.StatusForbidden(`Forbidden: you don't have permission to this resource`)
	}

	return nil

}
