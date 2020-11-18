package knowledgeItem

import (
	"github.com/puti-projects/puti/internal/admin/api"
	"github.com/puti-projects/puti/internal/admin/service"
	"github.com/puti-projects/puti/internal/pkg/errno"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Update update the knowledge item handler
func Update(c *gin.Context) {
	ID := c.Param("id")
	kItemID, _ := strconv.Atoi(ID)

	ir := service.KnowledgeItemUpdateInfoRequest{}
	cr := service.KnowledgeItemUpdateContentRequest{}
	if errInfo := c.ShouldBindBodyWith(&ir, binding.JSON); errInfo == nil {
		err := service.UpdateKnowledgeItemInfo(&ir, uint64(kItemID))
		if err != nil {
			api.SendResponse(c, err, nil)
			return
		}

		api.SendResponse(c, nil, nil)
		return
	} else if errContent := c.ShouldBindBodyWith(&cr, binding.JSON); errContent == nil {
		rsp, err := service.UpdateKnowledgeItemContent(&cr, uint64(kItemID))
		if err != nil {
			api.SendResponse(c, err, nil)
			return
		}

		api.SendResponse(c, nil, rsp)
		return
	}

	api.SendResponse(c, errno.ErrBind, nil)
	return
}
