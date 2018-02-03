package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model

	Nickname, Email, Password string
}

func NewUserModel(
	nickname string,
	email string,
	password string,
) User {
	return User{
		Nickname: nickname,
		Email:    email,
		Password: password,
	}
}
