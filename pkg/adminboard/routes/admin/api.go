package admin

import (
	"log"
	"strings"

	"github.com/adminboard/adminboard/pkg/adminboard/db"
	"github.com/eqto/api-server"
	"github.com/eqto/dbm"
	"github.com/eqto/go-json"
)

func Apis(ctx api.Context) error {
	stmt := dbm.Select(`*`).From(db.Prefix(`api`)).Where(`status IS NOT NULL`).OrderBy(`url`)

	cn, e := ctx.Database()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	rs, e := cn.Select(cn.SQL(stmt))
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	return ctx.Write(rs)
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
		`url`:         js.GetString(`url`),
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
	apiID := ctx.URL().Query().Get(`id`)
	if apiID == `` {
		return ctx.StatusBadRequest(`invalid id`)
	}
	cn, e := ctx.Database()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}

	result := dbm.Resultset{}

	data := ctx.URL().Query().Get(`data`)
	split := strings.Split(data, `,`)

	for _, str := range split {
		switch strings.TrimSpace(str) {
		case `queries`:
			stmt := dbm.Select(`id, query, params, property, sequence`).
				From(db.Prefix(`api_query`)).Where(`api_id = ?`).OrderBy(`sequence`)

			rs, e := cn.Select(cn.SQL(stmt), apiID)
			if e != nil {
				return ctx.StatusInternalServerError(e.Error())
			}
			result[`queries`] = rs
		case `groups`:
			stmt := dbm.Select(`g.id, g.name AS label`).
				From(db.Prefix(`group_api ga`)).InnerJoin(db.Prefix(`group g`), `g.id = ga.group_id`).
				Where(`ga.api_id = ?`, `g.status = 'enabled'`).OrderBy(`g.id`)

			rs, e := cn.Select(cn.SQL(stmt), apiID)
			if e != nil {
				return ctx.StatusInternalServerError(e.Error())
			}
			result[`groups`] = rs
		}

	}

	return ctx.Write(result)
}

func ApisQueryAdd(ctx api.Context) error {
	apiID := strings.TrimSpace(ctx.URL().Query().Get(`api_id`))
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
