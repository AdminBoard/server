package main

import (
	"github.com/adminboard/server/internal/app"
	"github.com/eqto/command"
	log "github.com/eqto/go-logger"
	_ "github.com/go-sql-driver/mysql"
)

var (
	svr *command.Service
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.E(r)
		}
	}()

	log.SetFile(app.DefaultLog)
	switch command.Cmd() {
	case `start`:
		command.StartService(`run`)
	case `stop`:
		command.StopService()
	case `run`:
		if e := app.Run(); e != nil {
			panic(e)
		}
	default:
		println(`Usage: adminboard [start/stop]`)
	}
}
