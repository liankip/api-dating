package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func VerifyJWT(tokenStr string) (uint, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return 0, err
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := uint(claims["userID"].(float64))

	return userID, nil
}
