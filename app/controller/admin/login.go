package admin

import (
	"github.com/gin-gonic/gin"
	"student/lib/response"
	adminCode "student/lib/response/code"
)

func Login(c *gin.Context) {
	response.AdminResponse(c, adminCode.SUCCESS, nil)
}
