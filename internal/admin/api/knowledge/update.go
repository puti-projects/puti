package knowledge

import (
	"github.com/gin-gonic/gin"
	"github.com/puti-projects/puti/internal/admin/api"
	"github.com/puti-projects/puti/internal/admin/service"
	"github.com/puti-projects/puti/internal/pkg/errno"
)

// Update update the knowledge info handler
func Update(c *gin.Context) {
	var r service.KnowledgeUpdateRequest
	if err := c.Bind(&r); err != nil {
		api.SendResponse(c, errno.ErrBind, nil)
		return
	}

	// check params
	if err := checkUpdateParam(&r); err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	// Update changed fields.
	if err := service.UpdateKnowledge(&r); err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	api.SendResponse(c, nil, nil)
}

func checkUpdateParam(r *service.KnowledgeUpdateRequest) error {
	if !service.CheckKnowledgeType(r.Type) {
		return errno.New(errno.ErrKnowledgeType, nil)
	}

	if r.Name == "" {
		return errno.New(errno.ErrValidation, nil).Add("name is empty.")
	}

	if r.Slug == "" {
		r.Slug = r.Name
	}

	return nil
}
