package media

import (
	"github.com/puti-projects/puti/internal/admin/api"
	"github.com/puti-projects/puti/internal/admin/service"

	"github.com/gin-gonic/gin"
)

// Upload file upload handler
func Upload(c *gin.Context) {
	// get userId and file
	userID := c.PostForm("userId")
	usage := c.DefaultPostForm("usage", "common")
	file, _ := c.FormFile("file")

	svc := service.New(c.Request.Context())
	ID, GUID, err := svc.UploadMedia(c, userID, usage, file)
	if err != nil {
		api.SendResponse(c, err, nil)
	}

	rsp := service.MediaUploadResponse{
		ID:  ID,
		URL: GUID,
	}
	api.SendResponse(c, nil, rsp)
}
