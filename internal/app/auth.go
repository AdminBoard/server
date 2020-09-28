package app

import "github.com/eqto/api-server"

func authMiddleware(ctx api.RequestCtx) error {
	ctx.Session().Put(`role_id`, 1)
	ctx.Session().Put(`user_id`, 2)
	return nil
}
