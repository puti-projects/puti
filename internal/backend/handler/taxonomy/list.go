package taxonomy

import (
	Response "github.com/puti-projects/puti/internal/backend/handler"
	"github.com/puti-projects/puti/internal/backend/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// List taxonomy tree
func List(c *gin.Context) {
	taxonomyType := c.Query("type")

	if taxonomyType == "" {
		Response.SendResponse(c, errno.ErrTypeEmpty, nil)
		return
	}

	termTaxonomy, err := service.GetTaxonomyList(taxonomyType)
	if err != nil {
		Response.SendResponse(c, err, nil)
		return
	}

	Response.SendResponse(c, nil, termTaxonomy)
}
