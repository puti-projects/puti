package user

import (
	"strconv"

	"github.com/puti-projects/puti/internal/admin/api"
	"github.com/puti-projects/puti/internal/admin/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Update user
func Update(c *gin.Context) {
	// Get user id
	userID, _ := strconv.Atoi(c.Param("id"))

	var r service.UserUpdateRequest
	var s service.UserUpdateStatusRequest
	var p service.UserUpdatePasswordRequest

	if err := c.ShouldBindBodyWith(&r, binding.JSON); err == nil {
		if err := service.UpdateUser(&r, userID); err != nil {
			api.SendResponse(c, err, nil)
			return
		}

		api.SendResponse(c, nil, nil)
		return
	} else if errStatus := c.ShouldBindBodyWith(&s, binding.JSON); errStatus == nil {
		if err := service.UpdateUserStatus(&s, userID); err != nil {
			api.SendResponse(c, err, nil)
			return
		}

		api.SendResponse(c, nil, nil)
		return
	} else if errPassword := c.ShouldBindBodyWith(&p, binding.JSON); errPassword == nil {
		// check two input password
		if p.Password != p.PasswordAgain {
			api.SendResponse(c, errno.ErrValidation, nil)
			return
		}

		if err := service.UpdateUserPassword(&p, userID); err != nil {
			api.SendResponse(c, err, nil)
			return
		}

		api.SendResponse(c, nil, nil)
		return
	}

	api.SendResponse(c, errno.ErrBind, nil)
	return
}
