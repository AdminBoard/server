package admin

import (
	"errors"

	"github.com/adminboard/adminboard/pkg/adminboard/db"
	"github.com/eqto/dbm"
	"github.com/eqto/go-json"
)

func GetMenu(activeOnly bool, groupID int) ([]json.Object, error) {
	cn := db.CN()

	stmt := dbm.Select(`m.id, m.parent_id, m.name, m.icon, m.sequence, p.url`).From(db.Prefix(`menu m`)).
		LeftJoin(db.Prefix(`page p`), `m.page_id = p.id AND p.status = 'publish'`)

	if activeOnly {
		stmt.Where(`m.status = 'active'`)
	} else {
		stmt.Where(`m.status IS NOT NULL`)
	}
	stmt.OrderBy(`parent_id, sequence`)

	menus, e := cn.Select(cn.SQL(stmt))
	if e != nil {
		return nil, e
	}

	alloweds := map[int]struct{}{}

	if groupID > 0 {
		stmt := dbm.Select(`m.id`).From(db.Prefix(`group_page gp`)).InnerJoin(db.Prefix(`menu m`), `gp.page_id = m.page_id`).
			Where(`m.status = 'active'`, `gp.group_id = ?`)

		rs, e := cn.Select(cn.SQL(stmt), groupID)
		if e != nil {
			return nil, e
		}
		if len(rs) == 0 {
			return nil, errors.New(`invalid group permission`)
		}
		for _, r := range rs {
			alloweds[r.Int(`id`)] = struct{}{}
		}
	}

	idxMap := map[int]int{}

	menu := []json.Object{}

	for idx, m := range menus {
		jsMenu := json.Object{
			`id`:       m.Int(`id`),
			`name`:     m.String(`name`),
			`icon`:     m.String(`icon`),
			`url`:      m.String(`url`),
			`sequence`: m.Int(`sequence`),
		}
		parentID := m.Int(`parent_id`)

		if parentID == 0 {
			menu = append(menu, jsMenu)
			idxMap[m.Int(`id`)] = idx
		} else {
			allow := false
			if groupID == 0 {
				allow = true
			} else if _, ok := alloweds[m.Int(`id`)]; ok {
				allow = true
			}
			if allow {
				if i, ok := idxMap[parentID]; ok {
					submenu := menu[i].GetArray(`children`)
					menu[i].Put(`children`, append(submenu, jsMenu))
				}
			}
		}
	}

	filtered := []json.Object{}
	for _, m := range menu {
		if len(m.GetArray(`children`)) > 0 || m.GetString(`url`) != `` {
			filtered = append(filtered, m)
		}
	}

	return filtered, nil
}
