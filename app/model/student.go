package model

import (
	"time"
)

// Admin [...]
type Admin struct {
	ID         uint      `gorm:"primaryKey;column:id;type:int unsigned;not null"`
	Username   string    `gorm:"column:username;type:varchar(20);not null"`
	Password   string    `gorm:"column:password;type:char(32);not null"`
	Salt       string    `gorm:"column:salt;type:char(32);not null"`
	Createtime time.Time `gorm:"column:createtime;type:timestamp"`
	Updatetime time.Time `gorm:"column:updatetime;type:timestamp"`
}
