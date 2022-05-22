package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func GenHashedPassword(password string) (string, error) {
	pwdBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(pwdBytes), nil
}
