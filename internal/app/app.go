package app

import (
	"fmt"

	"github.com/adminboard/server/internal/pkg/config"

	log "github.com/eqto/go-logger"

	"github.com/eqto/api-server"
)

var (
	svr *api.Server
)

//APIServer ..
func APIServer() *api.Server {
	return svr
}

//Init init server before serving. This function used if you need modify server or any property before serving http. If you don't need modification just run Serve without Init
func Init() error {
	if e := config.Load(); e != nil {
		return e
	}
	if logFile := config.Get(`Server.log_file`); logFile != `` {
		log.SetFile(logFile)
	}

	svr = api.New()
	svr.SetLogger(log.D, log.W, log.E)

	//Open database connection
	hostname := config.GetOr(`Database.hostname`, `localhost`)
	port := config.GetIntOr(`Database.port`, 3306)
	username := config.GetOr(`Database.username`, `adminboard`)
	password := config.GetOr(`Database.password`, `adminboard`)
	name := config.GetOr(`Database.name`, `adminboard`)
	log.D(fmt.Sprintf(`Open database %s:xxx@%s:%d/%s`, username, hostname, port, name))
	if e := svr.OpenDatabase(hostname, port, username, password, name); e != nil {
		log.D(fmt.Sprintf(`%s:xxx@%s:%d/%s`, username, hostname, port, name))
		return e
	}

	if e := loadRoutes(); e != nil {
		return e
	}
	route := api.NewFuncRoute(apiRoute)
	svr.SetRoute(api.MethodPost, `/api`, route)
	svr.SetRoute(api.MethodGet, `/api`, route)

	publicRoute := api.NewFuncRoute(apiPublic)
	publicRoute.SetSecure(false)
	svr.SetRoute(api.MethodPost, `/api/public`, publicRoute)
	svr.SetRoute(api.MethodGet, `/api/public`, publicRoute)

	svr.AddAuthMiddleware(authMiddleware)
	return nil
}

//Run run and wait until finish or error
func Run() error {
	if svr == nil {
		if e := Init(); e != nil {
			return e
		}
	}
	return svr.Serve(config.GetIntOr(`Server.port`, 8100))
}
