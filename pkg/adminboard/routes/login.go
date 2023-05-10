package routes

import (
	"time"

	"github.com/adminboard/adminboard/pkg/adminboard/crypto"
	"github.com/adminboard/adminboard/pkg/adminboard/db"
	"github.com/eqto/api-server"
	"github.com/eqto/config"
	"github.com/eqto/dbm"
	"github.com/eqto/go-json"
	"github.com/golang-jwt/jwt"
)

func Login(ctx api.Context) error {
	js := ctx.Request().JSON()
	stmt := dbm.Select(`id, group_id, secret, TIMEDIFF(expired_at, NOW()) < 0 AS expired`).
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
	claims := jwt.MapClaims{
		`s`: sessionID,
		`u`: userID,
		`g`: rs.Int(`group_id`),
	}
	if rs.Int(`expired`) == 1 {
		claims[`e`] = 1
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, e := token.SignedString([]byte(config.Get(`Security.secret`)))
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}

	ctx.Response().Header().SetCookie(`adminboard`, tokenStr, 30*time.Minute)

	return nil
}
