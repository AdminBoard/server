package api

import (
	"fmt"

	"github.com/adminboard/server/internal/pkg/query"
	"github.com/eqto/api-server"
	"github.com/eqto/go-db"
	"github.com/eqto/go-json"
)

//Page ...
func Page(ctx api.Context, path string) (interface{}, error) {
	rs, e := ctx.Tx().Get(query.Get(query.Page), path)
	if e != nil {
		return nil, e
	}
	if rs != nil {
		jsResp := json.Object{}
		jsResp.Put(`id`, rs.Int(`id`)).Put(`title`, rs.String(`title`))
		rsContent, e := ctx.Tx().Select(query.Get(query.PageContent), rs.Int(`id`))
		if e != nil {
			return nil, e
		}
		content := [][]db.Resultset{}
		var row []db.Resultset
		for _, rs := range rsContent {
			if rs.Int(`sequence`)%100 == 1 {
				if len(row) > 0 {
					content = append(content, row)
				}
				row = []db.Resultset{rs}
			} else {
				row = append(row, rs)
			}
		}
		if len(row) > 0 {
			content = append(content, row)
		}
		jsResp.Put(`content`, content)
		return jsResp, nil
	}
	return api.ResponseError(api.StatusNotFound, fmt.Errorf(`Page %s not found`, path))
}
