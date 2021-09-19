package config

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

type database struct {
	Driver      string
	Url         string
	MaxIdle     int
	MaxOpen     int
	MaxLifetime time.Duration
}

func setDatabase(cfg *viper.Viper) database {
	driver := cfg.GetString("driver")
	var url string
	switch driver {
	case "mysql":
		mysqlCfg := cfg.Sub("mysql")
		if mysqlCfg == nil {
			log.Fatal("配置文件中mysql配置不存在")
		}
		url = setMysql(mysqlCfg)
	default:
		log.Fatal("不支持的数据库类型")
	}
	return database{
		Driver:      driver,
		Url:         url,
		MaxIdle:     cfg.GetInt("maxIdle"),
		MaxOpen:     cfg.GetInt("maxOpen"),
		MaxLifetime: cfg.GetDuration("maxLifetime"),
	}
}

func setMysql(cfg *viper.Viper) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=PRC",
		cfg.GetString("user"),
		cfg.GetString("password"),
		cfg.GetString("host"),
		cfg.GetString("port"),
		cfg.GetString("db"))
}
