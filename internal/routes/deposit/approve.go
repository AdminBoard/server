package deposit

import (
	"fmt"

	"github.com/eqto/api-server"
)

func Approve(ctx api.Context) error {
	cn, e := ctx.Database()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}

	id := ctx.Request().JSON().GetInt(`id`)

	rs, e := cn.Get(`SELECT reseller_id, amount, description FROM deposit WHERE id = ?`, id)
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	if rs == nil {
		return ctx.Status(99, fmt.Sprintf(`Deposit dengan id:%d tidak ditemukan`, id))
	}
	resellerID := rs.Int(`reseller_id`)
	amount := rs.Int(`amount`)
	description := rs.String(`description`)

	groupID := ctx.Session().GetInt(`groupID`)

	rs, e = cn.Get(`SELECT * FROM deposit_permission WHERE group_id = ? AND reseller_id = ?`, groupID, resellerID)
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	if rs == nil {
		return ctx.StatusForbidden(`Akses ditolak. Tidak memiliki akses deposit untuk reseller ini`)
	}

	tx, e := ctx.Tx()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}

	userID := ctx.Session().GetInt(`userID`)
	res, e := tx.Exec(`UPDATE deposit SET status = 'verified', approved_by = ? WHERE status = 'unverified' AND id = ?`, userID, id)
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	row, e := res.RowsAffected()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	if row == 0 {
		return ctx.StatusBadRequest(`Deposit sedang/sudah approval, atau tidak ditemukan`)
	}

	res, e = tx.Exec(`UPDATE reseller SET balance = balance + ? WHERE id = ? AND status = 'active'`, amount, resellerID)
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	row, e = res.RowsAffected()
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	if row == 0 {
		return ctx.StatusBadRequest(`Tidak dapat menambah balance reseller, pastikan status reseller aktif`)
	}
	rs, e = tx.Get(`SELECT balance FROM reseller WHERE id = ?`, resellerID)
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}
	if rs == nil {
		return ctx.StatusInternalServerError(`Tidak dapat mengakses data reseller, mohon coba lagi.`)
	}
	balance := rs.Int(`balance`)

	res, e = tx.Insert(`transaction`, map[string]interface{}{
		`partner_id`:   resellerID,
		`ref_id`:       id,
		`type`:         2000, //Tipe deposit
		`partner_type`: `reseller`,
		`amount`:       amount,
		`balance`:      balance,
		`description`:  description,
	})
	if e != nil {
		return ctx.StatusInternalServerError(e.Error())
	}

	return nil
}
