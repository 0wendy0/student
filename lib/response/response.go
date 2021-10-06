package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	adminCode "student_server/lib/response/code"
)

func AdminResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  adminCode.MsgMap[code],
	})
}
