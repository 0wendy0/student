package studentModel

import "time"

type Student struct {
	ID        uint      `gorm:"primarykey;column:id" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Code      string    `gorm:"column:code" json:"code"`
	Sex       int       `gorm:"column:sex" json:"sex"`
	CreatedAt time.Time `gorm:"autoCreateTime;type:datetime;column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoCreateTime;type:datetime;column:updated_at" json:"updatedAt"`
}
