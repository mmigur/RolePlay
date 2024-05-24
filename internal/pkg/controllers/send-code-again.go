package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type SendCodeAgainRequest struct {
	Email string `json:"email"`
}

type SendCodeAgainResponse struct {
	Success      bool   `json:"success"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// SendCodeAgain godoc
// @Summary Отправка кода подтверждения повторно
// @Description Отправляет код подтверждения повторно на указанный email.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body SendCodeAgainRequest true "Тело запроса, содержащее email"
// @Success 200 {object} SendCodeAgainResponse "Успешный ответ"
// @Failure 400 {object} SendCodeAgainResponse "Неверный запрос"
// @Failure 500 {object} SendCodeAgainResponse "Внутренняя ошибка сервера"
// @Router /auth/send-code-again [post]
func (s *Server) SendCodeAgain(c *gin.Context) {
	var request SendCodeAgainRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, SendCodeAgainResponse{ErrorMessage: err.Error()})
		return
	}
	err := s.storage.SendCodeAgain(request.Email, *s.cfg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, SendCodeAgainResponse{ErrorMessage: err.Error()})
		return
	}
	c.JSON(http.StatusOK, SendCodeAgainResponse{Success: true})
}
