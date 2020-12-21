package media

import (
	"strconv"

	"github.com/puti-projects/puti/internal/admin/api"
	"github.com/puti-projects/puti/internal/admin/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Update update media info handler
func Update(c *gin.Context) {
	// Get user id
	userID, _ := strconv.Atoi(c.Param("id"))

	var r service.MediaUpdateRequest
	if err := c.ShouldBind(&r); err != nil {
		api.SendResponse(c, errno.ErrBind, nil)
		return
	}

	if r.Title == "" {
		api.SendResponse(c, errno.ErrTitleEmpty, nil)
		return
	}

	svc := service.New(c.Request.Context())
	if err := svc.UpdateMedia(&r, userID); err != nil {
		api.SendResponse(c, err, nil)
	}

	api.SendResponse(c, nil, nil)
	return
}
