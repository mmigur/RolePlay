package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetProductByIdRequest struct {
	Id uint `form:"id"`
}
type GetProductByIdResponse struct {
	ErrorMessage      string  `json:"errorMessage,omitempty"`
	Id                uint    `json:"id,omitempty"`
	ImageUrl          string  `json:"imageUrl,omitempty"`
	Name              string  `json:"name,omitempty"`
	Description       string  `json:"description,omitempty"`
	CategoryId        uint    `json:"categoryId,omitempty"`
	OldPrice          float32 `json:"oldPrice,omitempty"`
	Price             float32 `json:"price,omitempty"`
	Weight            float32 `json:"weight,omitempty"`
	ShelfLife         string  `json:"shelfLife,omitempty"`
	StorageConditions string  `json:"storageConditions,omitempty"`
	Brand             string  `json:"brand,omitempty"`
	Country           string  `json:"country,omitempty"`
}

// GetProductById godoc
// @Summary Получить информацию о продукте по идентификатору
// @Description Возвращает информацию о продукте по идентификатору
// @Tags Product
// @Accept json
// @Produce json
// @Param id query uint true "Идентификатор продукта"
// @Success 200 {object} GetProductByIdResponse
// @Failure 400 {object} GetProductByIdResponse
// @Failure 500 {object} GetProductByIdResponse
// @Router /product/id [get]
func (s *Server) GetProductById(c *gin.Context) {
	var request GetProductByIdRequest
	err := c.ShouldBindQuery(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, GetProductByIdResponse{ErrorMessage: err.Error()})
		return
	}
	product, err := s.storage.GetProductById(request.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, GetProductByIdResponse{ErrorMessage: err.Error()})
		return
	}
	c.JSON(http.StatusOK, GetProductByIdResponse{
		Id:                product.Id,
		ImageUrl:          product.ImageUrl,
		Name:              product.Name,
		Description:       product.Description,
		CategoryId:        product.CategoryId,
		OldPrice:          product.OldPrice,
		Price:             product.Price,
		Weight:            product.Weight,
		ShelfLife:         product.ShelfLife,
		StorageConditions: product.StorageConditions,
		Brand:             product.Brand,
		Country:           product.Country,
	})
}
