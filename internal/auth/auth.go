package auth

import (
	"api-integracao/internal/models"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	_ "github.com/joho/godotenv/autoload"
)

type TokenGenerator interface {
	GenerateJwtToken(user models.UserMetadata) (string, error)
}

type MyCustomClaims struct {
	Data map[string]string
	jwt.RegisteredClaims
}

func GenerateJwtToken(user models.UserMetadata) (string, error) {
	key := []byte(os.Getenv("SECRET_KEY"))

	claims := MyCustomClaims{
		Data: map[string]string{
			"id":          user.ID,
			"name":        user.Name,
			"departments": strings.Join(user.Departments, "|"), // Used '|' to avoid issues with commas
			"roles":       strings.Join(user.Roles, "|"),
			"permissions": strings.Join(user.Permissions, "|"),
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "api-sicredi",
			Subject:   user.Name,
			ID:        user.ID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // Changed to HS256 for simplicity with byte key
	signedString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return signedString, nil
}

func ValidateJwtToken(tokenString string) (*MyCustomClaims, error) {
	key := []byte(os.Getenv("SECRET_KEY"))

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return nil, fmt.Errorf("an error occurred while validating the token: %v", err)
	}

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token claims")
}
