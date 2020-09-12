package service

import (
	"github.com/puti-projects/puti/internal/backend/dao"
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

// Token represents a JSON web token.
type Token struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

// LoginAuth user login authentication
func LoginAuth(c *gin.Context, username string, password string) (*Token, error) {
	u, err := dao.Engine.GetUser(username)
	if err != nil {
		return nil, errno.New(errno.ErrUserNotFound, err)
	}

	if err := auth.Compare(u.Password, password); err != nil {
		return nil, errno.New(errno.ErrPasswordIncorrect, err)
	}

	t, err := token.Sign(c, token.Context{ID: u.ID, Username: u.Username}, "")
	if err != nil {
		return nil, errno.New(errno.ErrToken, err)
	}

	return &Token{Username: u.Username, Token: t}, nil
}
