package auth

import (
	"github.com/puti-projects/puti/internal/admin/api"
	"github.com/puti-projects/puti/internal/admin/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Info get an user by the user identifier.
func Info(c *gin.Context) {
	t := c.Query("token")

	svc := service.New(c.Request.Context())
	user, err := svc.GetUserByToken(t)
	if err != nil {
		api.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	api.SendResponse(c, nil, user)
}
