package routes

import (
	"github.com/adminboard/adminboard/pkg/adminboard/session"
	"github.com/eqto/api-server"
	"github.com/eqto/go-json"
)

func Session(ctx api.Context) error {
	token, e := session.Get(ctx)
	if e != nil {
		return ctx.StatusUnauthorized(e.Error())
	} else if token.Expired == 1 {
		return ctx.Write(json.Object{`status`: 1, `message`: `Password expired`})
	}
	return nil
}
