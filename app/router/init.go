package router

import (
	"net/http"
	"student_server/app/middlerware"
	adminRouter "student_server/app/router/admin"

	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {

	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404")
	})

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "计科201 20201131142 干宏琴")
	})

	router.Use(middlerware.Cors)
	{
		admin := router.Group("/admin")
		{
			adminRouter.InitAdminRouter(admin)
		}
	}
}
