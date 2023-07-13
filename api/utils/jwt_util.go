package utils

import (
	"fmt"
	"log"
	"time"
	"os"

	"github.com/dgrijalva/jwt-go"
)

var (
	secretKey = []byte(os.Getenv("JWT_SECRET_KEY")) // Replace with your own secret key
)

func GenerateJWT(username string) (string, error) {
	// Create a new token object
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims (payload) of the token
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // Token expiration time 1 hour

	// Generate the JWT string
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Println(err);
		return "", err;
	}

	return tokenString, nil;
}

func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	// Verify if the token is valid
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
