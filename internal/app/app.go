package app

import (
	"errors"
	"fmt"

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
