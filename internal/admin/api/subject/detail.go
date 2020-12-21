package subject

import (
	"strconv"

	"github.com/puti-projects/puti/internal/admin/api"
	"github.com/puti-projects/puti/internal/admin/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Detail get subject detail handler
func Detail(c *gin.Context) {
	ID := c.Param("id")
	subjectID, _ := strconv.Atoi(ID)

	svc := service.New(c.Request.Context())
	subject, err := svc.GetSubjectInfo(uint64(subjectID))
	if err != nil {
		api.SendResponse(c, errno.ErrSubjectNotFount, nil)
		return
	}

	api.SendResponse(c, nil, subject)
}
