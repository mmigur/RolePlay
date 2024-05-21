package controllers

import (
	"RolePlayModule/internal/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetCategoriesResponse struct {
	Categories []models.Category `json:"categories"`
}

func (s *Server) GetCategories(c *gin.Context) {
	categories, err := s.storage.GetCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, GetCategoriesResponse{Categories: categories})
}
