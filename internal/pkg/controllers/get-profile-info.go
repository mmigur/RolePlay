package controllers

import (
	"RolePlayModule/internal/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetProfileInfoRequest struct {
	Id int `form:"id"`
}
type GetProfileInfoResponse struct {
	ErrorMessage string `form:"errorMessage,omitempty"`
	Username     string `form:"username"`
	Email        string `form:"email"`
	FirstName    string `form:"firstName"`
	MiddleName   string `form:"middleName"`
	LastName     string `form:"lastName"`
	Address      string `form:"address"`
}

// GetProfileInfo godoc
// @Summary Получить информацию о профиле
// @Description Метод для получения информации о профиле пользователя
// @Tags Profile
// @Accept json
// @Produce json
// @Param id query int true "Идентификатор пользователя"
// @Success 200 {object} GetProfileInfoResponse
// @Failure 401 {object} GetProfileInfoResponse
// @Failure 500 {object} GetProfileInfoResponse
// @Router /profile [get]
func (s *Server) GetProfileInfo(c *gin.Context) {
	claims, err := services.GetUserClaimsFromJWT(c, *s.cfg)
	if err != nil {
		c.JSON(http.StatusUnauthorized, GetProfileInfoResponse{ErrorMessage: err.Error()})
		return
	}
	user, err := s.storage.GetProfileInfo(claims.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, GetProfileInfoResponse{ErrorMessage: err.Error()})
		return
	}

	c.JSON(http.StatusOK, GetProfileInfoResponse{
		Username:   user.Username,
		Email:      user.Email,
		FirstName:  user.FirstName,
		MiddleName: user.MiddleName,
		LastName:   user.LastName,
		Address:    user.Address,
	})
}
