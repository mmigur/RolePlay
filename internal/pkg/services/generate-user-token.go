package services

import (
	"RolePlayModule/internal/pkg/models"
	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	UserId   uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"nickname"`
	jwt.StandardClaims
}

func GenerateUserToken(secret []byte, user models.User) (string, error) {
	claims := UserClaims{
		UserId:   user.ID,
		Email:    user.Email,
		Username: user.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
