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

// ParseDIYToMap is used to convert the json object to map[string]string
func ParseDIYToMap(cf config.Configer, name string) (m map[string]string, err error) {
	DIY, err := cf.DIY(name)
	if err != nil {
		return
	}
	itf, ok := DIY.(map[string]interface{})
	if !ok {
		err = errors.New("value is not map[string]interface{}")
		return
	}
	m = make(map[string]string)
	for k, v := range itf {
		m[k] = v.(string)
	}
	return
}

// ParseDIYToMap is used to convert the json object to map[string]map[string]string
func ParseDIYToMaps(cf config.Configer, name string) (m map[string]map[string]string, err error) {
	DIY, err := cf.DIY(name)
	if err != nil {
		return
	}
	itf, ok := DIY.(map[string]interface{})
	if !ok {
		err = errors.New("value is not map[string]interface{}")
		return
	}
	m = make(map[string]map[string]string)

	for k, v := range itf {
		mTmp := v.(map[string]interface{})
		for k2, v2 := range mTmp {
			if _, exist := m[k]; !exist {
				m[k] = make(map[string]string)
			}
			m[k][k2] = v2.(string)
		}
	}
	return
}
