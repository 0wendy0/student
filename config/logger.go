package config

import "github.com/spf13/viper"

type logger struct {
	Driver string
	Path   string
}

func setLogger(cfg *viper.Viper) logger {
	return logger{
		Driver: cfg.GetString("driver"),
		Path:   cfg.GetString("path"),
	}
}
