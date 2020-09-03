package user

import (
	Response "github.com/puti-projects/puti/internal/backend/handler"
	"github.com/puti-projects/puti/internal/model"
	"github.com/puti-projects/puti/internal/pkg/errno"
	"github.com/puti-projects/puti/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/puti-projects/puti/internal/pkg/logger"
	"go.uber.org/zap"
)

// CreateRequest is the create user request params struct
type CreateRequest struct {
	Account       string `form:"account"`
	Nickname      string `form:"nickname"`
	Email         string `form:"email"`
	Role          string `form:"role"`
	Password      string `form:"password"`
	PasswordAgain string `form:"passwordAgain"`
	Website       string `form:"website"`
}

// CreateResponse is the create user request's response struct
type CreateResponse struct {
	Account  string
	Nickname string
}

// Create creates a user
func Create(c *gin.Context) {
	logger.Info("User Create function called.", zap.String("X-request-Id", utils.GetReqID(c)))

	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		Response.SendResponse(c, errno.ErrBind, nil)
		return
	}

	// check params
	if err := r.checkParam(); err != nil {
		Response.SendResponse(c, err, nil)
		return
	}

	if "" == r.Nickname {
		r.Nickname = r.Account
	}

	// TODO
	u := model.UserModel{
		Username: r.Account,
		Password: r.Password,
		Nickname: r.Nickname,
		Email:    r.Email,
		PageURL:  r.Website,
		Status:   1,
		Roles:    r.Role,
	}

	// encrypt password
	if err := u.Encrypt(); err != nil {
		Response.SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	// Insert the user to the database. TODO 字段提示
	if err := u.Create(); err != nil {
		Response.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		Account:  r.Account,
		Nickname: r.Nickname,
	}

	// Show the user information.
	Response.SendResponse(c, nil, rsp)
}

func (r *CreateRequest) checkParam() error {
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
