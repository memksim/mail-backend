package model

type Avatar string

type User struct {
	Email     string  `json:"email"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Avatar    *Avatar `json:"avatar"`
}
