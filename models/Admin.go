package models

import "time"

type Admin struct {
	ID        uint      `json:"-" gorm:"primary_key;type:serial"`
	Name      string    `json:"name"gorm:"type:varchar(20);not null;" binding:"required"`
	Password  string    `json:"password"gorm:"type:varchar(128)" binding:"required"`
	CreatedAt time.Time `json:"createdAt" gorm:"type:timestamptz"`
}

func (admin Admin) IsEmpty() bool {
	return admin.ID == 0 && admin.CreatedAt.IsZero() && admin.Name == "" && admin.Password == ""
}
func (Admin) TableName() string {
	return "admins" // 数据库表名
}
