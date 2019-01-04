package auth

import (
	Response "github.com/puti-projects/puti/internal/backend/handler"
	"github.com/puti-projects/puti/internal/pkg/errno"
	"github.com/puti-projects/puti/internal/pkg/token"
	"github.com/puti-projects/puti/internal/backend/service"

	"github.com/gin-gonic/gin"
)

// Info gets an user by the user identifier.
func Info(c *gin.Context) {
	t := c.Query("token")

	userContext, err := token.ParseToken(t)

	// Get the user by the `username` from the database.
	user, err := service.GetUser(userContext.Username)
	if err != nil {
		Response.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	Response.SendResponse(c, nil, user)
}
