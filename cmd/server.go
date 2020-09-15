package main

import (
	"errors"

	"github.com/adminboard/server/internal/pkg/util"
	"github.com/eqto/config"
	log "github.com/eqto/go-logger"
	"github.com/eqto/go-service"

	"github.com/adminboard/server/internal/pkg/constant"
)

func main() {
	switch service.Cmd() {
	case `start`:

	case `stop`:

	case `run`:
		run()
	default:
		println(`Usage: adminboard [start/stop]`)
	}
}

func run() {
	defer func() {
		if r := recover(); r != nil {
			log.E(r)
		}
	}()
	var cfg *config.Config
	for _, file := range constant.ConfigLocations() {
		if util.FileExist(file) {
			c, e := config.ParseFile(file)
			if e != nil {
				panic(e)
			}
			cfg = c
			break
		}
	}
	if cfg == nil {
		panic(errors.New(`no config file founds`))
	}
	log.D(cfg)
	// config.ParseFile(`config.conf`)

}
