package admin

import (
	"github.com/adminboard/adminboard/pkg/adminboard/db"
	"github.com/eqto/api-server"
	"github.com/eqto/dbm"
)

func Pages(ctx api.Context) error {
	cn := db.CN()

	stmt := dbm.Select(`name, url, title, description, status`).From(db.Prefix(`page`))

	rs, e := cn.Select(cn.SQL(stmt))
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	return ctx.Write(rs)
}

func PageAdd(ctx api.Context) error {
	cn, e := ctx.Database()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	js := ctx.Request().JSON()

	stmt := dbm.InsertInto(db.Prefix(`page`), `name, url, title, description, status`).Values(`?, ?, ?, ?, 'draft'`)

	_, e = cn.Exec(cn.SQL(stmt), js.GetString(`name`), js.GetString(`url`), js.GetString(`title`), js.GetString(`description`))
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}

	return nil
}
