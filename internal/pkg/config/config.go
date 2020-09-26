package config

import (
	"errors"

	"github.com/adminboard/server/internal/pkg/constant"
	"github.com/adminboard/server/internal/pkg/util"
	"github.com/eqto/config"
)

var cfg *config.Config

//Load ...
func Load() error {
	for _, file := range constant.ConfigLocations() {
		if util.FileExist(file) {
			c, e := config.ParseFile(file)
			if e != nil {
				return e
			}
			cfg = c
			return nil
		}
	}
	return errors.New(`no config file founds`)
}

//Get ...
func Get(key string) string {
	return cfg.Get(key)
}

//GetOr ...
func GetOr(key, def string) string {
	return cfg.GetOr(key, def)
}

//GetIntOr ...
func GetIntOr(key string, def int) int {
	return cfg.GetIntOr(key, def)
}
