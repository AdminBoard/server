package session

import (
	enc "encoding/json"
	"errors"
	"fmt"

	"github.com/eqto/api-server"
	"github.com/eqto/config"
	"github.com/eqto/go-json"
	"github.com/golang-jwt/jwt"
)

type Token struct {
	SessionID int    `json:"sid"`
	UserID    int    `json:"uid"`
	Username  string `json:"uname"`
	GroupID   int    `json:"gid"`
	Expired   int    `json:"expired"`
	IsValid   bool
}

func Get(ctx api.Context) (*Token, error) {
	exist := false
	t := &Token{IsValid: false}
	defer func() {
		if !exist {
			sess := ctx.Session()
			sess.Put(`session`, t)
			sess.Put(`user_id`, t.UserID)
			sess.Put(`group_id`, t.GroupID)
			sess.Put(`session_id`, t.SessionID)
		}
	}()

	if t := ctx.Session().Get(`session`); t != nil {
		exist = true
		return t.(*Token), nil
	}

	tokenStr := ctx.Request().Header().Cookie(`adminboard`)
	if tokenStr == `` {
		return t, errors.New(`invalid token`)
	}
	token, e := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(config.Get(`Security.secret`)), nil
	})
	if e != nil {
		return t, errors.New(e.Error())
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		js, e := json.ParseObject(claims)
		if e != nil {
			return t, errors.New(e.Error())
		}
		enc.Unmarshal(js.ToBytes(), t)
		t.IsValid = true
		return t, nil
	}
	return t, errors.New(`invalid token`)
}
