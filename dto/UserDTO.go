package dto

import "time"

type UserDTO struct {
	ID        uint
	CreatedAt time.Time
	Name      string
	StudentId uint
	Grade     string
	ClientIp  string
}

func NewUserDTO(ID uint, name string, CreatedAt time.Time, studentID uint, grade string, clientIP string) *UserDTO {
	return &UserDTO{
		ID:        ID,
		Name:      name,
		CreatedAt: CreatedAt,
		StudentId: studentID,
		Grade:     grade,
		ClientIp:  clientIP,
	}
}
