package routes

import (
	"github.com/adminboard/adminboard/pkg/adminboard/admin"
	"github.com/eqto/api-server"
)

func Menu(ctx api.Context) error {
	menus, e := admin.GetMenu(true, ctx.Session().GetInt(`groupID`))
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	return ctx.Write(menus)
}
