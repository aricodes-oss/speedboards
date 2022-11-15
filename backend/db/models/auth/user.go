package auth

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name         string
	PasswordHash string `json:"-"`

	Email         string
	EmailVerified bool
}
