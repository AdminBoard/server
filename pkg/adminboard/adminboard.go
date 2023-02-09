package adminboard

import (
	"github.com/adminboard/adminboard/pkg/adminboard/db"
	"github.com/adminboard/adminboard/pkg/adminboard/routes"
	"github.com/adminboard/adminboard/pkg/adminboard/routes/middleware"
	"github.com/eqto/api-server"
	"github.com/eqto/config"
	"github.com/eqto/service"
)

var svr *api.Server

func init() {
	svr = api.New()
	svr.SetPrefixPath(`/api`)
	svr.NormalizeFunc(true)
	svr.AddMiddleware(middleware.AuthStore).Secure()
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
