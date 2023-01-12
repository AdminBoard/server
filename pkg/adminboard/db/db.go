package db

import (
	"regexp"
	"strings"
	"time"

	"github.com/adminboard/adminboard/pkg/adminboard/exception"
	"github.com/eqto/config"
	"github.com/eqto/dbm"
	_ "github.com/eqto/dbm/driver/mysql"
)

const (
	DefaultPrefix = `admin_`
)

var (
	stdCN  *dbm.Connection
	prefix = DefaultPrefix
)

func Connect() error {
	var e error
	stdCN, e = dbm.Connect(`mysql`,
		config.GetOr(`Database.hostname`, `localhost`),
		config.GetIntOr(`Database.port`, 3306),
		config.GetOr(`Database.username`, `root`),
		config.Get(`Database.password`),
		config.GetOr(`Database.name`, `adminboard`), dbm.OptionMaxLifetime(5*time.Minute))
	if e != nil {
		return exception.Wrap(e, `adminboard::Database`)
	}
	return nil
}

func CN() *dbm.Connection {
	return stdCN
}

func Prefix(table string) string {
	return prefix + table
}

func SetPrefix(str string) {
	prefix = str
}

func QueryWithPrefix(q string) string {
	regexp.MustCompile(``)
	return strings.ReplaceAll(q, `{prefix}`, prefix)
}
