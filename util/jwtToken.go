package util

import (
	"errors"
	"time"

	creds "github.com/LukaMijovic/role-mgmt-access-ctrl/credentials"
	"github.com/golang-jwt/jwt"
)

var secretKey *creds.SecretKey

func GenerateToken(email string, userId int64) (string, error) {
	if secretKey == nil {
		key, err := ParseJSONFile[creds.SecretKey]("credentials/secretKey.json")

		//fmt.Printf("Token: %v\n", *key)
		//fmt.Println(err.Error())

		if err != nil {
			return "", err
		}

		secretKey = key
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(secretKey.SecretKey))
}

func VerifyToken(token string) (interface{}, error) {
	if secretKey == nil {
		key, err := ParseJSONFile[creds.SecretKey]("credentials/secretKey.json")

		if err != nil {
			return nil, err
		}

		secretKey = key
	}

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Unexpected signing method")
		}

		return []byte(secretKey.SecretKey), nil
	})

	if err != nil {
		return -1, errors.New("Could not parse token")
	}

	if !parsedToken.Valid {
		return -1, errors.New("Token is not valid")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return -1, errors.New("Invalid token claims")
	}

	userId, ok := claims["userId"].(float64)

	if !ok {
		return -1, errors.New("Cannot get userId")
	}

	return int64(userId), nil
}
