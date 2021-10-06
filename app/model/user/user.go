package userModel

import (
	"gorm.io/gorm"
	"student_server/common/orm"
)

type User struct {
	gorm.Model
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
}

func (User) TableName() string {
	return "users"
}

func GetUserByUserName(username string) (User, error) {
	user := User{}
	err := orm.Db.Where("username = ?", username).First(&user).Error
	return user, err
}

func GetUserById(id uint) (User, error) {
	user := User{}
	err := orm.Db.Where("id = ?", id).First(&user).Error
	return user, err
}

func (u User) Create() error {
	return orm.Db.Create(&u).Error
}
