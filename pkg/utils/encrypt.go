package utils

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("12a3f4b5c6d7e8f90123456789abcdef123456789abcdef01234567890abcdef")

func GenrateToken(userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})
	log.Println(jwtSecret)

	return token.SignedString(jwtSecret)
}

func VerifyToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
}

func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPassword), nil
}

func CheckPassword(hashPassword string, planPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(planPassword))
}
