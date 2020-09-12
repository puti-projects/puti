package taxonomy

import (
	"strconv"

	"github.com/puti-projects/puti/internal/backend/api"
	"github.com/puti-projects/puti/internal/backend/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Get get term taxonomy handler
func Get(c *gin.Context) {
	termID, _ := strconv.Atoi(c.Param("id"))

	term, err := service.GetTaxonomyInfo(uint64(termID))
	if err != nil {
		api.SendResponse(c, errno.ErrTermNotFount, nil)
		return
	}

	api.SendResponse(c, nil, term)
}
