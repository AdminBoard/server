package routes

import (
	"fmt"
	"log"
	"strings"

	"github.com/adminboard/adminboard/pkg/adminboard/routes/admin"
	"github.com/eqto/api-server"
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

		`POST /admin/groups`:                   admin.Groups,
		`POST /admin/groups/add`:               admin.GroupsAdd,
		`GET /admin/groups/details`:            admin.GroupsDetails,
		`POST /admin/groups/add_permission`:    admin.GroupsAddPermissions,
		`POST /admin/groups/remove_permission`: admin.GroupsRemovePermissions,

		`POST /admin/users`: admin.Users,

		`GET /admin/search/groups`: admin.SearchGroups,

		`POST /admin/permissions`:             admin.Permissions,
		`POST /admin/permissions/add`:         admin.PermissionsAdd,
		`GET /admin/permissions/details`:      admin.PermissionsDetails,
		`POST /admin/permissions/update`:      admin.PermissionsUpdate,
		`POST /admin/permissions/add_item`:    admin.PermissionsAddItem,
		`POST /admin/permissions/remove_item`: admin.PermissionsRemoveItem,

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
