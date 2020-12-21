package knowledgeitem

import (
	"strconv"

	"github.com/puti-projects/puti/internal/admin/api"
	"github.com/puti-projects/puti/internal/admin/service"

	"github.com/gin-gonic/gin"
)

// Delete delete article handler
func Delete(c *gin.Context) {
	kItemID, _ := strconv.Atoi(c.Param("id"))

	svc := service.New(c.Request.Context())
	if err := svc.DeleteKnowledgeItem(uint64(kItemID)); err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	api.SendResponse(c, nil, nil)
}
