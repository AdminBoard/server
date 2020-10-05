package app

import (
	"fmt"
	"strings"

	"github.com/adminboard/server/internal/pkg/query"
	"github.com/eqto/api-server"
	"github.com/eqto/go-db"
	"github.com/eqto/go-json"
)

func getSession(tx *db.Tx, respHeader api.Header, rawCookie string) json.Object {
	cookies := strings.Split(rawCookie, `;`)

	for _, cookie := range cookies {
		if strings.HasPrefix(cookie, `session_id=`) {
			split := strings.Split(cookie, `=`)
			if len(split) == 2 {
				rs, e := tx.Get(query.Get(query.Session), split[1])
				if e == nil && rs != nil {
					setCookie(respHeader, split[1])
					return json.Object{
						`session_id`: split[1],
						`user_id`:    rs.Int(`id`),
						`role_id`:    rs.Int(`role_id`),
					}
				}
			}
		}
	}
	setCookie(respHeader, ``)
	return nil
}

func setCookie(header api.Header, sessionID string) {
	if header == nil {
		return
	}
	age := 60 * 5 //5 minutes
	if sessionID == `` {
		age = 0
	}
	cookie := fmt.Sprintf(`session_id=%s; Max-Age=%d; Path=/; HttpOnly`, sessionID, age)
	header.Add(`Set-Cookie`, cookie)
}
