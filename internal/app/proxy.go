package app

import "github.com/adminboard/server/internal/pkg/query"

func loadProxy() error {
	proxies, e := svr.Database().Select(query.Get(query.Proxy))
	if e != nil {
		return e
	}
	for _, proxy := range proxies {
		svr.Proxy(proxy.String(`path`), proxy.String(`destination`))
	}

	return nil
}
