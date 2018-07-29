package user

import (
	"fmt"
	Response "gingob/handler"
	"gingob/model"
	"gingob/pkg/errno"
	"gingob/util"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// Create creates a user
func Create(c *gin.Context) {
	log.Info("User Create function called.", lager.Data{"X-request-Id": util.GetReqID(c)})

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

	// TODO
	u := model.UserModel{
		Username:       r.Username,
		Password:       r.Password,
		Nickname:       r.Username,
		Email:          "example@example.com",
		Status:         0,
		Roles:          "administrator",
		RegisteredTime: time.Now(),
	}

	// encrypt password
	if err := u.Encrypt(); err != nil {
		Response.SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	// Insert the user to the database. TODO 字段提示
	if err := u.Create(); err != nil {
		fmt.Print(err)
		Response.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	// Show the user information.
	Response.SendResponse(c, nil, rsp)
}

func (r *CreateRequest) checkParam() error {
	if r.Username == "" {
		return errno.New(errno.ErrValidation, nil).Add("username is empty.")
	}

	if r.Password == "" {
		return errno.New(errno.ErrValidation, nil).Add("password is empty.")
	}

	return nil
}
