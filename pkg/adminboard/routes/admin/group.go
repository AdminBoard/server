package admin

import (
	"strings"

	"github.com/adminboard/adminboard/pkg/adminboard/db"
	"github.com/eqto/api-server"
	"github.com/eqto/dbm"
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
