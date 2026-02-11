package models

type RegisterInput struct {
	Username string
	Email    string
	Password string
}

type LoginInput struct {
	Email    string
	Password string
}
