// Package beeconfig is used to load config file under conf for beego
package beeconfig

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
)

const defaultAdapter = "json"

var configMap map[string]config.Configer

// Load return Configer
func Load(configName string, adapterName ...string) (cf config.Configer, err error) {

	adapter := defaultAdapter
	if len(adapterName) > 0 {
		adapter = adapterName[0]
	}

	adapterFile := fmt.Sprintf("%s.%s", configName, adapter)
	if c, ok := configMap[adapterFile]; ok {
		cf = c
		return
	}

	configFile := filepath.Join("conf", beego.BConfig.RunMode, adapterFile)

	if !fileExist(configFile) {
		configFile = filepath.Join("conf", adapterFile)
		if !fileExist(configFile) {
			err = errors.New("file not exist:" + configFile)
			return
		}
	}

	cf, err = config.NewConfig(adapter, configFile)
	if err != nil {
		configMap[adapterFile] = cf
	}

	return
}

func fileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
