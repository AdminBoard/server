package admin

import (
	"github.com/adminboard/adminboard/pkg/adminboard/db"
	"github.com/eqto/api-server"
	"github.com/eqto/dbm"
)

func Users(ctx api.Context) error {
	stmt := dbm.Select(`*`).From(db.Prefix(`user`))

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
