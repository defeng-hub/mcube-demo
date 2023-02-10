package api

import (
	"context"
	rbac "github.com/defeng-hub/mcube-demo/apps/rbac"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/http/response"
)

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) QueryUser(c *gin.Context) {
	response.Success(c.Writer, "666")
}

func (h *handler) CreateUser(c *gin.Context) {
	user, err := h.service.CreateUser(context.Background(), &rbac.CreateUserRequest{
		UserName: "www",
		Pwd:      "www",
		Email:    "www",
		Phone:    "www",
		Address:  "wwwwwwww",
	})
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}
	h.log.Infof("创建用户:%s", user)
	response.Success(c.Writer, user)
}
