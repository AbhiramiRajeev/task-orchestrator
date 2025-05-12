package model


type UserRequest struct {
	Email string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}


type User struct {
	ID int `json:"id"`
	Email string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}