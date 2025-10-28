package utils

import (
	"comeht/models"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.HashedPassword), 10)
	if err != nil {
		return err
	}
	user.HashedPassword = string(hashedPassword)
	return nil
}

func CheckHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

var (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
)

func GenerateToken(userId int) (string, error) {

	secret := os.Getenv(JWT_SECRET_KEY)
	exp := time.Now().Add(24 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     exp,
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func RemoveBearer(tokenString *string) {
	if strings.HasPrefix(*tokenString, "Bearer") {
		*tokenString = strings.TrimPrefix(*tokenString, "Bearer ")
	}
}

func VerifyToken(tokenString string) (float64, error) {

	secret := os.Getenv(JWT_SECRET_KEY)

	RemoveBearer(&tokenString)
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}
		return nil, fmt.Errorf("invalid token")
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}
	return claims["user_id"].(float64), nil
}
