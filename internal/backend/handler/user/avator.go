package user

import (
	Response "github.com/puti-projects/puti/internal/backend/handler"
	"github.com/puti-projects/puti/internal/backend/service"

	"github.com/gin-gonic/gin"
)

// Avatar upload user avatar handler
func Avatar(c *gin.Context) {
	userID := c.PostForm("userId")
	file, _ := c.FormFile("img")

	if err := service.UpdateUserAvatar(c, userID, file); err != nil {
		Response.SendResponse(c, err, nil)
		return
	}

	Response.SendResponse(c, nil, nil)
}
