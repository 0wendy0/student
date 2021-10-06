package middlerware

import (
	"github.com/gin-gonic/gin"
	userModel "student_server/app/model/user"
	"student_server/lib/jwtToken"
	"student_server/lib/response"
	adminCode "student_server/lib/response/code"
)

func CheckLogin(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	l := len(authorization)
	if l < 50 {
		response.AdminResponse(c, adminCode.TokenInvalid, nil)
		c.Abort()
		return
	}
	jwt := jwtToken.NewJWT()
	res, err := jwt.ParseToken(authorization)
	if err == jwtToken.TokenExpired {
		response.AdminResponse(c, adminCode.TokenExpired, nil)
		c.Abort()
		return
	} else if err != nil {
		response.AdminResponse(c, adminCode.TokenInvalid, nil)
		c.Abort()
		return
	}
	current, err := userModel.GetUserById(res.Uid)
	if err != nil {
		response.AdminResponse(c, adminCode.TokenInvalid, nil)
		c.Abort()
		return
	}
	c.Set("uid", res.Uid)
	c.Set("current", current)
	c.Next()
}
