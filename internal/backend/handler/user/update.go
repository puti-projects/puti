package user

import (
	"strconv"

	Response "github.com/puti-projects/puti/internal/backend/handler"
	"github.com/puti-projects/puti/internal/backend/service"
	"github.com/puti-projects/puti/internal/common/model"
	"github.com/puti-projects/puti/internal/common/utils"
	"github.com/puti-projects/puti/internal/pkg/errno"
	"github.com/puti-projects/puti/internal/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
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

// UpdatePasswordRequest user for reset user's password during the profile page
type UpdatePasswordRequest struct {
	ID            uint64 `json:"id"`
	Password      string `json:"password" binding:"required"`
	PasswordAgain string `json:"passwordAgain" binding:"required"`
}

// Update user
func Update(c *gin.Context) {
	logger.Info("Update function called.", zap.String("X-request-Id", utils.GetReqID(c)))

	// Get user id
	userID, _ := strconv.Atoi(c.Param("id"))

	var r UpdateRequest
	var s UpdateStatusRequest
	var p UpdatePasswordRequest

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
	} else if errPassword := c.ShouldBindBodyWith(&p, binding.JSON); errPassword == nil {
		p.ID = uint64(userID)

		// check two input password
		if p.Password != p.PasswordAgain {
			Response.SendResponse(c, errno.ErrValidation, nil)
			return
		}

		user := &model.UserModel{
			Model: model.Model{ID: p.ID},

			Password: p.Password,
		}

		// encrypt password
		if err := user.Encrypt(); err != nil {
			Response.SendResponse(c, errno.ErrEncrypt, nil)
			return
		}

		// Update changed fields.
		if errPassword := service.UpdateUserPassword(user); errPassword != nil {
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
