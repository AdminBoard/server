package main

import (
	"github.com/adminboard/server/internal/app"
	log "github.com/eqto/go-logger"
	"github.com/eqto/go-service"
)

var (
	svr *service.Service
)

func main() {
	log.SetFile(app.DefaultLog)
	switch service.Cmd() {
	case `start`:
		service.StartService(`run`)
	case `stop`:
		service.StopService()
	case `run`:
		if e := app.Run(); e != nil {
			log.E(e)
		}
	default:
		println(`Usage: adminboard [start/stop]`)
	}
}
