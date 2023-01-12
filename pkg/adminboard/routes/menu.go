package routes

import (
	"github.com/adminboard/adminboard/pkg/adminboard/admin"
	"github.com/eqto/api-server"
)

func Menu(ctx api.Context) error {
	groupID := ctx.Session().GetInt(`groupID`)
	if groupID == 0 {
		return ctx.StatusForbidden(`invalid group id`)
	}

	menus, e := admin.GetMenu(true, groupID)
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	return ctx.Write(menus)
}
