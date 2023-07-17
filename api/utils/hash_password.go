package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(input string) (string, error) {
	inputBytes := []byte(input)

	hashedBytes, err := bcrypt.GenerateFromPassword(inputBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hashedString := string(hashedBytes)

	return hashedString, nil
}
