package jwt_utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

/*
return example

	{
		"username": "fulan",
		"user_id": "1234",
		"role": "owner",
		"session_id": "1234",
		"exp": 12345,
	}
*/
func GenerateJwtToken(username string, user_id string, role string, session_id string, secret_key string, exp int) (string, error) {
	secretKey := []byte(secret_key)

	claims := jwt.MapClaims{
		"user_id":    user_id,
		"username":   username,
		"role":       role,
		"exp":        time.Now().Add(time.Hour * time.Duration(exp)).Unix(),
		"session_id": session_id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func ValidateJWT(token_string string, secret_key string) (jwt.MapClaims, error) {
	var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
	var JWT_SIGNATURE_KEY = []byte(secret_key)

	token, err := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != JWT_SIGNING_METHOD {
			return nil, fmt.Errorf("signing method invalid")
		}

		return JWT_SIGNATURE_KEY, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	claims := token.Claims.(jwt.MapClaims)
	return claims, nil
}
