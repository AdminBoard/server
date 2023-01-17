package routes

import (
	"github.com/adminboard/adminboard/pkg/adminboard/db"
	"github.com/eqto/api-server"
	"github.com/eqto/dbm"
	"github.com/eqto/dbm/stmt"
)

func Page(ctx api.Context) error {
	groupID := ctx.Session().GetInt(`groupID`)

	var stmt *stmt.Select
	params := []any{}
	params = append(params, groupID)

	if name := ctx.Request().QueryParam(`name`); name != `` {
		stmt = dbm.Select(`p.title, p.description, p.layout, gp.group_id`).From(db.Prefix(`page p`)).
			LeftJoin(db.Prefix(`group_page gp`), `p.id = gp.page_id AND gp.group_id = ?`).Where(`p.name = ?`).Limit(1)
		params = append(params, name)
	} else if url := ctx.Request().QueryParam(`url`); url != `` {
		stmt = dbm.Select(`p.title, p.description, p.layout, gp.group_id`).From(db.Prefix(`page p`)).
			LeftJoin(db.Prefix(`group_page gp`), `p.id = gp.page_id AND gp.group_id = ?`).Where(`p.url = ?`).Limit(1)
		params = append(params, url)
	} else {
		return ctx.StatusBadRequest(`invalid request, need parameter of name or url`)
	}

	cn, e := ctx.Database()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	rs, e := cn.Get(cn.SQL(stmt), params...)
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	if rs.Int(`group_id`) == 0 {
		return ctx.StatusForbidden(`access denied`)
	}

	if rs == nil {
		return ctx.StatusNotFound(`page not found`)
	}

	return ctx.Write(rs)
}
