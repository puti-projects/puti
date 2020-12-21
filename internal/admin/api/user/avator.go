package user

import (
	"github.com/puti-projects/puti/internal/admin/api"
	"github.com/puti-projects/puti/internal/admin/service"

	"github.com/gin-gonic/gin"
)

// Avatar upload user avatar handler
func Avatar(c *gin.Context) {
	userID := c.PostForm("userId")
	file, _ := c.FormFile("img")

	svc := service.New(c.Request.Context())
	if err := svc.UpdateUserAvatar(c, userID, file); err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	api.SendResponse(c, nil, nil)
}
