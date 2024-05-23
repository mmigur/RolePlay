package controllers

import (
	"RolePlayModule/internal/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateOrderRequest struct {
	ProductsIds map[uint]int `json:"productsIds"`
}

type CreateOrderResponse struct {
	Success      bool   `json:"success"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// CreateOrder godoc
// @Summary Создать заказ
// @Description Метод для создания заказа пользователя
// @Tags Orders
// @Accept json
// @Produce json
// @Param productsIds body map[uint]int true "Идентификаторы и количество товаров в заказе"
// @Success 200 {object} CreateOrderResponse
// @Failure 400 {object} CreateOrderResponse
// @Failure 401 {object} CreateOrderResponse
// @Failure 500 {object} CreateOrderResponse
// @Router /orders [post]
func (s *Server) CreateOrder(c *gin.Context) {
	claims, err := services.GetUserClaimsFromJWT(c, *s.cfg)
	if err != nil {
		c.JSON(http.StatusUnauthorized, GetProfileInfoResponse{ErrorMessage: err.Error()})
		return
	}
	var request CreateOrderRequest
	if err = c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, GetProfileInfoResponse{ErrorMessage: err.Error()})
		return
	}
	err = s.storage.CreateOrder(request.ProductsIds, claims.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, GetProfileInfoResponse{ErrorMessage: err.Error()})
		return
	}
	c.JSON(http.StatusOK, CreateOrderResponse{Success: true})

}
