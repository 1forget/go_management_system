package bean

type UserLoginRequest struct {
	StudentId int    `json:"studentId,omitempty"`
	Password  string `json:"password,omitempty"`
}

type AdminLoginRequest struct {
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserCreateRequest struct {
	StudentID uint    `json:"studentId,omitempty"`
	Name      *string `json:"name,omitempty"`
	Password  string  `json:"password,omitempty"`
	Grade     *string `json:"grade,omitempty"`
}
type AdminUpdateUserRequest struct {
	ID        uint    `json:"ID" binding:"required"`
	StudentID uint    `json:"studentId,omitempty"`
	Name      *string `json:"name,omitempty"`
	Password  string  `json:"password,omitempty"`
	Grade     *string `json:"grade,omitempty"`
}

type UpdateUserInfoRequest struct {
	Name        *string `json:"name,omitempty"`
	Password    string  `json:"newPassword,omitempty"`
	OldPassword string  `json:"oldPassword,omitempty"`
	Grade       *string `json:"grade,omitempty"`
}

type UserDeleteRequest struct {
	Password string `json:"password,omitempty"`
}
