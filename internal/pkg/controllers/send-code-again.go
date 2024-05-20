package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type SendCodeAgainRequest struct {
	Email string `json:"email"`
}

type SendCodeAgainResponse struct {
	Success      bool   `json:"success,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func (s *Server) SendCodeAgain(c *gin.Context) {
	var request SendCodeAgainRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, SendCodeAgainResponse{ErrorMessage: err.Error()})
		return
	}

	c.JSON(http.StatusOK, SendCodeAgainResponse{Success: true})
}
