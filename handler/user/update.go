package user

import (
	"strconv"

	Response "puti/handler"
	"puti/model"
	"puti/pkg/errno"
	"puti/util"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// UpdateRequest is the update user request params struct
type UpdateRequest struct {
	ID            uint64 `form:"id"`
	Nickname      string `form:"nickname"`
	Email         string `form:"email"`
	Role          string `form:"role"`
	Password      string `form:"password"`
	PasswordAgain string `form:"passwordAgain"`
	Website       string `form:"website"`
}

// Update user
func Update(c *gin.Context) {
	log.Info("Update function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	userID, _ := strconv.Atoi(c.Param("id"))

	// Binding the user data.
	var u model.UserModel
	var r UpdateRequest
	if err := c.Bind(&r); err != nil {
		Response.SendResponse(c, errno.ErrBind, nil)
		return
	}

	// We update the record based on the user id.
	r.ID = uint64(userID)

	if r.Nickname != "" {
		u.Nickname = r.Nickname
	}

	if r.Email != "" {
		u.Email = r.Email
	}

	if r.Role != "" {
		u.Roles = r.Role
	}

	if r.Password != "" {
		u.Password = r.Password

		// Encrypt the user password.
		if err := u.Encrypt(); err != nil {
			Response.SendResponse(c, errno.ErrEncrypt, nil)
			return
		}
	}

	if r.Website != "" {
		u.PageURL = r.Website
	}

	// check params
	if err := u.Validate(); err != nil {
		Response.SendResponse(c, err, nil)
		return
	}

	// Save changed fields.
	if err := u.Update(); err != nil {
		Response.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	Response.SendResponse(c, nil, nil)
}
