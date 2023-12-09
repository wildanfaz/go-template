package pkg

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type NewClaims struct {
	Example string
	jwt.RegisteredClaims
}

func GenerateToken(claims *NewClaims, jwtSecret []byte) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := jwtToken.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return token, err
}

func ValidateToken(token string, jwtSecret []byte) (*NewClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &NewClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := jwtToken.Claims.(*NewClaims)
	if !ok {
		return nil, errors.New("Invalid Token")
	}

	return claims, nil
}
