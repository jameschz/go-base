package config

import (
	"github.com/jameschz/go-base/lib/util"
	"github.com/spf13/viper"
)

var (
	_configMap = make(map[string]*viper.Viper)
	_configDir = ""
)

// Init :
func Init() {
	// init config dir
	if len(_configDir) == 0 {
		_configDir = util.GetRootPath() + "/etc/" + util.GetEnv()
	}
}

// Load :
func Load(configName string) *viper.Viper {
	// init
	Init()
	// cache config
	if _, exist := _configMap[configName]; !exist {
		_configMap[configName] = viper.New()
		_configMap[configName].SetConfigType("yaml")
		_configMap[configName].AddConfigPath(_configDir)
		_configMap[configName].SetConfigName(configName)
		if err := _configMap[configName].ReadInConfig(); err != nil {
			panic(err)
		}
	}
	// return config
	return _configMap[configName]
}
