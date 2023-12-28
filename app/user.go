package app

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
    ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
    Username  string     `json:"username" grom:"required"`
    Email     string     `json:"email" grom:"required,unique"`
    Password  string     `json:"password" gorm:"required,min=6"`
    CreatedAt time.Time  `json:"created_at"`
    UpdatedAt time.Time  `json:"updated_at"`
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func VerifyPassword(hashedPassword, inputPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
}