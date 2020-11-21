package app

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/adminboard/server/internal/pkg/util"

	"github.com/adminboard/server/internal/pkg/config"

	"github.com/adminboard/server/internal/pkg/query"
	"github.com/speps/go-hashids"

	"github.com/adminboard/server/internal/pkg/api"
	apis "github.com/eqto/api-server"
	"github.com/eqto/go-json"
)

func apiRoute(ctx apis.Context) (interface{}, error) {
	command := ctx.Request().URL().RawQuery
	subcommand := ``
	if idx := strings.Index(command, `=`); idx >= 0 {
		subcommand = command[idx+1:]
		command = command[:idx]
	}
	switch command {
	case `menu`:
		return api.Menu(ctx)
	case `page`:
		return api.PageByPath(ctx, subcommand)
	case `page_id`:
		id, e := strconv.Atoi(subcommand)
		if e != nil {
			ctx.SetStatus(apis.StatusNotFound)
			return nil, e
		}
		return api.PageByID(ctx, id)
	}
	ctx.SetStatus(apis.StatusNotFound)
	return nil, errors.New(ctx.Request().URL().RawQuery)
}

func apiPublic(ctx apis.Context) (interface{}, error) {
	switch ctx.Request().URL().RawQuery {
	case `login`:
		js := ctx.Request().JSON()

		//check time
		reqTime := js.GetInt(`time`)
		t := time.Now().Unix()
		if int64(reqTime) < t-(3*60) || int64(reqTime) > t+(3*60) {
			ctx.SetStatus(apis.StatusUnauthorized)
			return nil, fmt.Errorf(`invalid time %d`, reqTime)
		}

		rs, e := ctx.Tx().Get(query.Get(query.LoginUser), js.GetString(`username`))
		if e != nil {
			ctx.SetStatus(apis.StatusUnauthorized)
			return nil, e
		} else if rs == nil {
			ctx.SetStatus(apis.StatusUnauthorized)
			return nil, errors.New(`user not found`)
		}
		userID := rs.Int(`id`)

		sig := util.Sha1(rs.String(`secret`) + js.GetString(`time`))

		if js.GetString(`signature`) != sig {
			ctx.SetStatus(apis.StatusUnauthorized)
			return nil, errors.New(`invalid password`)
		}

		res, e := ctx.Tx().Exec(query.Get(query.LoginNewSession), userID, nil)
		if e != nil {
			return nil, e
		}
		id, e := res.LastInsertID()
		if e != nil || id == 0 {
			if e == nil {
				e = errors.New(`invalid session`)
			}
			return nil, e
		}
		sess := encodeIds(id, 40)
		ctx.Tx().Exec(query.Get(query.LoginUpdateSession), sess, id)

		cookie := fmt.Sprintf(`session_id=%s; Max-Age=%d; HttpOnly`, sess, 60*15)
		ctx.Response().Header().Set(`Set-Cookie`, cookie)

		return json.Object{
			`session_id`: sess,
		}, nil
	case `session`:
		if session := getSession(ctx); session != nil {
			return session, nil
		}
		return nil, nil
	}
	ctx.SetStatus(apis.StatusNotFound)
	return nil, errors.New(ctx.Request().URL().RawQuery)
}

func encodeIds(userID int, length int) string {
	hash := hashids.NewData()
	hash.Salt = config.GetOr(`Server.salt`, ``)
	hash.MinLength = length
	h, _ := hashids.NewWithData(hash)
	enc, _ := h.Encode([]int{userID})
	return enc
}

func decodeIds(userID string, length int) (int, error) {
	hash := hashids.NewData()
	hash.Salt = config.GetOr(`Server.salt`, ``)
	hash.MinLength = length
	h, _ := hashids.NewWithData(hash)
	dec, e := h.DecodeWithError(userID)
	if e != nil {
		return 0, e
	}
	return dec[0], nil
}
