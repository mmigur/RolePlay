package controllers

import (
	"RolePlayModule/internal/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CheckCodeRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type CheckCodeResponse struct {
	IsCorrect    bool   `json:"isCorrect"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func (s *Server) CheckCode(c *gin.Context) {
	var request CheckCodeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, CheckCodeResponse{ErrorMessage: err.Error()})
		return
	}
	isValid := services.ValidEmail(request.Email)
	if !isValid {
		c.JSON(http.StatusBadRequest, CheckCodeResponse{ErrorMessage: "invalid email"})
		return
	}

	result, err := s.storage.CheckCode(request.Email, request.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CheckCodeResponse{ErrorMessage: err.Error()})
		return
	}

	c.JSON(http.StatusOK, CheckCodeResponse{IsCorrect: result})
}
