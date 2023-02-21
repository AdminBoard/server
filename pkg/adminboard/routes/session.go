package routes

import (
	"github.com/adminboard/adminboard/pkg/adminboard/session"
	"github.com/eqto/api-server"
)

func Session(ctx api.Context) error {
	_, e := session.Get(ctx)
	if e != nil {
		return ctx.StatusUnauthorized(e.Error())
	}
	return nil
}
