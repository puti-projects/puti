package taxonomy

import (
	"github.com/puti-projects/puti/internal/admin/api"
	"github.com/puti-projects/puti/internal/admin/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// List taxonomy list handler
func List(c *gin.Context) {
	taxonomyType := c.Query("type")

	if taxonomyType == "" {
		api.SendResponse(c, errno.ErrTypeEmpty, nil)
		return
	}

	svc := service.New(c.Request.Context())
	termTaxonomy, err := svc.GetTaxonomyList(taxonomyType)
	if err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	api.SendResponse(c, nil, termTaxonomy)
}
