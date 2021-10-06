package config

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
)

type database struct {
	Driver      string
	Url         string
	Dialect     gorm.Dialector
	MaxIdle     int
	MaxOpen     int
	MaxLifetime time.Duration
}

func setDatabase(cfg *viper.Viper) database {
	driver := cfg.GetString("driver")
	var url string
	var dialect gorm.Dialector
	switch driver {
	case "mysql":
		cfg := cfg.Sub("mysql")
		if cfg == nil {
			log.Fatal("配置文件中mysql配置不存在")
		}
		url = setMysql(cfg)
		dialect = mysql.Open(url)
	case "sqlite3":
		cfg = cfg.Sub("sqlite3")
		if cfg == nil {
			log.Fatal("配置文件中sqlite3配置不存在")
		}
		url = cfg.GetString("path")
		dialect = sqlite.Open(url)
	default:
		log.Fatal("不支持的数据库类型")
	}
	return database{
		Driver:      driver,
		Dialect:     dialect,
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
