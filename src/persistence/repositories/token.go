package repositories

import (
    "github.com/jordyvandomselaar/dungeons_and_dragons_backend_go/src/domain/entities"
)

type Token struct {
    Repository
}

func (t Token) Create(token entities.Token) {
    //tokenModel := models.NewToken(token.UserId)
}

func (t Token) Find(token entities.Token) {
    panic("implement me")
}
