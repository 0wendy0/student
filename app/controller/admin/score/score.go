package score

import (
	"github.com/gin-gonic/gin"
	scoreModel "student_server/app/model/score"
	"student_server/common/orm"
	"student_server/lib/response"
	adminCode "student_server/lib/response/code"
)

type get struct {
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"pageSize" form:"pageSize"`
	Type     int    `json:"type" form:"type"`
	Name     string `json:"name" form:"name"`
	Code     string `json:"code" form:"code"`
}

type listRes struct {
	List  *[]scoreModel.Score `json:"list"`
	Total int64               `json:"total"`
}

func List(c *gin.Context) {
	param := get{}
	err := c.ShouldBindQuery(&param)
	if err != nil {
		response.AdminResponse(c, adminCode.ParamsError, nil)
		return
	}
	list := new([]scoreModel.Score)
	var total int64
	cDb := orm.Db.Model(&scoreModel.Score{}).Joins("Student")
	lDb := orm.Db.Model(&scoreModel.Score{}).Joins("Student")
	if param.Type != 0 {
		cDb.Where("type = ?", param.Type)
		lDb.Where("type = ?", param.Type)
	}
	if param.Name != "" {
		cDb.Where("Student.name LIKE ?", "%"+param.Name+"%")
		lDb.Where("Student.name LIKE ?", "%"+param.Name+"%")
	}
	if param.Code != "" {
		cDb.Where("Student.code LIKE ?", "%"+param.Code+"%")
		lDb.Where("Student.code LIKE ?", "%"+param.Code+"%")
	}
	cDb.Count(&total)
	if param.PageSize != 0 && param.Page != 0 {
		lDb.Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize)
	}
	lDb.Order("scores.id desc").Find(&list)
	res := listRes{
		List:  list,
		Total: total,
	}
	response.AdminResponse(c, adminCode.Success, res)

}

type post struct {
	StudentId int `json:"studentId" binding:"required"`
	Type      int `json:"type" binding:"required"`
	En        int `json:"en"`
	Zh        int `json:"zh"`
	Math      int `json:"math"`
}

func Add(c *gin.Context) {
	param := post{}
	err := c.ShouldBind(&param)
	if err != nil {
		response.AdminResponse(c, adminCode.ParamsError, nil)
		return
	}
	var score scoreModel.Score
	err = orm.Db.Model(&scoreModel.Score{}).Where("student_id = ?", param.StudentId).Where("type = ?", param.Type).First(&score).Error
	if err == nil {
		response.AdminResponse(c, adminCode.ScoreExit, nil)
		return
	}
	score = scoreModel.Score{StudentId: param.StudentId, Type: param.Type, En: param.En, Zh: param.Zh, Math: param.Math}
	orm.Db.Create(&score)
	response.AdminResponse(c, adminCode.Success, nil)
}

type put struct {
	Id   int `json:"id" binding:"required"`
	En   int `json:"en"`
	Zh   int `json:"zh"`
	Math int `json:"math"`
}

func Edit(c *gin.Context) {
	param := put{}
	err := c.ShouldBind(&param)
	if err != nil {
		response.AdminResponse(c, adminCode.ParamsError, nil)
		return
	}
	var score scoreModel.Score
	err = orm.Db.Model(&scoreModel.Score{}).Where("id = ?", param.Id).First(&score).Error
	if err != nil {
		response.AdminResponse(c, adminCode.ParamsError, nil)
		return
	}
	score.En = param.En
	score.Zh = param.Zh
	score.Math = param.Math
	orm.Db.Save(&score)
	response.AdminResponse(c, adminCode.Success, nil)
}

type del struct {
	Id int `json:"id" binding:"required"`
}

func Delete(c *gin.Context) {
	param := del{}
	err := c.ShouldBind(&param)
	if err != nil {
		response.AdminResponse(c, adminCode.ParamsError, nil)
		return
	}
	orm.Db.Delete(&scoreModel.Score{}, param.Id)
	response.AdminResponse(c, adminCode.Success, nil)
}
