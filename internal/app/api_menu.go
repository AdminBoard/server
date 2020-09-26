package app

import (
	"github.com/eqto/api-server"
	"github.com/eqto/go-db"
	"github.com/eqto/go-json"
)

func apiMenu(ctx api.Context) (interface{}, error) {
	rs, e := ctx.Tx().Get(getQuery(queryPage), ctx.Request().JSONBody().GetString(`path`))
	if e != nil {
		return nil, e
	}
	if rs != nil {
		jsResp := json.Object{}
		jsResp.Put(`id`, rs.Int(`id`)).Put(`title`, rs.String(`title`))
		rsContent, e := ctx.Tx().Select(getQuery(queryPageContent), rs.Int(`id`))
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
	return nil, nil
}
