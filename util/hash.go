package util

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 16)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func CheckPassword(password string, hashedPass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password)) == nil
}