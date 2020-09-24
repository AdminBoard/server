package constant

import "log"

//ConfigLocations ...
func ConfigLocations() []string {
	log.Println(`locations`)
	return []string{
		`./adminboard.conf`,
		`configs/adminboard.conf`,
		`/etc/adminboard.conf`,
	}
}
