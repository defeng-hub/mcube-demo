package api

import (
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/http/response"
)

func (h *handler) QueryUser(c *gin.Context) {
	response.Success(c.Writer, "666")
}
