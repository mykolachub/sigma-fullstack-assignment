package controller

import (
	"net/http"
	"sigma-inventory/proto"

	"github.com/gin-gonic/gin"
)

type InventoryHandler struct {
	invSvc InventoryService
}

func InitInventoryHandler(r *gin.Engine, invSvc InventoryService) {
	handler := InventoryHandler{invSvc: invSvc}

	r.GET("/api/inventory/", handler.GetAllInventory)
	r.GET("/api/inventory/:product_id", handler.GetInventory)

	r.DELETE("/api/inventory/:product_id", handler.DeleteInventory)
	r.PATCH("/api/inventory/:product_id", handler.UpdateInventory)
	r.PATCH("/api/inventory/:product_id/decrement", handler.DecrementInventory)
	r.POST("/api/inventory", handler.CreateInventory)

	r.POST("/api/inventory/reserve", handler.ReserveInventory)
	r.PATCH("/api/inventory/reserve/free", handler.FreeReservedInventory)
}

func (h InventoryHandler) GetInventory(c *gin.Context) {
	id := c.Param("product_id")
	res, err := h.invSvc.GetProduct(c, &proto.GetProductRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h InventoryHandler) GetAllInventory(c *gin.Context) {
	res, err := h.invSvc.GetAllProducts(c, &proto.GetAllProductsRequest{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h InventoryHandler) CreateInventory(c *gin.Context) {
	var req proto.CreateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		return
	}

	res, err := h.invSvc.CreateProduct(c, &req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h InventoryHandler) UpdateInventory(c *gin.Context) {
	id := c.Param("product_id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product id parameter is missing"})
		return
	}

	var req proto.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		return
	}
	req.Id = id

	res, err := h.invSvc.UpdateProduct(c, &req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h InventoryHandler) DecrementInventory(c *gin.Context) {
	id := c.Param("product_id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product id parameter is missing"})
		return
	}

	var req proto.DecrementInventoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		return
	}

	res, err := h.invSvc.DecrementInventory(c, &req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h InventoryHandler) DeleteInventory(c *gin.Context) {
	id := c.Param("product_id")
	res, err := h.invSvc.DeleteProduct(c, &proto.DeleteProductRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h InventoryHandler) ReserveInventory(c *gin.Context) {
	var req proto.ReserveInventoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		return
	}

	res, err := h.invSvc.ReserveInventory(c, &req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h InventoryHandler) FreeReservedInventory(c *gin.Context) {
	var req proto.FreeReservedInventoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		return
	}

	res, err := h.invSvc.FreeReservedInventory(c, &req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
