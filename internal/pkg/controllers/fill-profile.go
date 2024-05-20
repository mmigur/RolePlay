package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type FillProfileRequest struct {
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
	c.JSON(http.StatusOK, FillProfileResponse{Success: true})
}
