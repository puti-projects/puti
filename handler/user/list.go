package user

import (
	Response "puti/handler"
	"puti/pkg/errno"
	"puti/service"

	"github.com/gin-gonic/gin"
)

// ListRequest is the user list request struct
type ListRequest struct {
	Username string `form:"username"`
	Number   int    `form:"number"`
	Page     int    `form:"page"`
	Status   int    `form:"status"`
	Role     string `form:"role"`
}

// ListResponse is the use list response struct
type ListResponse struct {
	TotalCount uint64              `json:"totalCount"`
	UserList   []*service.UserInfo `json:"userList"`
}

// List list the users in the database.
func List(c *gin.Context) {
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		Response.SendResponse(c, errno.ErrBind, nil)
		return
	}

	infos, count, err := service.ListUser(r.Username, r.Role, r.Number, r.Page, r.Status)
	if err != nil {
		Response.SendResponse(c, err, nil)
		return
	}

	Response.SendResponse(c, nil, ListResponse{
		TotalCount: count,
		UserList:   infos,
	})
}
