package controllers

import (
	"RolePlayModule/internal/pkg/models"
	"github.com/gin-gonic/gin"
)

type GetOrdersRequest struct {
	id uint `form:"id"`
}

type OrderRecord struct {
	OrderID    uint             `json:"orderId"`
	Products   []models.Product `json:"products"`
	TotalPrice float32          `json:"totalPrice"`
	Date       string           `json:"date"`
}

type GetOrdersResponse struct {
	UserOrders []OrderRecord `json:"orders"`
}

func (s *Server) GetOrders(c *gin.Context) {

}
