package controller

import (
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	ordSvc OrderService
}

func InitOrderHandler(r *gin.Engine, ordSvc OrderService) {
	handler := OrderHandler{ordSvc: ordSvc}

	r.POST("/api/orders/", handler.CreateOrder)
	r.GET("/api/orders/:order_id", handler.GetOrder)
	r.GET("/api/orders/", handler.GetAllOrders)

	r.PATCH("/api/orders/", handler.UpdateOrder)
	r.DELETE("/api/orders/", handler.DeleteOrder)

	r.POST("/api/orders/:order_id/items", handler.AddOrderItem)
	r.DELETE("/api/orders/:order_id/items/:item_id", handler.RemoveOrderItem)

	r.DELETE("/api/orders/:order_id/pay", handler.PayOrderItem)
}

func (h OrderHandler) CreateOrder(c *gin.Context) {
}

func (h OrderHandler) GetOrder(c *gin.Context) {}

func (h OrderHandler) GetAllOrders(c *gin.Context) {}

func (h OrderHandler) UpdateOrder(c *gin.Context) {}

func (h OrderHandler) DeleteOrder(c *gin.Context) {}

func (h OrderHandler) AddOrderItem(c *gin.Context) {}

func (h OrderHandler) RemoveOrderItem(c *gin.Context) {}

func (h OrderHandler) PayOrderItem(c *gin.Context) {}
