package main

import (
	"github.com/gin-gonic/gin"
	"student_server/app/router"
	"student_server/common/logger"
	"student_server/common/orm"
	"student_server/config"
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
