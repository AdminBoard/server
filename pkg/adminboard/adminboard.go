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

func Run() error {
	port := service.GetInt(`port`)

	if e := config.Open(`config.conf`); e != nil {
		return e
	}
	db.SetPrefix(config.GetOr(`Database.table_prefix`, db.DefaultPrefix))

	if e := db.Connect(); e != nil {
		return e
	}

	svr = api.New()
	svr.SetDatabase(db.CN())
	svr.NormalizeFunc(true)
	// svr.SetPrefixPath(`/api`)

	svr.AddMiddleware(middleware.AuthMiddleware).Secure()
	svr.Group(`api`).AddMiddleware(middleware.AuthAPI).Secure()

	routes.LoadFromDatabase(svr)
	routes.Load(svr)

	return svr.Serve(port)
}

func Shutdown() {
	if svr != nil {
		svr.Shutdown()
	}
	svr = nil
}
