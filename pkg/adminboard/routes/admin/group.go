package admin

import (
	"strings"

	"github.com/adminboard/server/pkg/adminboard/db"
	"github.com/adminboard/server/pkg/adminboard/routes/params"
	"github.com/eqto/api-server"
	"github.com/eqto/dbm"
	"github.com/eqto/go-json"
)

func Groups(ctx api.Context) error {
	stmt := dbm.Select(`*`).From(db.Prefix(`group`))

	cn, e := ctx.Database()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	rs, e := cn.Select(cn.SQL(stmt))
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	return ctx.Write(rs)
}

func GroupsAdd(ctx api.Context) error {
	cn, e := ctx.Database()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	stmtInsert := dbm.InsertInto(db.Prefix(`group`), `name, status`)
	js := ctx.Request().JSON()
	name := strings.TrimSpace(js.GetString(`name`))
	if name == `` {
		return ctx.StatusBadRequest(`invalid parameter:name`)
	}

	_, e = cn.Exec(cn.SQL(stmtInsert), name, `enabled`)
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	return nil
}

func GroupsDetails(ctx api.Context) error {
	params, e := params.FetchQuery(ctx, `id`, ``)
	if e != nil {
		return ctx.StatusBadRequest(e.Error())
	}
	cn, e := ctx.Database()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}

	groupID := params.GetString(`id`)
	stmt := dbm.Select(`id, name, status`).From(db.Prefix(`group`)).Where(`id = ?`)
	rs, e := cn.Get(cn.SQL(stmt), groupID)
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	resp := json.Object{
		`id`:     rs.Int(`id`),
		`name`:   rs.String(`name`),
		`status`: rs.String(`status`),
	}
	stmtPermission := dbm.Select(`p.id, p.name, p.description, gp.group_id`).From(db.Prefix(`permission p`)).
		LeftJoin(db.Prefix(`group_permission gp`), `gp.permission_id = p.id AND gp.group_id = ?`).OrderBy(`p.name`)

	rsPermission, e := cn.Select(cn.SQL(stmtPermission), groupID)
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	resp.Put(`permissions`, rsPermission)

	return ctx.Write(resp)
}

func GroupsAddPermissions(ctx api.Context) error {
	js, e := params.Fetch(ctx, `group_id, permission_id`, ``)
	if e != nil {
		return ctx.StatusBadRequest(e.Error())
	}

	cn, e := ctx.Database()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}

	stmt := dbm.InsertInto(db.Prefix(`group_permission`), `group_id, permission_id`).Values(`?, ?`)
	_, e = cn.Exec(cn.SQL(stmt), js.GetInt(`group_id`), js.GetInt(`permission_id`))
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}

	return nil
}

func GroupsRemovePermissions(ctx api.Context) error {
	js, e := params.Fetch(ctx, `group_id, permission_id`, ``)
	if e != nil {
		return ctx.StatusBadRequest(e.Error())
	}

	cn, e := ctx.Database()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}

	stmt := dbm.DeleteFrom(db.Prefix(`group_permission`)).Where(`group_id = ? AND permission_id = ?`)
	_, e = cn.Exec(cn.SQL(stmt), js.GetInt(`group_id`), js.GetInt(`permission_id`))
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}

	return nil
}
