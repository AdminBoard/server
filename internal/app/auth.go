package app

import (
	"errors"

	"github.com/eqto/api-server"
)

func authenticator(ctx api.Context) error {
	session := getSession(ctx)
	if session != nil {
		ctx.Session().Put(`user_id`, session.GetString(`user_id`))
		ctx.Session().Put(`role_id`, session.GetString(`role_id`))
		return nil
	}
	return errors.New(`invalid session`)
}
