package controllers

import (
	"RolePlayModule/internal/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type GetOrdersRequest struct {
	Id uint `form:"id"`
}

type ProductRecord struct {
	ProductID    uint    `json:"productId"`
	Name         string  `json:"name"`
	Price        float32 `json:"price"`
	ProductCount int     `json:"productCount"`
}

type OrderRecord struct {
	OrderID    uint            `json:"orderId"`
	Products   []ProductRecord `json:"products"`
	TotalPrice float32         `json:"totalPrice"`
	Date       string          `json:"date"`
}

type GetOrdersResponse struct {
	ErrorMessage string        `json:"errorMessage,omitempty"`
	UserOrders   []OrderRecord `json:"orders"`
}

// GetOrders godoc
// @Summary      Получить заказы пользователя
// @Description  Получить список заказов для аутентифицированного пользователя
// @Tags         Orders
// @Produce      json
// @Success      200  {object}  GetOrdersResponse
// @Failure      401  {object}  GetOrdersResponse
// @Failure      500  {object}  GetOrdersResponse
// @Router       /orders [get]
func (s *Server) GetOrders(c *gin.Context) {
	claims, err := services.GetUserClaimsFromJWT(c, *s.cfg)
	if err != nil {
		c.JSON(http.StatusUnauthorized, GetOrdersResponse{ErrorMessage: err.Error()})
		return
	}
	orders, err := s.storage.GetOrders(claims.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, GetOrdersResponse{ErrorMessage: err.Error()})
		return
	}
	var orderRecords []OrderRecord
	for _, order := range orders {
		var productRecords []ProductRecord
		for _, detail := range order.OrderDetails {
			productRecords = append(productRecords, ProductRecord{
				ProductID:    detail.ProductID,
				Name:         detail.Product.Name,
				Price:        detail.Product.Price,
				ProductCount: detail.ProductCount,
			})
		}
		orderRecords = append(orderRecords, OrderRecord{
			OrderID:    order.ID,
			Products:   productRecords,
			TotalPrice: order.TotalPrice,
			Date:       order.CreatedAt.Format(time.RFC3339),
		})
	}

	c.JSON(http.StatusOK, GetOrdersResponse{UserOrders: orderRecords})
}
