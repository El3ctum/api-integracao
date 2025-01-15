package auth

// import (
// 	"log"
// 	"os"

// 	"github.com/golang-jwt/jwt/v5"
// 	_ "github.com/joho/godotenv/autoload"
// )

// var (
// 	err          error
// 	key          []byte
// 	token        *jwt.Token
// 	signedString string
// )

// func GenerateJwtToken(user User) (string, error) {
// 	key = []byte(os.Getenv("SECRET_KEY"))
// 	token = jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
// 		"sub": "davi",
// 	})
// 	signedString, err = token.SignedString(key)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
