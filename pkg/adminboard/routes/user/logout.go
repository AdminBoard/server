package user

import "github.com/eqto/api-server"

func Logout(ctx api.Context) error {
	ctx.Response().Header().SetCookie(`adminboard`, ``, 0)
	return nil
}
