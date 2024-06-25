package user

import (
	"time"

	"github.com/eqto/api-server"
	"github.com/eqto/config"
	"github.com/golang-jwt/jwt"
)

func setCookie(ctx api.Context, sessionID, userID, groupID int, username string, expired bool) error {
	claims := jwt.MapClaims{
		`sid`:   sessionID,
		`uid`:   userID,
		`uname`: username,
		`gid`:   groupID,
	}
	if expired {
		claims[`expired`] = expired
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, e := token.SignedString([]byte(config.Get(`Security.secret`)))
	if e != nil {
		return e
	}
	ctx.Response().Header().SetCookie(`adminboard`, tokenStr, 30*time.Minute)
	return nil
}
