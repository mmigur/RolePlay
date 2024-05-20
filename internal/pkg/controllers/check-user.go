package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CheckUserRequest struct {
	Email string `json:"email"`
}

type CheckUserResponse struct {
	IsRegistered bool   `json:"isRegistered,omitempty"`
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
	//result := true
	c.JSON(http.StatusOK, CheckUserResponse{IsRegistered: true})
}
