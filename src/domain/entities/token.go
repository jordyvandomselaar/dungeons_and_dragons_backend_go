package entities

type Token struct {
	UserId int
	Token  string
}

func NewToken(userId int, token string) *Token {
	return &Token{
		UserId: userId,
		Token:  token,
	}
}
