package knowledgeitem

import (
	"strconv"

	"github.com/puti-projects/puti/internal/admin/api"
	"github.com/puti-projects/puti/internal/admin/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Detail get knowledge item detail handler
func Detail(c *gin.Context) {
	ID := c.Param("id")
	kID, _ := strconv.Atoi(ID)

	if kID <= 0 {
		api.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	svc := service.New(c.Request.Context())
	knowledgeItem, err := svc.GetKnowledgeItemInfo(kID)
	if err != nil {
		api.SendResponse(c, errno.ErrKnowledgeItemNotFount, nil)
		return
	}

	api.SendResponse(c, nil, knowledgeItem)
}
