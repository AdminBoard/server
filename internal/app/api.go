package app

import (
	"errors"

	"github.com/adminboard/server/internal/pkg/api"
	apis "github.com/eqto/api-server"
	"github.com/eqto/go-json"
)

func apiRoute(ctx apis.Context) (interface{}, error) {
	switch ctx.Request().URL().RawQuery {
	case `menu`:
		return api.Menu(ctx)
	case `page`:
		return api.Page(ctx)
	}
	return apis.ResponseError(apis.StatusNotFound, errors.New(ctx.Request().URL().RawQuery))
}

func apiPublic(ctx apis.Context) (interface{}, error) {
	switch ctx.Request().URL().RawQuery {
	case `session`:
		//TODO handle session
		return json.Object{`id`: `xx`}, nil
	}
	return apis.ResponseError(apis.StatusNotFound, errors.New(ctx.Request().URL().RawQuery))
}
