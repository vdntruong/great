package dto

type CreateUserDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,password"`
}

type UpdateUserDTO struct {
	Username string `json:"username" validate:"required,username"`
}

type UserDTO struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}
