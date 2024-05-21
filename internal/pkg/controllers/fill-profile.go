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
	Username   string `json:"username"`
	Password   string `json:"password"`
}

type FillProfileResponse struct {
	Success      bool   `json:"success"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

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
	user.Username = request.Username
	user.Password = request.Password
	err := s.storage.FillProfile(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, FillProfileResponse{Success: false, ErrorMessage: err.Error()})
		return
	}
	c.JSON(http.StatusOK, FillProfileResponse{Success: true})
}
