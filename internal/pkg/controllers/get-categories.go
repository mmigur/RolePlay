package controllers

import (
	"RolePlayModule/internal/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetCategoriesResponse struct {
	Categories []models.Category `json:"categories"`
}

// GetCategories godoc
// @Summary Получение списка категорий
// @Description Возвращает список всех категорий.
// @Tags Category
// @Produce json
// @Success 200 {object} GetCategoriesResponse "Успешный ответ с списком категорий"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /category [get]
func (s *Server) GetCategories(c *gin.Context) {
	categories, err := s.storage.GetCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, GetCategoriesResponse{Categories: categories})
}
