package middleware

import (
	"strings"

	"github.com/adminboard/server/pkg/adminboard/db"
	"github.com/adminboard/server/pkg/adminboard/session"
	"github.com/eqto/api-server"
	"github.com/eqto/dbm"
)

func AuthMiddleware(ctx api.Context) error {
	token, e := session.Get(ctx)
	if e != nil {
		return ctx.StatusUnauthorized(e.Error())
	}

	cn, e := ctx.Database()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}

	stmt := dbm.Select(`us.user_id, u.group_id`).
		From(db.Prefix(`user_session us`)).
		InnerJoin(db.Prefix(`user u`), `us.user_id = u.id`).
		Where(`us.id = ?`)

	rs, e := cn.Get(cn.SQL(stmt), token.SessionID)
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	if rs == nil {
		return ctx.StatusUnauthorized(`session not found or user have no group`)
	}
	return nil
}

func AuthAPI(ctx api.Context) error {
	token, e := session.Get(ctx)
	if e != nil {
		return ctx.StatusUnauthorized(e.Error())
	}
	if token.GroupID == 1 {
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

	stmt := dbm.Select(`id, status`).From(db.Prefix(`api a`)).
		InnerJoin(db.Prefix(`permission_item pi`), `a.id = pi.ref_id AND pi.ref_type = 'api'`).
		InnerJoin(db.Prefix(`group_permission gp`), `pi.permission_id = gp.permission_id`).
		Where(`a.path = ?`, `gp.group_id = ?`)

	rs, e := cn.Get(cn.SQL(stmt), path, token.GroupID)
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	if rs == nil {
		return ctx.StatusForbidden(`you don't have permission to this resource`)
	} else if rs.String(`status`) != `active` {
		return ctx.StatusBadRequest(`invalid api status`)
	}

	return nil

}
