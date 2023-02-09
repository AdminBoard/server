package routes

import (
	"strconv"
	"strings"

	"github.com/adminboard/adminboard/pkg/adminboard/db"
	"github.com/eqto/api-server"
	"github.com/eqto/dbm"
	"github.com/eqto/go-json"
)

func Menu(ctx api.Context) error {
	cn, e := ctx.Database()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}

	groupID := ctx.Session().GetInt(`groupID`)

	stmtMenu := dbm.Select(`m.id, m.parent_id, m.name, m.icon, m.sequence, p.path`).
		From(db.Prefix(`menu m`)).
		InnerJoin(db.Prefix(`page p`), `m.page_id = p.id`)

	if groupID > 0 {
		stmtMenu.InnerJoin(db.Prefix(`group_page gp`), `gp.page_id = p.id`).
			Where(`p.status = 'publish'`).And(`gp.group_id = ?`).OrderBy(`sequence`)
	} else {
		stmtMenu.Where(`p.status = 'publish'`).OrderBy(`sequence`)
	}

	rootMenus := map[int][]json.Object{}

	rs, e := cn.Select(cn.SQL(stmtMenu))
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}

	parentIDs := []string{}
	for _, r := range rs {
		parentID := r.Int(`parent_id`)
		parentIDs = append(parentIDs, strconv.Itoa(parentID))
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
	}

	parents := strings.Join(parentIDs, `, `)

	rs, e = cn.Select(db.QueryWithPrefix(`SELECT m.id, m.name, m.icon, p.path, m.sequence FROM {prefix}menu m 
		LEFT JOIN {prefix}page p ON m.page_id = p.id WHERE m.id IN (` + parents + `) OR (m.parent_id = 0 AND p.path != '') ORDER BY m.sequence`))
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}

	menus := []json.Object{}

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

	menus = append(menus, json.Object{
		`id`:       0,
		`name`:     `Logout`,
		`icon`:     `logout`,
		`path`:     `/user/logout`,
		`sequence`: 1000,
	})

	return ctx.Write(menus)
}
