package admin

import (
	"strings"

	"github.com/adminboard/adminboard/pkg/adminboard/db"
	"github.com/adminboard/adminboard/pkg/adminboard/routes/params"
	"github.com/eqto/api-server"
	"github.com/eqto/dbm"
	"github.com/eqto/go-json"
)

func Permissions(ctx api.Context) error {
	cn, e := ctx.Database()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	stmt := dbm.Select(`id, name, description`).From(db.Prefix(`permission`)).OrderBy(`name`)

	rs, e := cn.Select(cn.SQL(stmt))
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}

	return ctx.Write(rs)
}

func PermissionsAdd(ctx api.Context) error {
	cn, e := ctx.Database()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	js := ctx.Request().JSON()
	if !js.Has(`name`) || !js.Has(`description`) {
		return ctx.StatusBadRequest(`required parameter: name, description`)
	}

	stmt := dbm.InsertInto(db.Prefix(`permission`), `name, description`)

	values := []any{}

	name := js.GetString(`name`)
	if name != `` {
		stmt.Values(`?, ?`)
		values = append(values, name)
	} else {
		stmt.Values(`NULL, ?`)
	}
	values = append(values, js.GetString(`description`))

	res, e := cn.Exec(cn.SQL(stmt), values...)
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	id, e := res.LastInsertID()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}

	return ctx.Write(id)
}

func PermissionsDetails(ctx api.Context) error {
	js, e := params.FetchQuery(ctx, `id`, `data`)
	if e != nil {
		return ctx.StatusBadRequest(e.Error())
	}
	id := js.GetInt(`id`)

	cn, e := ctx.Database()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}

	stmt := dbm.Select(`id, name, description`).From(db.Prefix(`permission`)).Where(`id = ?`)
	rs, e := cn.Get(cn.SQL(stmt), id)
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}

	resp := json.Object{
		`id`:          rs.Int(`id`),
		`name`:        rs.String(`name`),
		`description`: rs.String(`description`),
	}

	split := strings.Split(js.GetString(`data`), `,`)
	for _, str := range split {
		str = strings.TrimSpace(str)
		switch str {
		case `api`:
			stmt := dbm.Select(`a.id, a.method, a.secure, a.path, a.description, a.status`).
				From(db.Prefix(`api a`)).InnerJoin(db.Prefix(`permission_item pi`), `pi.ref_id = a.id AND pi.ref_type = 'api'`).
				Where(`pi.permission_id = ?`).And(`(a.status = 'active' OR a.status = 'draft')`).OrderBy(`a.path, a.method`)
			rs, e := cn.Select(cn.SQL(stmt), id)
			if e != nil {
				return ctx.StatusInternalServerError(e.Error())
			}
			resp.Put(`apis`, rs)
		case `page`:
			stmt := dbm.Select(`p.id, p.path, p.name, p.description, p.status`).From(db.Prefix(`page p`)).
				InnerJoin(db.Prefix(`permission_item pi`), `pi.ref_id = p.id AND pi.ref_type = 'page'`).
				Where(`pi.permission_id = ?`).And(`(p.status = 'publish' OR p.status = 'draft')`).OrderBy(`p.path, p.name, p.status`)
			rs, e := cn.Select(cn.SQL(stmt), id)
			if e != nil {
				return ctx.StatusInternalServerError(e.Error())
			}
			resp.Put(`pages`, rs)
		}
	}
	return ctx.Write(resp)
}

func PermissionsUpdate(ctx api.Context) error {
	js, e := params.Fetch(ctx, `id, description`, `name`)
	if e != nil {
		return ctx.StatusBadRequest(e.Error())
	}
	name := js.GetString(`name`)
	stmt := dbm.Update(db.Prefix(`permission`))

	values := []any{}
	if name == `` {
		stmt.Set(`name = NULL`)
	} else {
		stmt.Set(`name = ?`)
		values = append(values, js.GetString(`name`))
	}
	stmt.Set(`description = ?`).Where(`id = ?`)
	values = append(values, []any{js.GetString(`description`), js.GetInt(`id`)}...)

	cn, e := ctx.Database()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	_, e = cn.Exec(cn.SQL(stmt), values...)
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	return nil
}

func PermissionsAddItem(ctx api.Context) error {
	js, e := params.FetchQuery(ctx, ``, `search`)
	if e != nil {
		return ctx.StatusBadRequest(e.Error())
	}
	cn, e := ctx.Database()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	if js.Has(`search`) {
		search := js.GetString(`search`)
		search += `*`
		js, e := params.FetchQuery(ctx, `search_type`, ``)
		if e != nil {
			return ctx.StatusBadRequest(e.Error())
		}
		var stmt any
		switch js.GetString(`search_type`) {
		case `api`:
			stmt = dbm.Select(`id, CONCAT(method, ' ', path) AS label, description AS sublabel`).From(db.Prefix(`api`)).
				Where(`MATCH(path, description) AGAINST(? IN BOOLEAN MODE)`).And(`(status = 'active' OR status = 'draft')`).
				OrderBy(`path, method`)
		case `page`:
			stmt = dbm.Select(`id, CONCAT(name, ' (', path, ')') AS label, description AS sublabel`).From(db.Prefix(`page`)).
				Where(`MATCH(path, name, description) AGAINST(? IN BOOLEAN MODE)`).And(`(status = 'publish' OR status = 'draft')`).
				OrderBy(`path, name`)
		}

		rs, e := cn.Select(cn.SQL(stmt), search)
		if e != nil {
			return ctx.StatusInternalServerError(e.Error())
		}
		return ctx.Write(rs)
	} else {
		js, e := params.Fetch(ctx, `permission_id, ref_id, ref_type`, ``)
		if e != nil {
			return ctx.StatusBadRequest(e.Error())
		}
		stmt := dbm.InsertInto(db.Prefix(`permission_item`), `permission_id, ref_id, ref_type`).Values(`?, ?, ?`)
		_, e = cn.Exec(cn.SQL(stmt), js.GetInt(`permission_id`), js.GetInt(`ref_id`), js.GetString(`ref_type`))
		if e != nil {
			return ctx.StatusInternalServerError(e.Error())
		}
	}
	return nil
}

func PermissionsRemoveItem(ctx api.Context) error {
	js, e := params.Fetch(ctx, `permission_id, ref_id, ref_type`, ``)
	if e != nil {
		return ctx.StatusBadRequest(e.Error())
	}
	cn, e := ctx.Database()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	stmt := dbm.DeleteFrom(db.Prefix(`permission_item`)).Where(`permission_id = ? AND ref_id = ? AND ref_type = ?`)

	_, e = cn.Exec(cn.SQL(stmt), js.GetInt(`permission_id`), js.GetInt(`ref_id`), js.GetString(`ref_type`))
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	return nil

}
