package app

import (
	"errors"
	"strings"

	"github.com/adminboard/server/internal/pkg/api"
	apis "github.com/eqto/api-server"
	"github.com/eqto/go-json"
)

func apiRoute(ctx apis.Context) (interface{}, error) {
	command := ctx.Request().URL().RawQuery
	subcommand := ``
	if idx := strings.Index(command, `=`); idx >= 0 {
		subcommand = command[idx+1:]
		command = command[:idx]
	}
	switch command {
	case `menu`:
		return api.Menu(ctx)
	case `page`:
		return api.Page(ctx, subcommand)
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
