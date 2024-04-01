package controller

import (
	"net/http"
	"sigma-test/config"
	"sigma-test/internal/util"

	"github.com/gin-gonic/gin"
)

type PageHandler struct {
	pageSvc PageService
}

func InitPageHandler(r *gin.Engine, pageSvc PageService) {
	handler := PageHandler{pageSvc: pageSvc}

	r.POST("/api/page/track", handler.TrackPage)
	r.GET("/api/page/track", handler.GetPageCount)
}

func (h PageHandler) TrackPage(c *gin.Context) {
	pageName := c.Query(config.Page)
	if pageName == "" {
		message := util.MakeMessage(util.MessageError, config.ErrMissingPagePar.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, message)
		return
	}

	err := h.pageSvc.TrackPage(pageName)
	if err != nil {
		message := util.MakeMessage(util.MessageError, config.ErrFailedTrackPage.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, message)
		return
	}
	c.JSON(http.StatusOK, util.MakeMessage(util.MessageSuccess, config.MsgPageCountet, nil))
}

func (h PageHandler) GetPageCount(c *gin.Context) {
	pageName := c.Query(config.Page)
	if pageName == "" {
		message := util.MakeMessage(util.MessageError, config.ErrMissingPagePar.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, message)
		return
	}

	page, err := h.pageSvc.GetPageCount(pageName)
	if err != nil {
		message := util.MakeMessage(util.MessageError, config.ErrFailedGetPage.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, message)
		return
	}
	c.JSON(http.StatusOK, util.MakeMessage(util.MessageSuccess, config.MsgEmpty, page))
}
