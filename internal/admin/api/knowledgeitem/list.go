package knowledgeitem

import (
	"strconv"

	"github.com/puti-projects/puti/internal/admin/api"
	"github.com/puti-projects/puti/internal/admin/service"

	"github.com/gin-gonic/gin"
)

// List knowledge item list handler
func List(c *gin.Context) {
	ID := c.Param("id")
	kID, _ := strconv.Atoi(ID)

	svc := service.New(c.Request.Context())
	knowledgeItemList, err := svc.GetKnowledgeItemList(kID)
	if err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	api.SendResponse(c, nil, knowledgeItemList)
}
