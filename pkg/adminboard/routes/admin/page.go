package admin

import (
	"github.com/adminboard/adminboard/pkg/adminboard/db"
	"github.com/eqto/api-server"
	"github.com/eqto/dbm"
	"github.com/eqto/dbm/stmt"
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

func PageUpdate(ctx api.Context) error {
	pageName := ctx.URL().Query().Get(`name`)
	if pageName == `` {
		return ctx.StatusBadRequest(`invalid page name`)
	}

	s := dbm.Update(db.Prefix(`page`))
	var stmtSet *stmt.UpdateFields
	params := []interface{}{}

	js := ctx.Request().JSON()

	for key := range js {
		stmtSet = s.Set(key + ` = ?`)
		params = append(params, js.GetString(key))
	}

	if len(params) == 0 {
		return ctx.StatusBadRequest(`no update`)
	}
	stmtSet.Where(`name = ?`)
	params = append(params, pageName)

	cn, e := ctx.Database()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	_, e = cn.Exec(cn.SQL(s), params...)
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}

	return nil
}
