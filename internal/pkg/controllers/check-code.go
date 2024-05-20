package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CheckCodeRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type CheckCodeResponse struct {
	IsCorrect    bool   `json:"isCorrect,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func (s *Server) CheckCode(c *gin.Context) {
	var request CheckCodeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, CheckCodeResponse{ErrorMessage: err.Error()})
		return
	}

	result := true

	c.JSON(http.StatusOK, CheckCodeResponse{IsCorrect: result})
}
