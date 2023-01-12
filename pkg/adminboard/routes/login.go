package routes

import (
	"time"

	"github.com/adminboard/adminboard/pkg/adminboard/crypto"
	"github.com/adminboard/adminboard/pkg/adminboard/db"
	"github.com/eqto/api-server"
	"github.com/eqto/config"
	"github.com/eqto/go-json"
)

func Login(ctx api.Context) error {
	js := ctx.Request().JSON()
	rs, e := db.CN().Get(db.QueryWithPrefix(`SELECT * FROM {prefix}user WHERE username = ?`), js.GetString(`username`))
	if e != nil {
		return ctx.StatusForbidden(e.Error())
	}
	if rs == nil {
		return ctx.StatusForbidden(`Username not found`)
	}

	signature := crypto.MD5(rs.String(`secret`) + js.GetString(`time`))

	if signature != js.GetString(`signature`) {
		return ctx.StatusForbidden(`Invalid password`)
	}
	header := ctx.Request().Header()
	client := json.Object{
		`Host`:       header.Get(`Host`),
		`User-Agent`: header.Get(`User-Agent`),
	}
	res, e := db.CN().Exec(db.QueryWithPrefix(`INSERT INTO {prefix}user_session(user_id, client) VALUES(?, ?)`), rs.Int(`id`), client.ToString())
	if e != nil {
		return ctx.StatusInternalServerError(`invalid session: ` + e.Error())
	}
	sessionID, e := res.LastInsertID()
	if e != nil {
		return ctx.StatusInternalServerError(`invalid session: ` + e.Error())
	}
	id := crypto.EncodeIDWithSalt(sessionID, config.Get(`Security.salt`))

	ctx.Response().Header().SetCookie(`SESS`, id, 15*time.Minute)
	return nil
}
