package routes

import (
	"github.com/adminboard/adminboard/pkg/adminboard/crypto"
	"github.com/eqto/api-server"
	"github.com/eqto/config"
)

func Session(ctx api.Context) error {
	// cn, e := ctx.Database()
	// if e != nil {
	// 	return ctx.StatusInternalServerError(e.Error())
	// }

	_, e := GetSessionID(ctx)
	if e != nil {
		return e
	}
	// stmt := dbm.Select(`u.password_expired_at`).From(db.Prefix(`user_session us`)).
	// 	InnerJoin(db.Prefix(`user u`), `us.user_id = u.id`).Where(`us.id = ?`)

	// row, e := cn.Get(cn.SQL(stmt), sessionID)
	// if e != nil {
	// 	return ctx.StatusInternalServerError(e.Error())
	// }
	// expired := row.TimeNil(`password_expired_at`)

	// if expired == nil {
	// 	return ctx.Error(1, `Expired password`)
	// }

	return nil
}

func GetSessionID(ctx api.Context) (int, error) {
	session := ctx.Request().Header().Cookie(`SESS`)
	if session == `` {
		return 0, ctx.StatusUnauthorized(`invalid session`)
	}
	sessionID, e := crypto.DecodeIDWithSalt(session, config.Get(`Security.salt`))
	if e != nil {
		return 0, ctx.StatusUnauthorized(e.Error())
	}
	return sessionID, nil
}
