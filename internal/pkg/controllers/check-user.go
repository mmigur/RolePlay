package controllers

import (
	"RolePlayModule/internal/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CheckUserRequest struct {
	Email string `json:"email"`
}

type CheckUserResponse struct {
	IsRegistered bool   `json:"isRegistered"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// CheckUser godoc
//
//	@Summary		Проверка на зарегистрирован ли пользователь
//	@Description	Проверяет существование пользователя и в случае его отсутсвия, отправляет код на email
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		CheckUserRequest	true	"Запрос на проверку пользователя"
//	@Success		200		{object}	CheckUserResponse
//	@Router			/auth/check-user [post]
func (s *Server) CheckUser(c *gin.Context) {
	var request CheckUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, CheckUserResponse{ErrorMessage: err.Error()})
		return
	}
	isValid := services.ValidEmail(request.Email)
	if !isValid {
		c.JSON(http.StatusBadRequest, CheckUserResponse{ErrorMessage: "invalid email"})
		return
	}
	result, err := s.storage.CheckUser(request.Email, *s.cfg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CheckUserResponse{ErrorMessage: err.Error()})
		return
	}
	if result {
		c.JSON(http.StatusOK, CheckUserResponse{IsRegistered: true})
		return
	}
	c.JSON(http.StatusOK, CheckUserResponse{IsRegistered: false})
}
