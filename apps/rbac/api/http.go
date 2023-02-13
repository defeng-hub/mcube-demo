package api

import (
	"github.com/defeng-hub/mcube-demo/apps/rbac"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	h = &handler{}
)

type handler struct {
	service rbac.UserServiceServer
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(rbac.AppName)
	h.service = app.GetGrpcApp(rbac.AppName).(rbac.UserServiceServer)
	return nil
}

func (h *handler) Name() string {
	return rbac.AppName
}

func (h *handler) Registry(r gin.IRouter) {
	r.POST("/QueryUser", h.QueryUser)
	r.POST("/CreateUser", h.CreateUser)
	r.POST("/DeleteUser", h.DeleteUser)
}

func init() {
	app.RegistryGinApp(h)
}
