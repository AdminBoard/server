package main

import (
	"log"

	"github.com/adminboard/adminboard/pkg/adminboard"
	"github.com/eqto/service"
)

func main() {
	defer service.HandlePanic()
	service.OnPanic(log.Println)
	service.OnStop(adminboard.Shutdown)

	if e := service.Run(adminboard.Run); e != nil {
		panic(e)
	}
}
