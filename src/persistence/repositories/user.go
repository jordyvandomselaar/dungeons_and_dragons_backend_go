package repositories

import (
    "github.com/jinzhu/gorm"
    "github.com/jordyvandomselaar/dungeons_and_dragons_backend_go/src/domain/entities"
    "github.com/jordyvandomselaar/dungeons_and_dragons_backend_go/src/persistence/models"
    "github.com/jordyvandomselaar/dungeons_and_dragons_backend_go/src/providers/database"
    "golang.org/x/crypto/bcrypt"
)

type User struct {
    Repository
}

func (u User) Create(user entities.User) error {
    password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

    usermodel := models.User{
        Email:    user.Email,
        Nickname: user.Nickname,
        Password: string(password),
    }

    database.GetDb().Create(&usermodel)

    return err
}

func NewUserRepository(db *gorm.DB) *User {
    return &User{
        Repository{
            db: db,
        },
    }
}
