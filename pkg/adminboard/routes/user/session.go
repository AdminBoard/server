package user

import (
	"github.com/adminboard/server/pkg/adminboard/session"
	"github.com/eqto/api-server"
)

func Session(ctx api.Context) error {
	token, e := session.Get(ctx)
	if e != nil {
		return ctx.StatusUnauthorized(e.Error())
	} else if token.Expired == 1 {
		return ctx.Error(1, `password expired`)
	}
	return nil
}
