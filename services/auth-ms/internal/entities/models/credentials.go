package models

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Credential struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
