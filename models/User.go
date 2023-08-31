package models

import (
	"fmt"
	"time"
)

type User struct {
	ID        uint      `json:"ID" gorm:"primary_key;type:serial"`
	CreatedAt time.Time `json:"createdAt" gorm:"type:timestamptz"`
	Name      string    `json:"name" gorm:"type:varchar(16);not null;"`
	StudentId uint      `json:"studentId" gorm:"unique" binding:"required"`
	Grade     string    `json:"grade" gorm:"type:varchar(16)"`
	Role      uint      `json:"-" gorm:"type:smallint"`
	Password  string    `json:"password" gorm:"type:varchar(128);not null"  binding:"required" `
}

func (u User) ToString() string {
	return fmt.Sprintf("User[ID: %d, CreatedAt: %s, Name: %s, StudentId: %d, Grade: %s, Password: %s]",
		u.ID, u.CreatedAt.String(), u.Name, u.StudentId, u.Grade, u.Password)
}
func (u User) IsEmpty() bool {
	return u.ID == 0 && u.CreatedAt.IsZero() && u.Name == "" && u.StudentId == 0 && u.Grade == "" && u.Password == ""
}
func (User) TableName() string {

	return "users" // 数据库表名
}
