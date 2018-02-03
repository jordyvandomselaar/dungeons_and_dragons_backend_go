package jwt

import "github.com/jordyvandomselaar/dungeons_and_dragons_backend_go/src/domain/entities"

type TokenRepositoryInterface interface {
    Create(token entities.Token)
    Find(token entities.Token)
}
