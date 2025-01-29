package utils

import (
	"crypto/ecdsa"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var privateKey *ecdsa.PrivateKey

func GenrateToken(userId uint) (string, error) {
	// Load the private key from PEM file
    privateKeyData, err := os.ReadFile("private_key.pem")
    if err != nil {
        log.Fatalf("Error reading private key: %v", err)
    }

    // Parse the ECDSA private key
    privateKey, err := jwt.ParseECPrivateKeyFromPEM(privateKeyData)
    if err != nil {
        log.Fatalf("Error parsing ECDSA private key: %v", err)
    }
	
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})
	// log.Println(jwtSecret)

	return token.SignedString(privateKey)
}

func VerifyToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return &privateKey.PublicKey, nil
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
