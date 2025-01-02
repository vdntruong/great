package dtos

type CreateUserDTO struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}

type UpdateUserDTO struct {
	Username string `json:"username" validate:"required,username"`
}

type UserDTO struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type LoginReq struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}

type LoginRes struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
