package entities

type User struct {
	Nickname, Email, Password string
}

func NewUserEntity(
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