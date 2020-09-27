package app

import (
	"github.com/adminboard/server/internal/pkg/query"
	"github.com/eqto/api-server"
	"github.com/eqto/go-db"
)

func loadRoutes() error {
	// dashboardPrefix := cfg.Get(`Dashboard.prefix_url`)
	// apiPrefix := cfg.Get(`API.prefix_url`)

	rsRoutes, e := svr.Database().Select(query.Get(query.Route))
	if e != nil {
		return e
	}
	for _, rsRoute := range rsRoutes {
		route, e := loadRoute(svr.Database(), rsRoute.Int(`id`))
		if e != nil {
			return e
		}
		svr.SetRoute(rsRoute.String(`method`), rsRoute.String(`path`), route)

	}

	return nil
}

func loadRoute(cn *db.Connection, id int) (*api.Route, error) {
	route := api.NewRoute()
	rsActions, e := cn.Select(query.Get(query.Action), id)
	if e != nil {
		return nil, e
	}
	for _, rsAction := range rsActions {
		if query := rsAction.String(`query`); query != `` {
			route.AddQueryAction(query, rsAction.String(`params`), rsAction.String(`property`))
		}
		//TODO if not query action
	}

	return route, nil
}
