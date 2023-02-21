package params

import (
	"errors"
	"strings"

	"github.com/eqto/api-server"
	"github.com/eqto/go-json"
)

/*
Fetch get params from JSON body. Mandatory and optional is separated by comma (,).
Return error when one of fields from mandatory does not exist.
*/
func Fetch(ctx api.Context, mandatory string, optional string) (json.Object, error) {
	split := strings.Split(mandatory, `,`)

	jsResp := json.Object{}

	jsReq := ctx.Request().JSON()
	for _, str := range split {
		s := strings.TrimSpace(str)
		if s == `` {
			continue
		}
		if !jsReq.Has(s) {
			return nil, errors.New(`invalid required parameter: ` + s)
		}
		jsResp.Put(s, jsReq.Get(s))
	}
	split = strings.Split(optional, `,`)
	for _, str := range split {
		s := strings.TrimSpace(str)
		if s == `` {
			continue
		}
		jsResp.Put(s, jsReq.Get(s))
	}

	return jsResp, nil
}

/*
FetchQuery get params from query string. Mandatory and optional is separated by comma (,).
Return error when one of fields from mandatory does not exist.
*/
func FetchQuery(ctx api.Context, mandatory string, optional string) (json.Object, error) {
	split := strings.Split(mandatory, `,`)

	jsResp := json.Object{}

	for _, str := range split {
		s := strings.TrimSpace(str)
		if s == `` {
			continue
		}
		val := ctx.Request().QueryParam(s)
		if val == `` {
			return nil, errors.New(`invalid required query parameter: ` + s)
		}
		jsResp.Put(s, val)
	}
	split = strings.Split(optional, `,`)
	for _, str := range split {
		s := strings.TrimSpace(str)
		if s == `` {
			continue
		}
		if ctx.Request().URL().Query().Has(s) {
			jsResp.Put(s, ctx.Request().QueryParam(s))
		}
	}

	return jsResp, nil
}
