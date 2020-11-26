package knowledgeItem

import (
	"github.com/gin-gonic/gin"
	"github.com/puti-projects/puti/internal/admin/api"
	"github.com/puti-projects/puti/internal/admin/service"
	"github.com/puti-projects/puti/internal/pkg/errno"
	"github.com/puti-projects/puti/internal/pkg/token"
)

// Create create knowledge item handler
func Create(c *gin.Context) {
	// get token and parse
	t := c.Query("token")
	userContext, err := token.ParseToken(t)
	if err != nil {
		api.SendResponse(c, errno.ErrTokenInvalid, nil)
		return
	}

	var r service.KnowledgeItemCreateRequest
	if err := c.ShouldBind(&r); err != nil {
		api.SendResponse(c, errno.ErrBind, nil)
		return
	}

	// check params
	if err := checkCreateParam(&r); err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	rsp, err := service.CreateKnowledgeItem(&r, userContext.ID)
	if err != nil {
		api.SendResponse(c, errno.ErrKnowledgeItemCreateFailed, nil)
		return
	}

	api.SendResponse(c, nil, rsp)
}

func checkCreateParam(r *service.KnowledgeItemCreateRequest) error {
	if r.Title == "" {
		return errno.New(errno.ErrValidation, nil).Add("Title can not be empty.")
	}

	return nil
}
