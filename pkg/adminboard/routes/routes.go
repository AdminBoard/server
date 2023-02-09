package routes

import (
	"fmt"
	"log"
	"strings"

	"github.com/adminboard/adminboard/pkg/adminboard/db"
	"github.com/adminboard/adminboard/pkg/adminboard/routes/admin"
	"github.com/eqto/api-server"
	"github.com/eqto/dbm"
)

func Load(svr *api.Server) {
	sysRoutes := map[string]func(api.Context) error{
		`GET /admin/menu`:          admin.Menu,
		`POST /admin/menu/add`:     admin.MenuAdd,
		`POST /admin/menu/reorder`: admin.MenuReorder,
		`POST /admin/menu/update`:  admin.MenuUpdate,
		`POST /admin/menu/status`:  admin.MenuStatus,

		`POST /admin/pages`:       admin.Pages,
		`POST /admin/pages/add`:   admin.PageAdd,
		`POST /admin/page/update`: admin.PageUpdate,

		`POST /admin/apis`:              admin.Apis,
		`POST /admin/apis/add`:          admin.ApisAdd,
		`GET /admin/apis/details`:       admin.ApisDetails,
		`POST /admin/apis/query/add`:    admin.ApisQueryAdd,
		`POST /admin/apis/query/update`: admin.ApisQueryUpdate,
		`POST /admin/apis/add_group`:    admin.ApisGroupsAdd,

		`POST /admin/groups`:     admin.Groups,
		`POST /admin/groups/add`: admin.GroupsAdd,

		`POST /admin/users`: admin.Users,

		`GET /admin/search/groups`: admin.SearchGroups,

		`POST /login`:  nil,
		`GET /logout`:  nil,
		`GET /session`: nil,
		`GET /menu`:    nil,
		`GET /page`:    nil,
	}

	g := svr.Group(`api`)
	for path, route := range sysRoutes {
		admin.Routes = append(admin.Routes, path)
		if route == nil {
			continue
		}
		split := strings.SplitN(path, ` `, 2)
		switch split[0] {
		case `GET`:
			g.Get(split[1]).Secure().AddAction(route)
		case `POST`:
			g.Post(split[1]).Secure().AddAction(route)
		default:
			log.Println(fmt.Sprintf(`invalid method %s`, path))
		}
	}
	svr.Post(`/login`).AddAction(Login)
	svr.Get(`/logout`).AddAction(Logout)
	svr.Get(`/session`).AddAction(Session)

	svr.Get(`/menu`).Secure().AddAction(Menu)
	svr.Get(`/page`).Secure().AddAction(Page)

}

func LoadFromDatabase(svr *api.Server) {
	stmt := dbm.Select(`path, method, secure, query, params, property`).
		From(db.Prefix(`api a`)).InnerJoin(db.Prefix(`api_query q`), `a.id = q.api_id`).OrderBy(`a.id, q.sequence`)
	cn := db.CN()

	g := svr.Group(`api`)
	if routes, e := cn.Select(cn.SQL(stmt)); e == nil {
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
			r.AddQueryAction(db.QueryWithPrefix(route.String(`query`)), route.String(`params`)).AssignTo(route.String(`property`))
		}
	}

}
