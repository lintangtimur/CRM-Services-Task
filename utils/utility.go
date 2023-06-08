package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type IAuth interface {
	HashPassword(pass string) (string, error)
	VerifyPassword(hashedPassword string, candidatePassword string) error
}
type Utils struct{}

func (u Utils) HashPassword(pass string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("could not hash password %w", err)
	}
	return string(hashedPassword), nil
}
func (u Utils) VerifyPassword(hashedPassword string, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}
