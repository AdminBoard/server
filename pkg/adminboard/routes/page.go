package routes

import (
	"github.com/adminboard/adminboard/pkg/adminboard/db"
	"github.com/eqto/api-server"
)

func Page(ctx api.Context) error {
	url := ctx.Request().QueryParam(`url`)
	if url == `` {
		return ctx.StatusBadRequest(`invalid page url:` + url)
	}
	groupID := ctx.Session().GetInt(`groupID`)

	rs, e := db.CN().Get(db.QueryWithPrefix(`SELECT p.title, p.description, p.layout, gp.group_id FROM {prefix}page p
		LEFT JOIN {prefix}group_page gp ON p.id = gp.page_id AND gp.group_id = ? WHERE url = ?`), groupID, url)
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	if rs.Int(`group_id`) == 0 {
		return ctx.StatusForbidden(`access denied`)
	}

	if rs == nil {
		return ctx.StatusNotFound(`page not found`)
	}
	// js, _ := json.ParseString(rs.String(`layout`))
	// rs[`layout`] = js
	// delete(rs, `group_id`)

	return ctx.Write(rs)
}
