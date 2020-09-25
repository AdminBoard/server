package app

import (
	"strings"
)

const (
	queryPage string = `SELECT p.id, p.title 
	FROM {prefix}route r 
	INNER JOIN {prefix}page p ON r.id = p.route_id 
	WHERE r.method = 'GET' AND r.path = ?`

	queryPageContent string = `SELECT pc.id, w.name, pc.sequence 
	FROM {prefix}page_content pc 
	INNER JOIN {prefix}widget w ON pc.widget_id = w.id WHERE pc.page_id = ? AND pc.sequence > 0 ORDER BY pc.sequence`
)

func getQuery(query string) string {
	prefix := cfg.GetOr(`Database.prefix`, `admin_`)
	return strings.ReplaceAll(query, `{prefix}`, prefix)
}
