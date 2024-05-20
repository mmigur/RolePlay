package services

import (
	"RolePlayModule/internal/utils/config"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func DecodeJWT(tokenString string, cfg config.Config) (*UserClaims, error) {

	secretKey := []byte(cfg.SecretKey)

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	}

	claims := &UserClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, keyFunc)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}
