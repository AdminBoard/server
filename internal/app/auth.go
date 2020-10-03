package app

import (
	"errors"

	"github.com/eqto/api-server"
)

func authMiddleware(ctx api.RequestCtx) error {
	session := getSession(ctx.Tx(), nil, ctx.Header().Get(`Cookie`))
	if session != nil {
		ctx.Session().Put(`user_id`, session.GetString(`user_id`))
		ctx.Session().Put(`role_id`, session.GetString(`role_id`))
		return nil
	}
	return errors.New(`invalid session`)
}
