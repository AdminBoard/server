package app

import "github.com/eqto/api-server"

func apiRoute(ctx api.Context) (interface{}, error) {
	switch ctx.Request().URL().RawQuery {
	case `menu`:
		return apiMenu(ctx)
	case `page`:
		return apiPage(ctx)
	}
	return nil, nil
}
