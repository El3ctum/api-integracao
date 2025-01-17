package auth

import (
	"api-integracao/internal/models"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	_ "github.com/joho/godotenv/autoload"
)

type MyCustomClaims struct {
	Data map[string]string
	jwt.RegisteredClaims
}

func GenerateJwtToken(user models.User) (string, error) {
	key := []byte(os.Getenv("SECRET_KEY"))

	claims := MyCustomClaims{
		Data: map[string]string{
			"nome":        user.FirstName,
			"departments": strings.Join(user.Departments, ","),
			"roles":       strings.Join(user.Roles, ","),       // Transforms the MAP into string
			"permissions": strings.Join(user.Permissions, ","), //
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "api-sicredi",
			Subject:   user.FirstName,
			ID:        user.ID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, &claims)
	signedString, err := token.SignedString(key)

	if err != nil {
		return "", nil
	}

	return signedString, nil
}

func ValidateJwtToken(token string) bool {
	jwtValidator := jwt.NewParser(jwt.WithValidMethods([]string{jwt.SigningMethodES256.Name}))

	jwtValidator.Parse(token, jwt.Keyfunc{})
}
