package routes

import (
	"github.com/adminboard/adminboard/pkg/adminboard/db"
	"github.com/adminboard/adminboard/pkg/adminboard/routes/admin"
	"github.com/eqto/api-server"
	"github.com/eqto/dbm"
)

func Load(svr *api.Server) {
	g := svr.Group(`api`)
	g.GetAction(admin.Menu).Secure()
	g.Post(`/admin/menu/add`).Secure().AddAction(admin.MenuAdd)
	g.Post(`/admin/menu/reorder`).Secure().AddAction(admin.MenuReorder)
	g.Post(`/admin/menu/update`).Secure().AddAction(admin.MenuUpdate)
	g.Post(`/admin/menu/status`).Secure().AddAction(admin.MenuStatus)

	g.PostAction(admin.Pages).Secure()
	g.Post(`/admin/pages/add`).Secure().AddAction(admin.PageAdd)
	g.Post(`/admin/page/update`).Secure().AddAction(admin.PageUpdate)

	g.Post(`/admin/apis`).Secure().AddAction(admin.Apis)
	g.Post(`/admin/apis/add`).Secure().AddAction(admin.ApisAdd)
	g.Get(`/admin/apis/details`).Secure().AddAction(admin.ApisDetails)

	svr.Post(`/login`).AddAction(Login)
	svr.Get(`/logout`).AddAction(Logout)
	svr.Get(`/session`).AddAction(Session)

	svr.Get(`/menu`).Secure().AddAction(Menu)
	svr.Get(`/page`).Secure().AddAction(Page)
}

func LoadFromDatabase(svr *api.Server) {
	stmt := dbm.Select(`url, method, secure, query, params, property`).
		From(db.Prefix(`api a`)).InnerJoin(db.Prefix(`api_query q`), `a.id = q.api_id`).OrderBy(`a.id, q.sequence`)
	cn := db.CN()

	g := svr.Group(`api`)
	if routes, e := cn.Select(cn.SQL(stmt)); e == nil {
		for _, route := range routes {
			url := route.String(`url`)

			var r *api.Route
			switch route.String(`method`) {
			case api.MethodPost:
				r = g.Post(url)
			case api.MethodGet:
				r = g.Get(url)
			}
			if route.Int(`secure`) == 1 {
				r.Secure()
			}
			r.AddQueryAction(route.String(`query`), route.String(`params`)).AssignTo(route.String(`property`))
		}
	}

}
