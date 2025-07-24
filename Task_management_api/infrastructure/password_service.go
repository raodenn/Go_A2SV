package infrastructure

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordSvc struct{}

func NewPasswordSvc() *PasswordSvc {
	return &PasswordSvc{}
}

func (p *PasswordSvc) HashPassword(password string) *string {
	pass, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	hashed := string(pass)
	return &hashed
}

func (p *PasswordSvc) VerifyPassword(userPass string, foundPass string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(foundPass), []byte(userPass))

	if err != nil {
		return false, err
	}
	return true, nil
}
