package controller

import (
	"net/http"
	"sigma-order/proto"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	ordSvc OrderService
}

func InitOrderHandler(r *gin.Engine, ordSvc OrderService) {
	handler := OrderHandler{ordSvc: ordSvc}

	r.POST("/api/orders/", handler.CreateOrder)
	r.GET("/api/orders/", handler.GetAllOrders)
	r.GET("/api/orders/:order_id", handler.GetOrder)

	r.PATCH("/api/orders/", handler.UpdateOrder)
	r.DELETE("/api/orders/:order_id", handler.DeleteOrder)

	r.POST("/api/orders/:order_id/items", handler.AddOrderItem)
	r.DELETE("/api/orders/:order_id/items/:item_id", handler.RemoveOrderItem)

	r.PATCH("/api/orders/:order_id/pay", handler.PayOrderItem)
}

func (h OrderHandler) CreateOrder(c *gin.Context) {
	var req proto.CreateOrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		return
	}

	res, err := h.ordSvc.CreateOrder(c, &req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": res})
}

func (h OrderHandler) GetOrder(c *gin.Context) {
	order_id := c.Param("order_id")

	order, err := h.ordSvc.GetOrder(c, &proto.GetOrderRequest{Id: order_id})
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": order})
}

func (h OrderHandler) GetAllOrders(c *gin.Context) {
	res, err := h.ordSvc.GetAllOrders(c, &proto.GetAllOrdersRequest{})
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": res})
}

func (h OrderHandler) UpdateOrder(c *gin.Context) {
	var req proto.UpdateOrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		return
	}

	res, err := h.ordSvc.UpdateOrder(c, &req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": res})
}

func (h OrderHandler) DeleteOrder(c *gin.Context) {
	order_id := c.Param("order_id")

	res, err := h.ordSvc.DeleteOrder(c, &proto.DeleteOrderRequest{OrderId: order_id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": res})
}

func (h OrderHandler) AddOrderItem(c *gin.Context) {
	order_id := c.Param("order_id")
	var req proto.AddOrderItemRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		return
	}
	req.OrderId = order_id

	res, err := h.ordSvc.AddOrderItem(c, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": res})
}

func (h OrderHandler) RemoveOrderItem(c *gin.Context) {
	order_id := c.Param("order_id")
	item_id := c.Param("item_id")

	res, err := h.ordSvc.RemoveOrderItem(c, &proto.RemoveOrderItemRequest{OrderId: order_id, OrderItemId: item_id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": res})
}

func (h OrderHandler) PayOrderItem(c *gin.Context) {
	order_id := c.Param("order_id")

	res, err := h.ordSvc.PayOrder(c, &proto.PayOrderRequest{OrderId: order_id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": res})
}
