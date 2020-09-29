package main

import (
	"github.com/adminboard/server/internal/app"
	"github.com/eqto/command"
	log "github.com/eqto/go-logger"
)

var (
	svr *command.Service
)

func main() {
	log.SetFile(app.DefaultLog)
	switch command.Cmd() {
	case `start`:
		command.StartService(`run`)
	case `stop`:
		command.StopService()
	case `run`:
		if e := app.Run(); e != nil {
			log.E(e)
		}
	default:
		println(`Usage: adminboard [start/stop]`)
	}
}
