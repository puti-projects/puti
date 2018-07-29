package user

import (
	Response "gingob/handler"
	"gingob/model"
	"gingob/pkg/errno"
	"gingob/pkg/token"

	"github.com/gin-gonic/gin"
)

// Info gets an user by the user identifier.
func Info(c *gin.Context) {
	t := c.Query("token")

	userContext, err := token.ParseToken(t)

	// Get the user by the `username` from the database.
	user, err := model.GetUser(userContext.Username)
	if err != nil {
		Response.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	Response.SendResponse(c, nil, user)
}
