package user

import (
	"github.com/puti-projects/puti/internal/backend/api"
	"github.com/puti-projects/puti/internal/backend/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Get gets an user by the user identifier.
func Get(c *gin.Context) {
	username := c.Param("username")

	user, err := service.GetUser(username)
	if err != nil {
		api.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	api.SendResponse(c, nil, user)
}
