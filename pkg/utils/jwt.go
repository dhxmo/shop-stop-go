package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/copier"
)

const (
	TokenExpiredTime = 3600
)

func GenerateToken(payload interface{}) (string, error) {
	tokenContent := jwt.MapClaims{
		"payload": payload,
		"expiry":  time.Now().Add(time.Second * TokenExpiredTime).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}

func ValidateToken(jwtToken string) (*map[string]interface{}, error) {
	if jwtToken == "" {
		return nil, fmt.Errorf("token must not be empty")
	}
	cleanJWT := strings.Replace(jwtToken, "Bearer ", "", -1)
	tokenData := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(cleanJWT, tokenData, func(token *jwt.Token) (interface{}, error) {
		return []byte("TokenPassword"), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("token is invalid")
	}

	var data map[string]interface{}
	copier.Copy(&data, tokenData["payload"])
	return &data, nil
}
