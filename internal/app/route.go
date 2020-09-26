package app

import (
	"github.com/eqto/api-server"
	"github.com/eqto/go-db"
)

func loadRoutes() error {
	prefix := cfg.GetOr(`Database.prefix`, `admin_`)
	// dashboardPrefix := cfg.Get(`Dashboard.prefix_url`)
	// apiPrefix := cfg.Get(`API.prefix_url`)

	rsRoutes, e := svr.Database().Select(`SELECT * FROM ` + prefix + `route ORDER BY path`)
	if e != nil {
		return e
	}
	for _, rsRoute := range rsRoutes {
		route, e := loadRoute(svr.Database(), rsRoute.Int(`id`), rsRoute.String(`method`), rsRoute.String(`path`))
		if e != nil {
			return e
		}
		svr.SetRoute(route)

	}

	svr.SetRoute(api.NewFuncRoute(`/api`, apiFunc))

	return nil
}

func loadRoute(cn *db.Connection, id int, method, path string) (*api.Route, error) {
	prefix := cfg.GetOr(`Database.prefix`, `admin_`)
	route := api.NewRoute(method, path)
	rsActions, e := cn.Select(`SELECT * FROM `+prefix+`action WHERE route_id = ? AND sequence > 0 ORDER BY sequence`, id)
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
