package user

import (
	"github.com/adminboard/adminboard/pkg/adminboard/crypto"
	"github.com/adminboard/adminboard/pkg/adminboard/db"
	"github.com/adminboard/adminboard/pkg/adminboard/session"
	"github.com/eqto/api-server"
	"github.com/eqto/dbm"
)

func ChangePassword(ctx api.Context) error {
	token, e := session.Get(ctx)
	if e != nil {
		return ctx.StatusForbidden(e.Error())
	}
	js := ctx.Request().JSON()

	password, confirm := js.GetString(`password`), js.GetString(`confirm`)

	if len(password) < 6 {
		return ctx.Error(1000, `Invalid password, minimum length: 6.`)
	} else if password != confirm {
		return ctx.Error(1000, `The password confirmation does not match.`)
	}

	cn, e := ctx.Database()

	stmt := dbm.Update(db.Prefix(`user`)).Set(`secret = ?, expired_at = DATE_ADD(NOW(), INTERVAL 90 DAY)`).Where(`id = ?`)
	tx, e := ctx.Tx()
	secret := crypto.MD5(token.Username + password)
	if _, e = tx.Exec(cn.SQL(stmt), secret, token.UserID); e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	setCookie(ctx, token.SessionID, token.UserID, token.GroupID, token.Username, false)
	return nil
}
