package constant

import "log"

//ConfigLocations ...
func ConfigLocations() []string {
	log.Println(`locations`)
	return []string{
		`./adminboard.cnf`,
		`configs/adminboard.cnf`,
		`/etc/adminboard.cnf`,
	}
}
