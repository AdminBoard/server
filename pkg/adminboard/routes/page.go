package routes

import (
	"github.com/adminboard/server/pkg/adminboard/db"
	"github.com/adminboard/server/pkg/adminboard/routes/params"
	"github.com/adminboard/server/pkg/adminboard/session"
	"github.com/eqto/api-server"
	"github.com/eqto/dbm"
)

func Page(ctx api.Context) error {
	js, e := params.FetchQuery(ctx, `path`, ``)
	if e != nil {
		return ctx.StatusBadRequest(e.Error())
	}
	params := []any{}

	token, _ := session.Get(ctx)
	if !token.IsValid {
		return ctx.StatusBadRequest(`invalid token`)
	}
	stmt := dbm.Select(`p.name, p.description, p.layout`).From(db.Prefix(`page p`))
	path := js.GetString(`path`)
	if token.GroupID != 1 {
		stmt.InnerJoin(db.Prefix(`permission_item pi`), `p.id = pi.ref_id AND pi.ref_type = 'page'`).
			InnerJoin(db.Prefix(`group_permission gp`), `pi.permission_id = gp.permission_id`).
			Where(`gp.group_id = ?`, `p.path = ?`)
		params = append(params, token.GroupID, path)
	} else {
		stmt.Where(`p.path = ?`)
		params = append(params, path)
	}
	stmt.Limit(1)

	cn, e := ctx.Database()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}

	rs, e := cn.Get(cn.SQL(stmt), params...)
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}

	if rs == nil {
		return ctx.StatusNotFound(`page not found`)
	}

	return ctx.Write(rs)
}
