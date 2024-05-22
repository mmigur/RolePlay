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

// CreateCategory godoc
// @Summary Создание категории
// @Description Создает новую категорию с указанным именем.
// @Tags Category
// @Accept json
// @Produce json
// @Param request body CreateCategoryRequest true "Тело запроса, содержащее имя категории"
// @Success 200 {object} CreateCategoryResponse "Успешный ответ"
// @Failure 400 {object} CreateCategoryResponse "Неверный запрос"
// @Failure 500 {object} CreateCategoryResponse "Внутренняя ошибка сервера"
// @Router /category [post]
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
