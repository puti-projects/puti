package user

import (
	Response "gingob/handler"
	"gingob/pkg/errno"
	"gingob/service"

	"github.com/gin-gonic/gin"
)

// List list the users in the database.
func List(c *gin.Context) {
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		Response.SendResponse(c, errno.ErrBind, nil)
		return
	}

	infos, count, err := service.ListUser(r.Username, r.Offset, r.Limit)
	if err != nil {
		Response.SendResponse(c, err, nil)
		return
	}

	Response.SendResponse(c, nil, ListResponse{
		TotalCount: count,
		UserList:   infos,
	})
}
