package adminRouter

import (
	"github.com/gin-gonic/gin"
	"student_server/app/controller/admin/login"
	"student_server/app/controller/admin/score"
	"student_server/app/controller/admin/student"
	"student_server/app/middlerware"
)

func InitAdminRouter(router *gin.RouterGroup) {

	router.POST("/login", login.Login)

	router.Use(middlerware.CheckLogin)
	{
		router.GET("/user/info", login.Info)

		router.GET("/student", student.List)
		router.POST("/student", student.Add)
		router.PUT("/student", student.Edit)
		router.DELETE("/student", student.Delete)

		router.GET("/score", score.List)
		router.POST("/score", score.Add)
		router.PUT("/score", score.Edit)
		router.DELETE("/score", score.Delete)
	}
}
