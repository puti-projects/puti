package media

import (
	Response "github.com/puti-projects/puti/internal/backend/handler"
	"github.com/puti-projects/puti/internal/backend/service"

	"github.com/gin-gonic/gin"
)

// Upload file upload handler
func Upload(c *gin.Context) {
	// get userId and file
	userID := c.PostForm("userId")
	usage := c.DefaultPostForm("usage", "common")
	file, _ := c.FormFile("file")

	ID, GUID, err := service.UploadMedia(c, userID, usage, file)
	if err != nil {
		Response.SendResponse(c, err, nil)
	}

	rsp := service.MediaUploadResponse{
		ID:  ID,
		URL: GUID,
	}
	Response.SendResponse(c, nil, rsp)
}
