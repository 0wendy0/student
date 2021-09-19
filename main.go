package main

import (
	"github.com/gin-gonic/gin"
	"student/app/router"
	"student/common/logger"
	"student/common/orm"
	"student/config"
)

func main() {
	configName := "config.json"
	config.InitConfig(configName)
	orm.InitDatabase()
	logger.InitLog()
	gin.SetMode(config.App.Mode)
	r := gin.Default()
	router.InitRouter(r)
	_ = r.Run(config.App.Host + ":" + config.App.Port)
}
