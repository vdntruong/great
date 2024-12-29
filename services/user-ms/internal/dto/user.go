package dto

// CreateUserReq represents for client request to register new user
type CreateUserReq struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,password"`
}

// UpdateUserReq represents for client request to update user
type UpdateUserReq struct {
	Username string `json:"username" validate:"required,username"`
}

// UserRes represents for a user (aka user brief) response to client
type UserRes struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}
