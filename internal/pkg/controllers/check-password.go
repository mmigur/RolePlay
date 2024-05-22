package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CheckPasswordRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CheckPasswordResponse struct {
	Success      bool   `json:"success"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	Token        string `json:"token,omitempty"`
}

// CheckPassword godoc
// @Summary Проверка пароля пользователя
// @Description Проверяет пароль у пользователя с указанным email
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body CheckPasswordRequest true "Запрос на проверку пользователя"
// @Success 200 {object} CheckPasswordResponse
// @Failure 400 {object} CheckPasswordResponse "Неверный запрос"
// @Failure 500 {object} CheckPasswordResponse "Ошибка сервера"
// @Router /auth/check-password [post]
func (s *Server) CheckPassword(c *gin.Context) {
	var request CheckPasswordRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, CheckPasswordResponse{Success: false, ErrorMessage: err.Error()})
		return
	}
	token, err := s.storage.CheckPassword(request.Email, request.Password, *s.cfg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CheckPasswordResponse{Success: false, ErrorMessage: err.Error()})
		return
	}

	c.JSON(http.StatusOK, CheckPasswordResponse{Success: true, Token: token})

}
