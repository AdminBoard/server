package routes

import (
	"strconv"
	"strings"

	"github.com/adminboard/server/pkg/adminboard/db"
	"github.com/adminboard/server/pkg/adminboard/session"
	"github.com/eqto/api-server"
	"github.com/eqto/dbm"
	"github.com/eqto/go-json"
)

func Menu(ctx api.Context) error {
	cn, e := ctx.Database()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}

	stmt := dbm.Select(`m.id, m.parent_id, m.name, m.icon, m.sequence, p.path`).
		From(db.Prefix(`menu m`)).
		InnerJoin(db.Prefix(`page p`), `m.page_id = p.id`)

	token, e := session.Get(ctx)
	params := []any{}

	if token.GroupID > 1 {
		stmt.InnerJoin(db.Prefix(`permission_item pi`), `p.id = pi.ref_id AND pi.ref_type = 'page'`).
			InnerJoin(db.Prefix(`group_permission gp`), `pi.permission_id = gp.permission_id`).
			Where(`p.status = 'publish'`).And(`gp.group_id = ?`).OrderBy(`sequence`)
		params = append(params, token.GroupID)
	} else if token.IsValid {
		stmt.Where(`p.status = 'publish'`).OrderBy(`sequence`)
	}

	rootMenus := map[int][]json.Object{}

	rs, e := cn.Select(cn.SQL(stmt), params...)
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}

	parentIDs := []string{}
	for _, r := range rs {
		parentID := r.Int(`parent_id`)
		children := []json.Object{}
		if m, ok := rootMenus[parentID]; ok {
			children = m
		}

		children = append(children, json.Object{
			`id`:       r.Int(`id`),
			`name`:     r.String(`name`),
			`icon`:     r.String(`icon`),
			`path`:     r.String(`path`),
			`sequence`: r.Int(`sequence`),
		})

		rootMenus[parentID] = children
		if parentID > 0 {
			parentIDs = append(parentIDs, strconv.Itoa(parentID))
		} else {
			parentIDs = append(parentIDs, r.String(`id`))
		}
	}

	menus := []json.Object{}

	if len(parentIDs) > 0 {
		parents := strings.Join(parentIDs, `, `)
		stmt = dbm.Select(`m.id, m.name, m.icon, m.sequence, p.path`).From(db.Prefix(`menu m`)).LeftJoin(db.Prefix(`page p`), `m.page_id = p.id`) //
		stmt.Where(`m.id IN (` + parents + `)`).OrderBy(`m.sequence`)
		rs, e = cn.Select(cn.SQL(stmt))
		if e != nil {
			return ctx.StatusInternalServerError(e.Error())
		}

		for _, r := range rs {
			id := r.Int(`id`)
			menu := json.Object{
				`id`:       id,
				`name`:     r.String(`name`),
				`icon`:     r.String(`icon`),
				`path`:     r.String(`path`),
				`sequence`: r.Int(`sequence`),
				`children`: rootMenus[id],
			}
			menus = append(menus, menu)
		}
	}

	menus = append(menus, json.Object{
		`id`:       0,
		`name`:     `Logout`,
		`icon`:     `logout`,
		`path`:     `/user/logout`,
		`sequence`: 1000,
	})

	return ctx.Write(menus)
}
