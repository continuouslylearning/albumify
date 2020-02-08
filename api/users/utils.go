package users

import (
	"errors"
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func createToken(user *User) (string, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user.Normalize(),
		"exp":  time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	signedToken, e := token.SignedString(secret)
	return signedToken, e
}

func hashPassword(password string) (string, error) {
	hash, e := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(hash), e
}

func verifyToken(signedToken string) (jwt.MapClaims, error) {
	token, e := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		secret := []byte(os.Getenv("JWT_SECRET"))
		return secret, nil
	})
	if e != nil && !token.Valid {
		e = errors.New("the provided token was not valid")
	}

	return token.Claims.(jwt.MapClaims), e
}

func verifyPassword(hash string, password string) bool {
	e := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if e != nil {
		return false
	}
	return true
}
