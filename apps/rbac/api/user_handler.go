package api

import (
	"context"
	rbac "github.com/defeng-hub/mcube-demo/apps/rbac"
	"github.com/defeng-hub/mcube-demo/util/exception"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/http/response"
)

func (h *handler) QueryUser(c *gin.Context) {
	req := rbac.QueryUserRequest{Page: &rbac.PageRequest{}}
	c.Bind(&req)

	userSet, err := h.userService.QueryUser(context.Background(), &req)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}

	response.Success(c.Writer, 0, "用户列表查询成功", userSet)
}

func (h *handler) CreateUser(c *gin.Context) {
	userReq := rbac.CreateUserRequest{}
	err := c.Bind(&userReq)
	if err != nil {
		return
	}
	h.log.Infof("创建用户:%s", &userReq)

	user, err := h.userService.CreateUser(context.Background(), &userReq)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}
	response.Success(c.Writer, 0, "创建用户成功", user)
}

func (h *handler) DeleteUser(c *gin.Context) {
	req := rbac.DeleteUserRequest{}

	err := c.Bind(&req)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}

	user, err := h.userService.DeleteUser(context.Background(), &req)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			newErr := exception.DefaultException(-1, "用户不存在", nil)
			response.Failed(c.Writer, newErr)
			return
		}

		response.Failed(c.Writer, err)
		return
	}

	response.Success(c.Writer, 0, "删除用户成功", user)
	return
}

func (h *handler) CreateUserRole(c *gin.Context) {
	req := rbac.CreateUserRoleRequest{}
	err := c.Bind(&req)
	if err != nil {
		newErr := exception.DefaultException(-1, "请求解析失败", nil)
		response.Failed(c.Writer, newErr)
		return
	}

	role, err := h.userRoleService.CreateUserRole(context.Background(), &req)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	} else {
		response.Success(c.Writer, 0, "添加权限成功", role)
	}
}

func (h *handler) DeleteUserRole(c *gin.Context) {
	req := rbac.DeleteUserRoleRequest{}
	err := c.Bind(&req)
	if err != nil {
		newErr := exception.DefaultException(-1, "请求解析失败", nil)
		response.Failed(c.Writer, newErr)
		return
	}
	role, err := h.userRoleService.DeleteUserRole(context.Background(), &req)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	} else {
		response.Success(c.Writer, 0, "删除权限成功", role)
	}

}
