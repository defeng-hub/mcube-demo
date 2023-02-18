package api

import (
	"context"
	"github.com/defeng-hub/mcube-demo/apps/rbac"
	"github.com/defeng-hub/mcube-demo/util/exception"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/http/response"
)

func (h *handler) CreateRole(c *gin.Context) {
	req := rbac.CreateRoleRequest{}
	err := c.Bind(&req)
	if err != nil {
		newErr := exception.DefaultException(-1, "请求解析失败", nil)
		response.Failed(c.Writer, newErr)
		return
	}

	role, err := h.roleService.CreateRole(context.Background(), &req)
	if err != nil {
		newErr := exception.DefaultException(-1, err.Error(), nil)
		response.Failed(c.Writer, newErr)
		return
	}

	response.Success(c.Writer, 0, "角色创建成功", role)
}

func (h *handler) DeleteRole(c *gin.Context) {
	req := rbac.DeleteRoleRequest{}
	err := c.Bind(&req)
	if err != nil {
		newErr := exception.DefaultException(-1, "请求解析失败", nil)
		response.Failed(c.Writer, newErr)
		return
	}
	if _, err = h.roleService.DeleteRole(context.Background(), &req); err != nil {
		response.Failed(c.Writer, err)
	} else {
		response.Success(c.Writer, 0, "删除成功", nil)
		return
	}
}

func (h *handler) QueryRole(c *gin.Context) {
	req := rbac.QueryRoleRequest{}
	err := c.Bind(&req)
	if err != nil {
		newErr := exception.DefaultException(-1, "请求解析失败", nil)
		response.Failed(c.Writer, newErr)
		return
	}

	if res, err := h.roleService.QueryRole(context.Background(), &req); err != nil {
		response.Failed(c.Writer, err)
	} else {
		response.Success(c.Writer, 0, "获取成功", res)
		return
	}
}
