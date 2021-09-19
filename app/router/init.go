package router

import (
	"net/http"
	"student/app/middlerware"
	adminRouter "student/app/router/admin"

	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {

	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404")
	})

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})

	router.StaticFile("/favicon.ico", "./favicon.ico")

	router.Use(middlerware.Cors)
	{
		admin := router.Group("/admin")
		{
			adminRouter.InitAdminRouter(admin)
		}
	}
}
