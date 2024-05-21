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
