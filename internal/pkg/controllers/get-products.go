package controllers

import (
	"RolePlayModule/internal/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetProductsRequest struct{}

type Category struct {
	Name     string                `json:"name"`
	Products []models.ShortProduct `json:"products"`
}

type GetProductsResponse struct {
	Categories   []Category `json:"categories"`
	ErrorMessage string     `json:"errorMessage,omitempty"`
}

func (s *Server) GetProducts(c *gin.Context) {
	categories, err := s.storage.GetCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get categories"})
		return
	}
	responseCategories := make([]Category, len(categories))
	for i, category := range categories {
		responseCategories[i].Name = category.Name

		products, err := s.storage.GetProductsByCategory(category.Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, GetProductsResponse{ErrorMessage: err.Error()})
			return
		}
		responseCategories[i].Products = make([]models.ShortProduct, len(products))
		for j, product := range products {
			responseCategories[i].Products[j] = models.ShortProduct{
				Id:       product.Id,
				Name:     product.Name,
				ImageUrl: product.ImageUrl,
				OldPrice: product.OldPrice,
				Price:    product.Price,
			}
		}
	}

	c.JSON(http.StatusOK, GetProductsResponse{Categories: responseCategories})
}
