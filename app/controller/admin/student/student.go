package student

import (
	"github.com/gin-gonic/gin"
	scoreModel "student_server/app/model/score"
	studentModel "student_server/app/model/student"
	"student_server/common/orm"
	"student_server/lib/response"
	adminCode "student_server/lib/response/code"
)

type get struct {
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"pageSize" form:"pageSize"`
	Name     string `json:"name" form:"name"`
	Code     string `json:"code" form:"code"`
}

type listRes struct {
	List  *[]studentModel.Student `json:"list"`
	Total int64                   `json:"total"`
}

func List(c *gin.Context) {
	param := get{}
	err := c.ShouldBindQuery(&param)
	if err != nil {
		response.AdminResponse(c, adminCode.ParamsError, nil)
		return
	}
	list := new([]studentModel.Student)
	var total int64
	cDb := orm.Db.Model(&studentModel.Student{})
	lDb := orm.Db.Model(&studentModel.Student{})
	if param.Name != "" {
		cDb.Where("name LIKE ?", "%"+param.Name+"%")
		lDb.Where("name LIKE ?", "%"+param.Name+"%")
	}
	if param.Code != "" {
		cDb.Where("code LIKE ?", "%"+param.Code+"%")
		lDb.Where("code LIKE ?", "%"+param.Code+"%")
	}
	cDb.Count(&total)
	if param.PageSize != 0 && param.Page != 0 {
		lDb.Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize)
	}
	lDb.Order("id desc").Find(&list)
	res := listRes{
		List:  list,
		Total: total,
	}
	response.AdminResponse(c, adminCode.Success, res)

}

type post struct {
	Name string `json:"name" binding:"required"`
	Code string `json:"code" binding:"required"`
	Sex  int    `json:"sex" binding:"required"`
}

func Add(c *gin.Context) {
	param := post{}
	err := c.ShouldBind(&param)
	if err != nil {
		response.AdminResponse(c, adminCode.ParamsError, nil)
		return
	}
	var student studentModel.Student
	err = orm.Db.Model(&studentModel.Student{}).Where("code = ?", param.Code).First(&student).Error
	if err == nil {
		response.AdminResponse(c, adminCode.CodeExit, nil)
		return
	}
	student = studentModel.Student{Name: param.Name, Code: param.Code, Sex: param.Sex}
	orm.Db.Create(&student)
	response.AdminResponse(c, adminCode.Success, nil)
}

type put struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Code string `json:"code" binding:"required"`
	Sex  int    `json:"sex" binding:"required"`
}

func Edit(c *gin.Context) {
	param := put{}
	err := c.ShouldBind(&param)
	if err != nil {
		response.AdminResponse(c, adminCode.ParamsError, nil)
		return
	}
	var student studentModel.Student
	err = orm.Db.Model(&studentModel.Student{}).Where("id = ?", param.Id).First(&student).Error
	if err != nil {
		response.AdminResponse(c, adminCode.ParamsError, nil)
		return
	}
	if param.Code != student.Code {
		var s studentModel.Student
		err = orm.Db.Model(&studentModel.Student{}).Where("code = ?", param.Code).First(&s).Error
		if err == nil {
			response.AdminResponse(c, adminCode.CodeExit, nil)
			return
		}
	}
	student.Code = param.Code
	student.Name = param.Name
	student.Sex = param.Sex
	orm.Db.Save(&student)
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
	orm.Db.Delete(&studentModel.Student{}, param.Id)
	orm.Db.Where("student_id = ?", param.Id).Delete(&scoreModel.Score{})
	response.AdminResponse(c, adminCode.Success, nil)
}
