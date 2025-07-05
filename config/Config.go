package config

import "mail/model"

type currentUser model.User

var CurrentUser = currentUser{
	Email:     "admin@example.com",
	FirstName: "Admin",
	LastName:  "Adminovich",
	Avatar:    nil,
}
