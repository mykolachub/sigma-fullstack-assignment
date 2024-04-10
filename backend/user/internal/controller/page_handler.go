package controller

import (
	"net/http"
	"sigma-user/config"
	"sigma-user/internal/util"

	"github.com/adrianbrad/queue"
	"github.com/gin-gonic/gin"
	"github.com/sony/gobreaker"
)

type PageHandler struct {
	pageSvc PageService
	cb      *gobreaker.CircuitBreaker
	q       *queue.Linked[string]
}

func InitPageHandler(r *gin.Engine, pageSvc PageService, cb *gobreaker.CircuitBreaker) {
	handler := PageHandler{pageSvc: pageSvc, cb: cb, q: &queue.Linked[string]{}}

	r.POST("/api/page/track", handler.TrackPage)
	r.GET("/api/page/track", handler.GetPageCount)
}

func (h PageHandler) TrackPage(c *gin.Context) {
	pageName := c.Query(config.Page)
	if pageName == "" {
		svcCode := config.SvcMissingPagePar
		c.JSON(http.StatusBadRequest, util.NewErrResponse(svcCode.Message, svcCode.Code))
		return
	}

	svcCode, err := h.cb.Execute(func() (interface{}, error) {
		return h.pageSvc.TrackPage(h.q, pageName)
	})

	if err != nil {
		message := util.NewErrResponse(svcCode.(config.ServiceCode).Message, svcCode.(config.ServiceCode).Code)
		c.JSON(http.StatusServiceUnavailable, message)
		return
	}
	c.JSON(http.StatusOK, util.NewResponse(svcCode.(config.ServiceCode).Message, svcCode.(config.ServiceCode).Code))
}

func (h PageHandler) GetPageCount(c *gin.Context) {
	pageName := c.Query(config.Page)
	if pageName == "" {
		svcCode := config.SvcMissingPagePar
		c.JSON(http.StatusBadRequest, util.NewErrResponse(svcCode.Message, svcCode.Code))
		return
	}

	page, svcCode, err := h.pageSvc.GetPageCount(pageName)
	if err != nil {
		message := util.NewErrResponse(svcCode.Message, svcCode.Code)
		c.JSON(http.StatusUnprocessableEntity, message)
		return
	}
	c.JSON(http.StatusOK, util.NewResponse(svcCode.Message, svcCode.Code).AddKey(pageName, page))
}
