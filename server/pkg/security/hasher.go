package security

import (
	"golang.org/x/crypto/bcrypt"
)

type Hasher interface {
	HashPassword(password string) (string, error)
	VerifyPassword(inputPassword, hashedPassword string) (bool, error)
}

type hasher struct{}

func NewHasher() *hasher {
	return &hasher{}
}

func (h *hasher) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (h *hasher) VerifyPassword(inputPassword, hasherdPassword string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hasherdPassword), []byte(inputPassword)); err != nil {
		return false, err
	}

	return true, nil
}
