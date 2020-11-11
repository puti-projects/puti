package knowledge

import (
	"github.com/puti-projects/puti/internal/admin/api"
	"github.com/puti-projects/puti/internal/admin/service"

	"github.com/gin-gonic/gin"
)

// List knowledge list handler
func List(c *gin.Context) {
	knowledgeList, err := service.GetKnowledgeList()
	if err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	api.SendResponse(c, nil, knowledgeList)
}
