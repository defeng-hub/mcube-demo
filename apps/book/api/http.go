package api

import (
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/defeng-hub/mcube-demo/apps/book"
)

var (
	h = &handler{}
)

type handler struct {
	service book.ServiceServer
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(book.AppName)
	h.service = app.GetGrpcApp(book.AppName).(book.ServiceServer)
	return nil
}

func (h *handler) Name() string {
	return book.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(r gin.IRouter) {
	r.POST("/", h.CreateBook)
	r.GET("/", h.QueryBook)
}

func init() {
	app.RegistryGinApp(h)
}
