package deposit

import (
	"strconv"
	"strings"

	"github.com/eqto/api-server"
)

func Add(ctx api.Context) error {
	js := ctx.Request().JSON()

	resellerID := js.GetInt(`reseller_id`)
	amount, e := strconv.Atoi(strings.ReplaceAll(js.GetString(`amount`), `.`, ``))
	if e != nil || amount == 0 {
		return ctx.StatusBadRequest(`invalid 'amount'`)
	}

	if resellerID == 0 {
		return ctx.StatusBadRequest(`invalid 'reseller_id'`)
	}

	cn, e := ctx.Database()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	userID := ctx.Session().GetInt(`userID`)

	res, e := cn.Insert(`deposit`, map[string]interface{}{
		`reseller_id`: resellerID,
		`amount`:      amount,
		`added_by`:    userID,
		`status`:      `unverified`,
		`description`: js.GetString(`description`),
	})
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	count, e := res.RowsAffected()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	if count == 0 {
		return ctx.StatusInternalServerError(`Gagal menambah deposit`)
	}
	return nil
}
