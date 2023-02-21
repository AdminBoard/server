package adminboard

import (
	"github.com/adminboard/adminboard/pkg/adminboard/db"
	"github.com/adminboard/adminboard/pkg/adminboard/routes"
	"github.com/adminboard/adminboard/pkg/adminboard/routes/middleware"
	"github.com/eqto/api-server"
	"github.com/eqto/config"
	"github.com/eqto/go-json"
	"github.com/eqto/service"
)

type RouteFunc func(ctx api.Context) error
type MiddlewareFunc func(ctx api.Context) error

type PrepareSession func(session json.Object) error

const (
	MethodGet  = api.MethodGet
	MethodPost = api.MethodPost
)

var svr *api.Server

func init() {
	svr = api.New()
	svr.SetPrefixPath(`/api`)
	svr.NormalizeFunc(true)
	svr.AddMiddleware(middleware.AuthMiddleware).Secure()
	svr.Group(`api`).AddMiddleware(middleware.AuthAPI).Secure()
}

func Run() error {
	port := service.GetInt(`port`)

	if e := config.Open(`config.conf`); e != nil {
		return e
	}
	db.SetPrefix(config.GetOr(`Database.table_prefix`, db.DefaultPrefix))

	if e := db.Connect(); e != nil {
		return e
	}

	svr.SetDatabase(db.CN())

	routes.LoadFromDatabase(svr)
	routes.Load(svr)

	return svr.Serve(port)
}

func Server() *api.Group {
	return svr.Group(`api`)
}

func Shutdown() {
	if svr != nil {
		svr.Shutdown()
	}
	svr = nil
}

func RegisterApi(method, path string, route RouteFunc) {
	g := svr.Group(`api`)

	switch method {
	case api.MethodGet:
		g.Get(path).AddAction(route)
	case api.MethodPost:
		g.Post(path).AddAction(route)
	}
}

func RegisterSecureApi(method, path string, route RouteFunc) {
	g := svr.Group(`api`)
	switch method {
	case api.MethodGet:
		g.Get(path).Secure().AddAction(route)
	case api.MethodPost:
		g.Post(path).Secure().AddAction(route)
	}
}

func AddMiddleware(m MiddlewareFunc) {
	svr.AddMiddleware(m)
}
