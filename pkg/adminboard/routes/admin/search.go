package admin

import (
	"strings"

	"github.com/adminboard/adminboard/pkg/adminboard/db"
	"github.com/eqto/api-server"
	"github.com/eqto/dbm"
)

func SearchGroups(ctx api.Context) error {
	name := ctx.Request().QueryParam(`name`)

	name = strings.TrimSpace(name)

	if name == `` {
		return ctx.StatusBadRequest(`invalid parameter: name`)
	}

	cn, e := ctx.Database()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	name += `*`
	stmt := dbm.Select(`id, name as label`).From(db.Prefix(`group`)).Where(`status = 'enabled'`, `MATCH(name) AGAINST(? IN BOOLEAN MODE)`).OrderBy(`name`)

	rs, e := cn.Select(cn.SQL(stmt), name)
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}

	return ctx.Write(rs)
}
