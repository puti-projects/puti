package subject

import (
	"strconv"

	"github.com/puti-projects/puti/internal/admin/api"
	"github.com/puti-projects/puti/internal/admin/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Delete delete the taxonomy directly without soft delete
func Delete(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	subjectID := uint64(ID)

	if err := service.DeleteSubject(subjectID); err != nil {
		api.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	api.SendResponse(c, nil, nil)
}
