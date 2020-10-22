package adminboard

import "github.com/adminboard/server/internal/app"

//StaticProxy ...
func StaticProxy(url string) {
	app.StaticProxy(url)
}

//Run ...
func Run() error {
	return app.Run()
}
