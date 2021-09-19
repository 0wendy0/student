package adminRouter

import (
	"github.com/gin-gonic/gin"
	"student/app/controller/admin"
)

func InitAdminRouter(router *gin.RouterGroup) {

	router.POST("/login", admin.Login)

	//router.Use(apiMiddlerware.CheckLogin)
	//{
	//	router.GET("/user/info", apiController.Info)
	//}

}
