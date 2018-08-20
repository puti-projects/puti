package user

import (
	"strconv"

	Response "puti/handler"
	"puti/model"
	"puti/pkg/errno"
	"puti/service"
	"puti/util"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// UpdateRequest is the update user request params struct
type UpdateRequest struct {
	ID       uint64 `form:"id"`
	Nickname string `form:"nickname"`
	Email    string `form:"email"`
	Role     string `form:"role"`
	Website  string `form:"website"`
}

// Update user
func Update(c *gin.Context) {
	log.Info("Update function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	userID, _ := strconv.Atoi(c.Param("id"))

	// Binding the user data.
	var r UpdateRequest
	if err := c.Bind(&r); err != nil {
		Response.SendResponse(c, errno.ErrBind, nil)
		return
	}

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
}
