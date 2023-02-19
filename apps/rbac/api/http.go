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
	userService     rbac.UserServiceServer
	roleService     rbac.RoleServiceServer
	userRoleService rbac.UserRoleServiceServer
	log             logger.Logger
}

func init() {
	app.RegistryGinApp(h)
}

func (h *handler) Config() error {
	h.log = zap.L().Named(rbac.AppName)
	h.userService = app.GetGrpcApp(rbac.AppName).(rbac.UserServiceServer)
	h.roleService = app.GetGrpcApp(rbac.AppName).(rbac.RoleServiceServer)
	h.userRoleService = app.GetGrpcApp(rbac.AppName).(rbac.UserRoleServiceServer)
	return nil
}

func (h *handler) Name() string {
	return rbac.AppName
}

func (h *handler) Version() string {
	return ""
}

func (h *handler) Registry(r gin.IRouter) {
	userRouter := r.Group("/user")
	{
		userRouter.POST("/QueryUser", h.QueryUser)
		userRouter.POST("/CreateUser", h.CreateUser)
		userRouter.POST("/DeleteUser", h.DeleteUser)
		userRouter.POST("/CreateUserRole", h.CreateUserRole)
		userRouter.POST("/DeleteUserRole", h.DeleteUserRole)
	}

	roleRouter := r.Group("/role")
	{
		roleRouter.POST("/CreateRole", h.CreateRole)
		roleRouter.POST("/DeleteRole", h.DeleteRole)
		roleRouter.POST("/QueryRole", h.QueryRole)
	}
}
