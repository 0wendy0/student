package config

import "github.com/spf13/viper"

type app struct {
	Host string
	Port string
	Name string
	Mode string
	Url  string
}

func setApp(cfg *viper.Viper) app {
	return app{
		Host: cfg.GetString("host"),
		Port: cfg.GetString("port"),
		Name: cfg.GetString("name"),
		Mode: cfg.GetString("mode"),
		Url:  cfg.GetString("url"),
	}
}
