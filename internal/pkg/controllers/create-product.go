package controllers

import (
	"RolePlayModule/internal/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateProductRequest struct {
	Name              string  `json:"name"`
	Description       string  `json:"description"`
	CategoryId        uint    `json:"category"`
	Price             float32 `json:"price"`
	Weight            float32 `json:"weight"`
	ShelfLife         string  `json:"shelfLife"`
	StorageConditions string  `json:"storageConditions"`
	Brand             string  `json:"brand"`
	Country           string  `json:"country"`
}

type CreateProductResponse struct {
	Success      bool   `json:"success,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func (s *Server) CreateProduct(c *gin.Context) {
	var request CreateProductRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, CreateProductResponse{ErrorMessage: err.Error()})
		return
	}
	var newProduct = models.Product{
		Name:              request.Name,
		Description:       request.Description,
		CategoryId:        request.CategoryId,
		Price:             request.Price,
		Weight:            request.Weight,
		ShelfLife:         request.ShelfLife,
		StorageConditions: request.StorageConditions,
		Brand:             request.Brand,
		Country:           request.Country,
	}
	err := s.storage.CreateProduct(newProduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CreateProductResponse{ErrorMessage: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, CreateProductResponse{Success: true})
}
