package login

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	userModel "student_server/app/model/user"
	"student_server/common/orm"
	"student_server/lib/jwtToken"
	"student_server/lib/response"
	adminCode "student_server/lib/response/code"
)

type Form struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Token struct {
	Token string `json:"token"`
}

func Login(c *gin.Context) {
	form := &Form{}
	err := c.ShouldBind(form)
	if err != nil {
		response.AdminResponse(c, adminCode.ParamsError, nil)
		return
	}
	form.Password = fmt.Sprintf("%x", md5.Sum([]byte(form.Password)))
	user, err := userModel.GetUserByUserName(form.Username)
	if err != nil {
		// 账户不存在
		user = userModel.User{Username: form.Username, Password: form.Password}
		orm.Db.Create(&user)
	} else {
		// 账户存在
		if form.Password != user.Password {
			response.AdminResponse(c, adminCode.PasswordError, nil)
			return
		}
	}
	jwt := jwtToken.NewJWT()
	data := jwtToken.CustomClaims{}
	data.Uid = user.ID
	res, _ := jwt.CreateToken(data)
	token := Token{res}
	response.AdminResponse(c, adminCode.Success, token)
}

func Info(c *gin.Context) {
	current, _ := c.Get("current")
	response.AdminResponse(c, adminCode.Success, current)
}
