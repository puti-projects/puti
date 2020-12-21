package knowledgeitem

import (
	"strconv"

	"github.com/puti-projects/puti/internal/admin/api"
	"github.com/puti-projects/puti/internal/admin/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Update update the knowledge item handler
func Update(c *gin.Context) {
	ID := c.Param("id")
	kItemID, _ := strconv.Atoi(ID)

	ir := service.KnowledgeItemUpdateInfoRequest{}
	cr := service.KnowledgeItemUpdateContentRequest{}
	svc := service.New(c.Request.Context())
	if errInfo := c.ShouldBindBodyWith(&ir, binding.JSON); errInfo == nil {
		err := svc.UpdateKnowledgeItemInfo(&ir, uint64(kItemID))
		if err != nil {
			api.SendResponse(c, err, nil)
			return
		}

		api.SendResponse(c, nil, nil)
		return
	} else if errContent := c.ShouldBindBodyWith(&cr, binding.JSON); errContent == nil {
		rsp, err := svc.UpdateKnowledgeItemContent(&cr, uint64(kItemID))
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
