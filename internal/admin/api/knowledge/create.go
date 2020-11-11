package knowledge

import (
	"github.com/puti-projects/puti/internal/admin/api"
	"github.com/puti-projects/puti/internal/admin/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Create create knowledge handler
func Create(c *gin.Context) {
	var r service.KnowledgeCreateRequest
	if err := c.Bind(&r); err != nil {
		api.SendResponse(c, errno.ErrBind, nil)
		return
	}

	// check params
	if err := checkCreateParam(&r); err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	if err := service.CreateKnowledge(&r); err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	api.SendResponse(c, nil, nil)
}

func checkCreateParam(r *service.KnowledgeCreateRequest) error {
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
