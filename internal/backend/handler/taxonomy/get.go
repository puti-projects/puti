package taxonomy

import (
	Response "github.com/puti-projects/puti/internal/backend/handler"
	"github.com/puti-projects/puti/internal/backend/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Get gets the term taxonomy info by term_id
func Get(c *gin.Context) {
	termID := c.Param("id")

	term, err := service.GetTaxonomyInfo(termID)
	if err != nil {
		Response.SendResponse(c, errno.ErrTermNotFount, nil)
		return
	}

	Response.SendResponse(c, nil, term)
}
