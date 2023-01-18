package admin

import (
	"github.com/adminboard/adminboard/pkg/adminboard/db"
	"github.com/eqto/api-server"
	"github.com/eqto/dbm"
)

func Apis(ctx api.Context) error {
	stmt := dbm.Select(`*`).From(db.Prefix(`api`)).Where(`status IS NOT NULL`).OrderBy(`url`)

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

func ApiAdd(ctx api.Context) error {
	cn, e := ctx.Database()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	js := ctx.Request().JSON()
	res, e := cn.Insert(db.Prefix(`api`), map[string]any{
		`method`:      js.GetStringOr(`method`, `GET`),
		`secure`:      js.GetIntOr(`secure`, 0),
		`url`:         js.GetString(`url`),
		`description`: js.GetString(`description`),
		`status`:      `draft`,
	})
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	row, e := res.RowsAffected()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	if row == 0 {
		return ctx.StatusBadRequest(`operation failed`)
	}
	return nil
}
