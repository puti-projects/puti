package auth

import (
	"github.com/puti-projects/puti/internal/backend/api"
	"github.com/puti-projects/puti/internal/backend/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Login is the Login handler
func Login(c *gin.Context) {
	var u service.LoginRequest
	if err := c.Bind(&u); err != nil {
		api.SendResponse(c, errno.ErrBind, nil)
		return
	}

	token, err := service.LoginAuth(c, u.Username, u.Password)
	if err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	api.SendResponse(c, nil, token)
}
