package auth

import (
	Response "github.com/puti-projects/puti/internal/backend/handler"
	"github.com/puti-projects/puti/internal/backend/service"
	"github.com/puti-projects/puti/internal/common/model"
	"github.com/puti-projects/puti/internal/pkg/auth"
	"github.com/puti-projects/puti/internal/pkg/errno"
	"github.com/puti-projects/puti/internal/pkg/token"

	"github.com/gin-gonic/gin"
)

// LoginRequest is the login request params struct
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

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

	Response.SendResponse(c, nil, service.Token{Username: u.Username, Token: t})
}
