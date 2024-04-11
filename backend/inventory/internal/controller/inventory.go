package controller

import (
	"log"
	"net/http"
	"sigma-inventory/proto"

	"github.com/gin-gonic/gin"
)

type InventoryService interface {
	proto.InventoryServiceClient
}

type InventoryHandler struct {
	invSvc InventoryService
}

func InitInventoryHandler(r *gin.Engine, invSvc InventoryService) {
	handler := InventoryHandler{invSvc: invSvc}

	r.GET("/api/inventory", handler.test)
}

func (h InventoryHandler) test(c *gin.Context) {
	res, err := h.invSvc.CreateProduct(c, &proto.CreateProductRequest{})
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, res)
}
