package user

import (
	"strconv"

	Response "puti/handler"
	"puti/model"
	"puti/pkg/errno"
	"puti/service"
	"puti/util"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// UpdateRequest is the update user request params struct
type UpdateRequest struct {
	ID       uint64 `json:"id"`
	Nickname string `json:"nickname"`
	Email    string `json:"email" binding:"required"`
	Role     string `json:"role" binding:"required"`
	Website  string `json:"website"`
}

// UpdateStatusRequest only use for update user status
type UpdateStatusRequest struct {
	ID     uint64 `json:"id"`
	Status int    `json:"status" binding:"required"`
}

// Update user
func Update(c *gin.Context) {
	log.Info("Update function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	// Get user id
	userID, _ := strconv.Atoi(c.Param("id"))

	var r UpdateRequest
	var s UpdateStatusRequest

	if err := c.ShouldBindBodyWith(&r, binding.JSON); err == nil {
		// We update the record based on the user id.
		r.ID = uint64(userID)

		user := &model.UserModel{
			Model: model.Model{ID: r.ID},

			Nickname: r.Nickname,
			Email:    r.Email,
			PageURL:  r.Website,
			Roles:    r.Role,
		}

		// Update changed fields.
		if err := service.UpdateUser(user); err != nil {
			Response.SendResponse(c, errno.ErrDatabase, nil)
			return
		}

		Response.SendResponse(c, nil, nil)
		return
	} else if errStatus := c.ShouldBindBodyWith(&s, binding.JSON); errStatus == nil {
		// We update the record based on the user id.
		s.ID = uint64(userID)

		user := &model.UserModel{
			Model: model.Model{ID: s.ID},

			Status: s.Status,
		}

		// Update changed fields.
		if errStatus := service.UpdateUserStatus(user); errStatus != nil {
			Response.SendResponse(c, errno.ErrDatabase, nil)
			return
		}

		Response.SendResponse(c, nil, nil)
		return
	} else {
		Response.SendResponse(c, errno.ErrBind, nil)
		return
	}
}
