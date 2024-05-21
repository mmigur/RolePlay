package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateCategoryRequest struct {
	Name string `json:"name"`
}

type CreateCategoryResponse struct {
	Success      bool   `json:"success,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func (s *Server) CreateCategory(c *gin.Context) {
	var request CreateCategoryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, CreateCategoryResponse{ErrorMessage: err.Error()})
		return
	}
	err := s.storage.CreateCategory(request.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CreateCategoryResponse{ErrorMessage: err.Error()})
		return
	}

	c.JSON(http.StatusOK, CreateCategoryResponse{Success: true})
}
