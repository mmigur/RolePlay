package services

import (
	"RolePlayModule/internal/utils/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func GetUserClaimsFromJWT(c *gin.Context, cfg config.Config) (*UserClaims, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return nil, &gin.Error{
			Err:  http.ErrNoCookie,
			Type: gin.ErrorTypePublic,
		}
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
