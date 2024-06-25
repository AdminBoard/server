package user

import (
	"github.com/adminboard/server/pkg/adminboard/crypto"
	"github.com/adminboard/server/pkg/adminboard/db"
	"github.com/eqto/api-server"
	"github.com/eqto/dbm"
	"github.com/eqto/go-json"
)

func Login(ctx api.Context) error {
	js := ctx.Request().JSON()
	stmt := dbm.Select(`id, username, group_id, secret, TIMEDIFF(expired_at, NOW()) < 0 AS expired`).
		From(db.Prefix(`user`)).Where(`username = ?`)
	cn, e := ctx.Database()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	rs, e := cn.Get(cn.SQL(stmt), js.GetString(`username`))
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
	userID := rs.Int(`id`)
	stmtInsert := dbm.InsertInto(db.Prefix(`user_session`), `user_id, client`).Values(`?, ?`)
	res, e := db.CN().Exec(cn.SQL(stmtInsert), userID, client.ToString())
	if e != nil {
		return ctx.StatusInternalServerError(`invalid session: ` + e.Error())
	}
	sessionID, e := res.LastInsertID()
	if e != nil {
		return ctx.StatusInternalServerError(`invalid session: ` + e.Error())
	}

	expired := rs.Int(`expired`) == 1

	if e := setCookie(ctx, sessionID, userID, rs.Int(`group_id`), rs.String(`username`), expired); e != nil {
		return ctx.StatusUnauthorized(e.Error())
	}

	if expired {
		return ctx.Error(1, `password expired`)
	}

	return nil
}
