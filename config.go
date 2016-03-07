// Package beeconfig is used to load config file under conf for beego
package beeconfig

import (
	"errors"
	"path/filepath"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/utils"
)

var (
	configFiles map[string]config.Configer
)

func init() {
	configFiles = make(map[string]config.Configer)
}

// Load return Configer
// you can load the specific config file by run mode
func Load(configFile string) (cf config.Configer, err error) {
	if c, ok := configFiles[configFile]; ok {
		cf = c
		return
	}
	fullPathFile := filepath.Join("conf", beego.BConfig.RunMode, configFile)
	if !utils.FileExists(fullPathFile) {
		fullPathFile = filepath.Join("conf", configFile)
		if !utils.FileExists(fullPathFile) {
			err = errors.New(fullPathFile + " not found")
			return
		}
	}
	adapter := strings.TrimLeft(filepath.Ext(fullPathFile), ".")
	cf, err = config.NewConfig(adapter, fullPathFile)
	if err != nil {
		configFiles[configFile] = cf
	}
	return
}
