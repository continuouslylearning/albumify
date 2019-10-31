package users

import (
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	hash, e := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(hash), e
}

func validatePassword(hash string, password string) bool {
	e := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return e == nil
}
