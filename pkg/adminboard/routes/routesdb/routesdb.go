package routesdb

import (
	"log"

	"github.com/adminboard/adminboard/pkg/adminboard/db"
	"github.com/eqto/api-server"
	"github.com/eqto/dbm"
)

var server *api.Server

func SetServer(svr *api.Server) {
	server = svr
}

// Load load all apis (apiID = 0) or single api when apiID != 0
func Load(apiID int) {
	params := []any{}
	stmt := dbm.Select(`path, method, secure, query, parameter, property`).
		From(db.Prefix(`api a`)).InnerJoin(db.Prefix(`api_query q`), `a.id = q.api_id`)
	if apiID > 0 {
		stmt.Where(`a.id = ?`)
		params = append(params, apiID)
	}
	stmt.OrderBy(`a.id, q.sequence`)

	cn := db.CN()

	g := server.Group(`api`)
	if routes, e := cn.Select(cn.SQL(stmt), params...); e == nil {
		for _, route := range routes {
			path := route.String(`path`)

			var r *api.Route
			switch route.String(`method`) {
			case api.MethodPost:
				r = g.Post(path)
			case api.MethodGet:
				r = g.Get(path)
			}
			if route.Int(`secure`) == 1 {
				r.Secure()
			}
			r.AddQueryAction(db.QueryWithPrefix(route.String(`query`)), route.String(`parameter`)).AssignTo(route.String(`property`))
		}
	} else {
		log.Println(e.Error())
	}

}

// Unload
func Unload(apiID int) {

	//TODO

}
