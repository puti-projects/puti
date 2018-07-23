package user

import (
	Response "gingob/handler"
	"gingob/model"
	"gingob/pkg/auth"
	"gingob/pkg/errno"
	"gingob/pkg/token"

	"github.com/gin-gonic/gin"
)

// Login is the Login function
func Login(c *gin.Context) {
	var u LoginRequest
	if err := c.Bind(&u); err != nil {
		Response.SendResponse(c, errno.ErrBind, nil)
		return
	}

	d, err := model.GetUser(u.Username)
	if err != nil {
		Response.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	if err := auth.Compare(d.Password, u.Password); err != nil {
		Response.SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	t, err := token.Sign(c, token.Context{ID: d.ID, Username: d.Username}, "")
	if err != nil {
		Response.SendResponse(c, errno.ErrToken, nil)
		return
	}

	Response.SendResponse(c, nil, model.Token{Username: u.Username, Token: t})
}
