package knowledge

import (
	"strconv"

	"github.com/puti-projects/puti/internal/admin/api"
	"github.com/puti-projects/puti/internal/admin/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Detail get knowledge info detail handler
func Detail(c *gin.Context) {
	ID := c.Param("id")
	kID, _ := strconv.Atoi(ID)

	knowledge, err := service.GetKnowledgeInfo(uint64(kID))
	if err != nil {
		api.SendResponse(c, errno.ErrSubjectNotFount, nil)
		return
	}

	api.SendResponse(c, nil, knowledge)
}
