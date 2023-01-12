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
