package api

import (
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/http/response"

	"github.com/defeng-hub/mcube-demo/apps/book"
)

func (h *handler) CreateBook(c *gin.Context) {
	req := book.NewCreateBookRequest()

	if err := c.BindJSON(req); err != nil {
		response.Failed(c.Writer, err)
		return
	}

	set, err := h.service.CreateBook(c.Request.Context(), req)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}

	response.Success(c.Writer, 0, "", set)
}

func (h *handler) QueryBook(c *gin.Context) {
	response.Success(c.Writer, 0, "666", nil)
}
