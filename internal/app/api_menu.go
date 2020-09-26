package app

import (
	"sort"

	"github.com/adminboard/server/internal/pkg/query"
	"github.com/eqto/api-server"
	"github.com/eqto/go-json"
)

func apiMenu(ctx api.Context) (interface{}, error) {
	roleID := ctx.Session().GetInt(`role_id`)

	rsMenu, e := ctx.Tx().Select(query.Get(query.Menu), roleID)
	if e != nil {
		return nil, e
	}
	menuMap := make(map[int]json.Object)

	for _, rs := range rsMenu {
		js := json.Object{
			`id`:          rs.Int(`id`),
			`kind`:        rs.String(`kind`),
			`caption`:     rs.String(`caption`),
			`description`: rs.String(`description`),
			`sequence`:    rs.String(`sequence`),
		}
		if parentID := rs.Int(`parent_id`); parentID == 0 {
			menuMap[js.GetInt(`id`)] = js
		} else {
			submenu := menuMap[parentID].GetArray(`submenu`)
			submenu = append(submenu, js)
			menuMap[parentID].Put(`submenu`, submenu)
		}
	}
	menu := []json.Object{}
	for _, m := range menuMap {
		menu = append(menu, m)
	}
	sort.SliceStable(menu, func(i, j int) bool {
		return menu[i].GetInt(`sequence`) < menu[j].GetInt(`sequence`)
	})
	return menu, nil
}
