package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(password string) (string, error) {
	hashedPassword, errHashedPassword := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if errHashedPassword != nil {
		return string(hashedPassword), errHashedPassword
	}

	return string(hashedPassword), nil
}

func ComparePassword(hashedPassword []byte, password []byte) error {
	compareHashPassword := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if compareHashPassword != nil {
		return compareHashPassword
	}

	return nil
}
