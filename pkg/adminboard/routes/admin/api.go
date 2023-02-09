package admin

import (
	"log"
	"sort"
	"strings"

	"github.com/adminboard/adminboard/pkg/adminboard/db"
	"github.com/eqto/api-server"
	"github.com/eqto/dbm"
	"github.com/eqto/dbm/stmt"
	"github.com/eqto/go-json"
)

func Apis(ctx api.Context) error {
	stmt := dbm.Select(`*`).From(db.Prefix(`api`)).Where(`status IS NOT NULL`).OrderBy(`path`)

	cn, e := ctx.Database()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	rs, e := cn.Select(cn.SQL(stmt))
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	jsRoutes := []json.Object{}
	for _, route := range Routes {
		split := strings.SplitN(route, ` `, 2)
		jsRoutes = append(jsRoutes, json.Object{
			`method`: split[0],
			`path`:   split[1],
			`status`: `system`,
		})
	}
	sort.SliceStable(jsRoutes, func(i, j int) bool {
		return jsRoutes[i].GetString(`path`) < jsRoutes[j].GetString(`path`)
	})

	for _, r := range rs {
		jsRoutes = append(jsRoutes, json.Object{
			`id`:          r.Int(`id`),
			`description`: r.String(`description`),
			`method`:      r.String(`method`),
			`path`:        r.String(`path`),
			`secure`:      r.Int(`secure`),
			`status`:      r.String(`status`),
		})
	}

	return ctx.Write(jsRoutes)
}

func ApisAdd(ctx api.Context) error {
	cn, e := ctx.Database()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	js := ctx.Request().JSON()
	res, e := cn.Insert(db.Prefix(`api`), map[string]any{
		`method`:      js.GetStringOr(`method`, `GET`),
		`secure`:      js.GetIntOr(`secure`, 0),
		`path`:        js.GetString(`path`),
		`description`: js.GetString(`description`),
		`status`:      `draft`,
	})
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	row, e := res.RowsAffected()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	if row == 0 {
		return ctx.StatusBadRequest(`operation failed`)
	}
	return nil
}

func ApisDetails(ctx api.Context) error {
	path := ctx.Request().QueryParam(`path`)
	if path == `` {
		return ctx.StatusBadRequest(`invalid parameter: path`)
	}
	method := ctx.Request().QueryParam(`method`)
	if method == `` {
		return ctx.StatusBadRequest(`invalid parameter: method`)
	}
	cn, e := ctx.Database()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}

	result := dbm.Resultset{}

	data := ctx.Request().QueryParam(`data`)
	split := strings.Split(data, `,`)

	for _, str := range split {
		switch strings.TrimSpace(str) {
		case `queries`:
			stmt := dbm.Select(`aq.id, aq.query, aq.params, aq.property, aq.sequence`).
				From(db.Prefix(`api_query aq`)).InnerJoin(db.Prefix(`api a`), `a.id = aq.api_id`).
				Where(`a.path = ?`).And(`a.method = ?`).OrderBy(`sequence`)

			rs, e := cn.Select(cn.SQL(stmt), path, method)
			if e != nil {
				return ctx.StatusInternalServerError(e.Error())
			}
			result[`queries`] = rs
		case `groups`:
			// stmt := dbm.Select(`g.id, g.name AS label`).
			// 	From(db.Prefix(`group_api ga`)).InnerJoin(db.Prefix(`group g`), `g.id = ga.group_id`).
			// 	Where(`ga.api_id = ?`, `g.status = 'enabled'`).OrderBy(`g.id`)

			// rs, e := cn.Select(cn.SQL(stmt), apiID)
			// if e != nil {
			// 	return ctx.StatusInternalServerError(e.Error())
			// }
			// result[`groups`] = rs
		}

	}

	return ctx.Write(result)
}

func ApisQueryAdd(ctx api.Context) error {
	apiID := strings.TrimSpace(ctx.Request().QueryParam(`api_id`))
	if apiID == `` {
		return ctx.StatusBadRequest(`invalid parameter:api_id`)
	}

	js := ctx.Request().JSON()
	query := strings.TrimSpace(js.GetString(`query`))
	if query == `` {
		return ctx.StatusBadRequest(`invalid parameter:query`)
	}

	cn, e := ctx.Database()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}

	stmt := dbm.Select(`COUNT(*) AS count`).From(db.Prefix(`api_query`)).Where(`api_id = ?`).GroupBy(`api_id`)

	rs, e := cn.Get(cn.SQL(stmt), apiID)
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	count := rs.Int(`count`)

	stmtInsert := dbm.InsertInto(db.Prefix(`api_query`), `api_id, query, params, property, sequence`).Values(`?, ?, ?, ?, ?`)
	res, e := cn.Exec(cn.SQL(stmtInsert), apiID, query, strings.TrimSpace(js.GetString(`params`)), strings.TrimSpace(js.GetString(`property`)), count)
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	id, e := res.LastInsertID()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}

	return ctx.Write(json.Object{
		`id`:       id,
		`sequence`: count,
	})
}

func ApisGroupsAdd(ctx api.Context) error {
	js := ctx.Request().JSON()
	apiID := js.GetInt(`api_id`)
	log.Println(apiID)
	if apiID == 0 {
		return ctx.StatusBadRequest(`invalid parameter:api_id`)
	}
	groupID := js.GetInt(`group_id`)
	if groupID == 0 {
		return ctx.StatusBadRequest(`invalid parameter:group_id`)
	}
	cn, e := ctx.Database()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}

	stmtInsert := dbm.InsertInto(db.Prefix(`group_api`), `group_id, api_id`).Values(`?, ?`)
	res, e := cn.Exec(cn.SQL(stmtInsert), groupID, apiID)
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	row, e := res.RowsAffected()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	if row == 0 {
		return ctx.StatusInternalServerError(`affected row: 0`)
	}

	return nil
}

func ApisQueryUpdate(ctx api.Context) error {
	id := ctx.Request().QueryParam(`id`)
	if id == `` {
		return ctx.StatusBadRequest(`invalid parameter: id`)
	}
	s := dbm.Update(db.Prefix(`api_query`))
	var stmtSet *stmt.UpdateFields
	params := []any{}

	js := ctx.Request().JSON()
	for key := range js {
		stmtSet = s.Set(key + ` = ?`)
		params = append(params, js.GetString(key))
	}
	if len(params) == 0 {
		return ctx.StatusBadRequest(`no update`)
	}

	stmtSet.Where(`id = ?`)
	params = append(params, id)

	cn, e := ctx.Database()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	_, e = cn.Exec(cn.SQL(s), params...)
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}

	return nil
}
