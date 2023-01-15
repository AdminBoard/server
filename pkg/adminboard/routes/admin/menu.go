package admin

import (
	"fmt"

	"github.com/adminboard/adminboard/pkg/adminboard/db"
	"github.com/eqto/api-server"
	"github.com/eqto/dbm"
	"github.com/eqto/dbm/stmt"
	"github.com/eqto/go-json"
)

func Menu(ctx api.Context) error {
	query := ctx.URL().Query()
	menuID := query.Get(`id`)

	cn := db.CN()
	if menuID != `` {
		switch query.Get(`fetch`) {
		case `groups`:
			return ctx.Write(nil)

		default:
			stmt := dbm.Select(`m.id, m.parent_id, m.name, m.icon, m.sequence, m.status`).From(db.Prefix(`menu m`)).Where(`id = ?`)

			rsMenu, e := cn.Get(cn.SQL(stmt), menuID)
			if e != nil {
				return ctx.StatusInternalServerError(e.Error())
			}
			if rsMenu == nil {
				return ctx.Status(99, `menu not found`)
			}
			return ctx.Write(rsMenu)
		}
	} else {
		stmt := dbm.Select(`m.id, m.parent_id, m.name, m.icon, m.sequence, m.status, p.url`).From(db.Prefix(`menu m`)).
			LeftJoin(db.Prefix(`page p`), `m.page_id = p.id AND p.status = 'publish'`).
			Where(`m.status IS NOT NULL`).OrderBy(`parent_id, sequence`)

		rsMenus, e := cn.Select(cn.SQL(stmt))
		if e != nil {
			return ctx.StatusInternalServerError(e.Error())
		}
		idxMap := map[int]int{}

		menus := []json.Object{}

		for idx, m := range rsMenus {
			jsMenu := json.Object{
				`id`:       m.Int(`id`),
				`name`:     m.String(`name`),
				`icon`:     m.String(`icon`),
				`url`:      m.String(`url`),
				`sequence`: m.Int(`sequence`),
				`children`: []json.Object{},
				`status`:   m.String(`status`),
			}
			parentID := m.Int(`parent_id`)

			if parentID == 0 {
				menus = append(menus, jsMenu)
				idxMap[m.Int(`id`)] = idx
			} else {
				if i, ok := idxMap[parentID]; ok {
					submenu := menus[i].GetArray(`children`)
					menus[i].Put(`children`, append(submenu, jsMenu))
				}
			}
		}
		return ctx.Write(menus)
	}
}

func MenuAdd(ctx api.Context) error {
	cn, e := ctx.Database()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	tx, e := cn.Begin()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	defer tx.Rollback()

	js, e := ctx.Request().ValidJSON(`parent_id`, `name`)
	if e != nil {
		return ctx.StatusBadRequest(e.Error())
	}
	parentID := js.GetInt(`parent_id`)

	stmtSelect := dbm.Select(`COUNT(*) AS count`).From(db.Prefix(`menu`)).Where(`parent_id = ?`)
	rs, e := tx.Get(cn.SQL(stmtSelect), parentID)
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	count := rs.Int(`count`)

	stmtInsert := dbm.InsertInto(db.Prefix(`menu`), `parent_id, name, sequence, status`).Values(`?, ?, ?, 'draft'`)
	res, e := tx.Exec(cn.SQL(stmtInsert), js.GetInt(`parent_id`), js.GetString(`name`), count)
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	id, e := res.LastInsertID()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	if e := tx.Commit(); e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	return ctx.Write(json.Object{`id`: id})
}

func MenuReorder(ctx api.Context) error {
	stmt := dbm.Update(db.Prefix(`menu`)).Set(`sequence = ?`).Where(`id = ?`)

	jsArr, e := json.ParseArray(ctx.Request().Body())
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	cn, e := ctx.Database()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}

	tx, e := ctx.Tx()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	for _, js := range jsArr {
		_, e := tx.Exec(cn.SQL(stmt), js.GetInt(`sequence`), js.GetInt(`id`))
		if e != nil {
			return ctx.StatusBadRequest(e.Error())
		}
	}
	return nil
}

func MenuUpdate(ctx api.Context) error {
	menuID := ctx.URL().Query().Get(`id`)
	if menuID == `` {
		return ctx.StatusBadRequest(`invalid menu id`)
	}

	s := dbm.Update(db.Prefix(`menu`))
	var stmtSet *stmt.UpdateFields
	params := []interface{}{}

	js := ctx.Request().JSON()

	if js.Has(`status`) {
		return ctx.StatusForbidden(`status not changed`)
	}

	cn, e := ctx.Database()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}

	if page := js.GetString(`page`); page != `` {
		stmt := dbm.Select(`id`).From(db.Prefix(`page`)).Where(`name = ?`)

		rs, e := cn.Get(cn.SQL(stmt), page)
		if e != nil {
			return ctx.StatusInternalServerError(e.Error())
		}
		if rs == nil {
			return ctx.StatusNotFound(fmt.Sprintf(`Page %s not found`, page))
		}
		stmtSet = s.Set(`page_id = ?`)
		params = append(params, rs.Int(`id`))
	} else {
		for key := range js {
			stmtSet = s.Set(key + ` = ?`)
			params = append(params, js.GetString(key))
		}
	}

	if len(params) == 0 {
		return ctx.StatusBadRequest(`no update`)
	}
	stmtSet.Where(`id = ?`)
	params = append(params, menuID)
	_, e = cn.Exec(cn.SQL(s), params...)
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}

	return nil
}

func MenuStatus(ctx api.Context) error {
	menuID := ctx.URL().Query().Get(`id`)
	if menuID == `` {
		return ctx.StatusBadRequest(`invalid menu id`)
	}
	js, e := ctx.Request().ValidJSON(`status`)
	if e != nil {
		return ctx.StatusBadRequest(e.Error())
	}

	stmt := dbm.Update(db.Prefix(`menu`)).Set(`status = ?`).Where(`id = ?`)
	cn, e := ctx.Database()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	_, e = cn.Exec(cn.SQL(stmt), js.GetString(`status`), menuID)
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	return nil
}
