package studentModel

import (
	studentModel "student_server/app/model/student"
	"time"
)

type Score struct {
	ID        uint                 `gorm:"primarykey;column:id" json:"id"`
	StudentId int                  `gorm:"column:student_id" json:"studentId"`
	Type      int                  `gorm:"column:type" json:"type"`
	CreatedAt time.Time            `gorm:"autoCreateTime;type:datetime;column:created_at" json:"createdAt"`
	UpdatedAt time.Time            `gorm:"autoCreateTime;type:datetime;column:updated_at" json:"updatedAt"`
	En        int                  `gorm:"column:en" json:"en"`
	Zh        int                  `gorm:"column:zh" json:"zh"`
	Math      int                  `gorm:"column:math" json:"math"`
	Student   studentModel.Student `gorm:"foreignkey:StudentId" json:"student"`
}
