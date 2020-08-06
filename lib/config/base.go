package config

import (
	"go-base/lib/util"

	"github.com/spf13/viper"
)

var (
	_config_map = make(map[string]*viper.Viper)
	_config_dir = ""
)

func Init() {
	// init config dir
	if len(_config_dir) == 0 {
		_config_dir = util.GetRootPath() + "/etc/" + util.GetEnv()
	}
}

func Load(configName string) *viper.Viper {
	// init
	Init()
	// cache config
	if _, exist := _config_map[configName]; !exist {
		_config_map[configName] = viper.New()
		_config_map[configName].SetConfigType("yaml")
		_config_map[configName].AddConfigPath(_config_dir)
		_config_map[configName].SetConfigName(configName)
		if err := _config_map[configName].ReadInConfig(); err != nil {
			panic(err)
		}
	}
	// return config
	return _config_map[configName]
}
