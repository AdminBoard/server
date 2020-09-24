package app

import (
	"errors"
	"fmt"

	"github.com/eqto/go-db"

	log "github.com/eqto/go-logger"

	"github.com/eqto/api-server"

	"github.com/adminboard/server/internal/pkg/util"
	"github.com/eqto/config"
)

var (
	cfg *config.Config
	svr *api.Server
)

func openConfig() error {
	c, e := util.LoadConfig()
	if e != nil {
		return e
	} else if c == nil {
		//TODO no config found, run install setup
		return errors.New(`no config file founds`)
	}
	cfg = c
	return nil
}

//Run ...
func Run() error {
	if e := openConfig(); e != nil {
		return e
	}
	if logFile := cfg.Get(`Server.log_file`); logFile != `` {
		log.SetFile(logFile)
	}

	svr = api.New()
	svr.SetLogger(log.D, log.W, log.E)

	//Open database connection
	hostname := cfg.GetOr(`Database.hostname`, `localhost`)
	port := cfg.GetIntOr(`Database.port`, 3306)
	username := cfg.GetOr(`Database.username`, `adminboard`)
	password := cfg.GetOr(`Database.password`, `adminboard`)
	name := cfg.GetOr(`Database.name`, `adminboard`)
	log.D(fmt.Sprintf(`Open database %s:xxx@%s:%d/%s`, username, hostname, port, name))
	if e := svr.OpenDatabase(hostname, port, username, password, name); e != nil {
		log.D(fmt.Sprintf(`%s:xxx@%s:%d/%s`, username, hostname, port, name))
		return e
	}

	if e := loadRoutes(); e != nil {
		return e
	}
	return svr.Serve(cfg.GetIntOr(`Server.port`, 8100))
}

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
