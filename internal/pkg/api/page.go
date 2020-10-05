package api

import (
	"fmt"

	"github.com/adminboard/server/internal/pkg/query"
	"github.com/eqto/api-server"
	"github.com/eqto/go-db"
	"github.com/eqto/go-json"
)

func page(ctx api.Context, q string, param interface{}) (interface{}, error) {
	rs, e := ctx.Tx().Get(q, param)
	if e != nil {
		return nil, e
	}
	if rs != nil {
		jsResp := json.Object{}
		jsResp.Put(`id`, rs.Int(`id`)).Put(`title`, rs.String(`title`))
		rsWidgets, e := ctx.Tx().Select(query.Get(query.PageContent), rs.Int(`id`))
		if e != nil {
			return nil, e
		}
		widgets := [][]json.Object{}
		var row []json.Object
		for _, rs := range rsWidgets {
			if rs.Int(`sequence`)%100 == 1 {
				if len(row) > 0 {
					widgets = append(widgets, row)
				}
				row = []json.Object{parseWidget(rs)}
			} else {
				row = append(row, parseWidget(rs))
			}
		}
		if len(row) > 0 {
			widgets = append(widgets, row)
		}
		jsResp.Put(`widgets`, widgets)
		return jsResp, nil
	}
	return api.ResponseError(api.StatusNotFound, fmt.Errorf(`Page %v not found`, param))
}

//PageByID ..
func PageByID(ctx api.Context, id int) (interface{}, error) {
	return page(ctx, query.Get(query.PageByID), id)
}

//PageByPath ...
func PageByPath(ctx api.Context, path string) (interface{}, error) {
	return page(ctx, query.Get(query.PageByPath), path)
}

func parseWidget(rs db.Resultset) json.Object {
	js := json.Object{
		`id`:          rs.Int(`id`),
		`sequence`:    rs.Int(`sequence`),
		`name`:        rs.String(`name`),
		`data_source`: rs.String(`data_source`),
	}
	params, _ := json.ParseString(rs.String(`params`))
	js.Put(`params`, params)
	return js
}
