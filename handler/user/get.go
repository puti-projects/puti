package user

import (
	Response "gingob/handler"
	"gingob/model"
	"gingob/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Get gets an user by the user identifier.
func Get(c *gin.Context) {
	username := c.Param("username")
	// Get the user by the `username` from the database.
	user, err := model.GetUser(username)
	if err != nil {
		Response.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	Response.SendResponse(c, nil, user)
}
