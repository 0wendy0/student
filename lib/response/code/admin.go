package adminCode

const (
	Success = 1000

	TokenInvalid = 2000
	TokenExpired = 2001

	ParamsError   = 3000
	PasswordError = 3001

	CreateError = 4000

	CodeExit  = 5000
	ScoreExit = 5001
)

var MsgMap = map[int]string{
	Success:       "成功",
	TokenInvalid:  "无效的登录凭证Token",
	TokenExpired:  "失效的登录凭证Token",
	ParamsError:   "参数错误",
	PasswordError: "密码错误",
	CreateError:   "创建失败",
	CodeExit:      "学号已存在",
	ScoreExit:     "成绩已录入",
}
