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

//Run ...
func Run() error {
	//Start open config file
	cfg, e := util.LoadConfig()
	if e != nil {
		return e
	} else if cfg == nil {
		//TODO no config found, run install setup
		return errors.New(`no config`)
	}
	if logFile := cfg.Get(`Server.log_file`); logFile != `` {
		log.SetFile(logFile)
	}

	//Start database connection
	hostname := cfg.GetOr(`Database.hostname`, `localhost`)
	port := cfg.GetIntOr(`Database.port`, 3306)
	username := cfg.GetOr(`Database.username`, `adminboard`)
	password := cfg.GetOr(`Database.password`, `adminboard`)
	name := cfg.GetOr(`Database.name`, `adminboard`)
	cn, e := db.NewConnection(hostname, port, username, password, name)
	if e != nil {
		log.D(fmt.Sprintf(`%s:xxx@%s:%d/%s`, username, hostname, port, name))
		return e
	}
	log.D(cn)
	// svr = api.New()
	return nil

}
