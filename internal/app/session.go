package app

import (
	"fmt"
	"strings"

	"github.com/adminboard/server/internal/pkg/query"
	"github.com/eqto/api-server"
	"github.com/eqto/go-json"
)

func getSession(ctx api.Context) json.Object {
	cookies := strings.Split(ctx.Request().Header().Get(`Cookie`), `;`)

	for _, cookie := range cookies {
		if strings.HasPrefix(cookie, `session_id=`) {
			split := strings.Split(cookie, `=`)
			if len(split) == 2 {
				rs, e := ctx.Tx().Get(query.Get(query.Session), split[1])
				if e == nil && rs != nil {
					setCookie(ctx.Response().Header(), split[1])
					return json.Object{
						`session_id`: split[1],
						`user_id`:    rs.Int(`id`),
						`role_id`:    rs.Int(`role_id`),
					}
				}
			}
		}
	}
	setCookie(ctx.Response().Header(), ``)
	return nil
}

func setCookie(header *api.ResponseHeader, sessionID string) {
	age := 60 * 15 //5 minutes
	if sessionID == `` {
		age = 0
	}
	cookie := fmt.Sprintf(`session_id=%s; Max-Age=%d; Path=/; HttpOnly`, sessionID, age)
	header.Set(`Set-Cookie`, cookie)
}
