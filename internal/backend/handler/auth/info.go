package auth

import (
	Response "github.com/puti-projects/puti/internal/backend/handler"
	"github.com/puti-projects/puti/internal/backend/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Info get an user by the user identifier.
func Info(c *gin.Context) {
	t := c.Query("token")

	user, err := service.GetUserByToken(t)
	if err != nil {
		Response.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	Response.SendResponse(c, nil, user)
}
