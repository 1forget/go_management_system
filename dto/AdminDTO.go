package dto

import (
	"time"
)

type AdminDTO struct {
	ID        uint
	CreatedAt time.Time
	Name      string
	Password  string
	ClientIp  string
}

func NewAdminDTO(ID uint, name string, CreatedAt time.Time, password string, clientIP string) *AdminDTO {
	return &AdminDTO{
		ID:        ID,
		Name:      name,
		CreatedAt: CreatedAt,
		Password:  password,
		ClientIp:  clientIP,
	}
}
