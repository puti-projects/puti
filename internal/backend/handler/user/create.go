package user

import (
	Response "github.com/puti-projects/puti/internal/backend/handler"
	"github.com/puti-projects/puti/internal/backend/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Create user create handler
func Create(c *gin.Context) {
	var r service.UserCreateRequest
	if err := c.Bind(&r); err != nil {
		Response.SendResponse(c, errno.ErrBind, nil)
		return
	}

	// check params
	if err := checkParam(&r); err != nil {
		Response.SendResponse(c, err, nil)
		return
	}

	username, nickname, err := service.CreateUser(&r)
	if err != nil {
		Response.SendResponse(c, err, nil)
	}

	rsp := &service.UserCreateResponse{
		Account:  username,
		Nickname: nickname,
	}

	// Show the user information.
	Response.SendResponse(c, nil, rsp)
}

func checkParam(r *service.UserCreateRequest) error {
	if r.Account == "" {
		return errno.New(errno.ErrValidation, nil).Add("account is empty.")
	}

	if r.Password == "" {
		return errno.New(errno.ErrValidation, nil).Add("password is empty.")
	}

	if r.PasswordAgain == "" {
		return errno.New(errno.ErrValidation, nil).Add("check password is empty.")
	}

	if r.Password != r.PasswordAgain {
		return errno.New(errno.ErrValidation, nil).Add("check password is incorrect.")
	}

	if r.Email == "" {
		return errno.New(errno.ErrValidation, nil).Add("Email is empty.")
	}

	if r.Role == "" {
		return errno.New(errno.ErrValidation, nil).Add("role is empty.")
	}

	if r.Role != "administrator" && r.Role != "writer" && r.Role != "subscriber" {
		return errno.New(errno.ErrValidation, nil).Add("role is incorrect.")
	}

	return nil
}
