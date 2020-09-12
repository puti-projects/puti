package auth

import (
	"github.com/puti-projects/puti/internal/backend/api"
	"github.com/puti-projects/puti/internal/backend/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Info get an user by the user identifier.
func Info(c *gin.Context) {
	t := c.Query("token")

	user, err := service.GetUserByToken(t)
	if err != nil {
		api.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	api.SendResponse(c, nil, user)
}
