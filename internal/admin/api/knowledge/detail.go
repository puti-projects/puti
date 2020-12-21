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

	if kID <= 0 {
		api.SendResponse(c, errno.ErrValidation, nil)
	}

	svc := service.New(c.Request.Context())
	knowledge, err := svc.GetKnowledgeInfo(uint64(kID))
	if err != nil {
		api.SendResponse(c, errno.ErrKnowledgeNotFount, nil)
		return
	}

	api.SendResponse(c, nil, knowledge)
}
