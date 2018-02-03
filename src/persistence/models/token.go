package models

import (
    "github.com/jinzhu/gorm"
)

type Token struct {
    gorm.Model

    UserID int
    Token string
}

func NewToken(userId int, token string) *Token {
    return &Token{
        UserID: userId,
        Token: token,
    }
}
