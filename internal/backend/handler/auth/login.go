package auth

import (
	Response "github.com/puti-projects/puti/internal/backend/handler"
	"github.com/puti-projects/puti/internal/backend/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Login is the Login handler
func Login(c *gin.Context) {
	var u service.LoginRequest
	if err := c.Bind(&u); err != nil {
		Response.SendResponse(c, errno.ErrBind, nil)
		return
	}

	token, err := service.LoginAuth(c, u.Username, u.Password)
	if err != nil {
		Response.SendResponse(c, err, nil)
		return
	}

	Response.SendResponse(c, nil, token)
}
