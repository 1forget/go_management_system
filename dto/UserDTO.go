package dto

import "time"

type UserDTO struct {
	ID        uint
	CreatedAt time.Time
	Name      string
	StudentId uint
	Grade     string
	Role      uint
	ClientIp  string
}

func NewUserDTO(ID uint, name string, role uint, studentID uint, clientIP string) *UserDTO {
	return &UserDTO{
		ID:        ID,
		Name:      name,
		Role:      role,
		StudentId: studentID,
		ClientIp:  clientIP,
	}
}
