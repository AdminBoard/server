package query

import (
	"strings"

	"github.com/adminboard/server/internal/pkg/config"
)

const (
	Route = `SELECT r.*
	FROM {prefix}route r
	LEFT JOIN {prefix}page p ON r.id = p.route_id
	WHERE p.route_id IS NULL ORDER BY path`

	Page = `SELECT p.id, p.title 
	FROM {prefix}route r 
	INNER JOIN {prefix}page p ON r.id = p.route_id 
	WHERE r.method = 'GET' AND r.path = ?`

	PageContent = `SELECT pw.id, pw.params, pw.sequence, w.name, r.path AS data_source 
	FROM {prefix}page_widget pw 
	INNER JOIN {prefix}widget w ON pw.widget_id = w.id 
	LEFT JOIN {prefix}route r ON pw.data_route_id = r.id
	WHERE pw.page_id = ? AND pw.sequence > 0 
	ORDER BY pw.sequence`

	Action = `SELECT * FROM {prefix}action WHERE route_id = ? AND sequence > 0 ORDER BY sequence`

	Menu = `SELECT m.id, m.parent_id, m.kind, r.path, m.caption, m.description, m.sequence 
	FROM {prefix}menu m 
	INNER JOIN {prefix}menu_role mr ON m.id = mr.menu_id 
	LEFT JOIN {prefix}page p ON m.page_id = p.id
	LEFT JOIN {prefix}route r ON p.route_id = r.id
	WHERE m.sequence > 0 AND mr.role_id = ?
	ORDER BY m.parent_id, m.sequence`
)

//Get ...
func Get(query string) string {
	prefix := config.GetOr(`Database.prefix`, `admin_`)
	return strings.ReplaceAll(query, `{prefix}`, prefix)
}
