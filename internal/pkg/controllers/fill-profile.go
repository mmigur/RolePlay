package controllers

import (
	"RolePlayModule/internal/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FillProfileRequest struct {
	Email      string `json:"email"`
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
	Address    string `json:"address"`
	Password   string `json:"password"`
}

type FillProfileResponse struct {
	Success      bool   `json:"success"`
	Token        string `json:"token,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// FillProfile godoc
// @Summary Заполнение профиля пользователя
// @Description Заполняет профиль пользователя данными из запроса.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body FillProfileRequest true "Тело запроса, содержащее данные пользователя"
// @Success 200 {object} FillProfileResponse "Успешный ответ с токеном"
// @Failure 400 {object} FillProfileResponse "Неверный запрос"
// @Failure 500 {object} FillProfileResponse "Внутренняя ошибка сервера"
// @Router /auth/fill-profile [post]
func (s *Server) FillProfile(c *gin.Context) {
	var request FillProfileRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, FillProfileResponse{Success: false, ErrorMessage: err.Error()})
		return
	}

	var user models.User
	user.Email = request.Email
	user.FirstName = request.FirstName
	user.MiddleName = request.MiddleName
	user.LastName = request.LastName
	user.Address = request.Address
	user.Password = request.Password
	token, err := s.storage.FillProfile(user, *s.cfg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, FillProfileResponse{Success: false, ErrorMessage: err.Error()})
		return
	}
	c.JSON(http.StatusOK, FillProfileResponse{Success: true, Token: token})
}
