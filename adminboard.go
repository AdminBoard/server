package adminboard

import (
	"github.com/adminboard/server/internal/app"
	"github.com/eqto/api-server"
)

var (
	proxy string
)

//Proxy ...
func Proxy(url string) {
	proxy = url
}

//Server ...
func Server() *api.Server {
	app.Init()
	return app.APIServer()
}

//Run ...
func Run() error {
	if e := app.Init(); e != nil {
		return e
	}
	if proxy != `` {
		app.APIServer().Proxy(`*`, proxy)
	}

	return app.Run()
}
