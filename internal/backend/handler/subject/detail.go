package subject

import (
	"strconv"

	Response "github.com/puti-projects/puti/internal/backend/handler"
	"github.com/puti-projects/puti/internal/backend/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Detail get subject detail
func Detail(c *gin.Context) {
	ID := c.Param("id")
	subjectID, err := strconv.Atoi(ID)

	subject, err := service.GetSubjectInfo(uint64(subjectID))
	if err != nil {
		Response.SendResponse(c, errno.ErrSubjectNotFount, nil)
		return
	}

	Response.SendResponse(c, nil, subject)
}
