package config

import (
	"github.com/spf13/viper"
	"log"
)

var App app
var Database database
var Logger logger

func InitConfig(path string) {
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("读取配置文件失败")
	}

	cfgApp := viper.Sub("app")
	if cfgApp == nil {
		log.Fatal("配置文件中app配置不存在")
	}
	App = setApp(cfgApp)

	cfgDatabase := viper.Sub("database")
	if cfgDatabase == nil {
		log.Fatal("配置文件中database配置不存在")
	}
	Database = setDatabase(cfgDatabase)

	cfgLogger := viper.Sub("logger")
	if cfgLogger == nil {
		log.Fatal("配置文件中logger配置不存在")
	}
	Logger = setLogger(cfgLogger)
}
