package taxonomy

import (
	Response "puti/handler"
	"puti/pkg/errno"
	"puti/service"

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
