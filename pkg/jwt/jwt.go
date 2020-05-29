package jwt

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// SecretKey being used to sign our tokens.
var SecretKey = []byte("secret")

// GenerateToken generates a jwt token and assign a username to it's claims and return it.
func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)
	/* Set token claims */
	claims["user"] = username
	claims["exp"] = time.Now().Add(15 * time.Minute).Unix()

	tokenStr, err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatal("Error generating token, ", err)
		return "", err
	}

	return tokenStr, nil
}

// ParseToken parses a jwt token and returns the username in it's claims.
func ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["user"].(string), nil
	}
	return "", err
}
