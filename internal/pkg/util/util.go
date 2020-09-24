package util

import (
	"errors"
	"os"

	"github.com/adminboard/server/internal/pkg/constant"
	"github.com/eqto/config"
)

//FileExist ...
func FileExist(file string) bool {
	info, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

//LoadConfig ...
func LoadConfig() (*config.Config, error) {
	// var cfg *config.Config
	for _, file := range constant.ConfigLocations() {
		if FileExist(file) {
			c, e := config.ParseFile(file)
			if e != nil {
				return nil, e
			}
			return c, nil
		}
	}
	return nil, errors.New(`no config file founds`)
}
