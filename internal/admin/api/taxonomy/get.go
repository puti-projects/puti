package taxonomy

import (
	"strconv"

	"github.com/puti-projects/puti/internal/admin/api"
	"github.com/puti-projects/puti/internal/admin/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Get get term taxonomy handler
func Get(c *gin.Context) {
	termID, _ := strconv.Atoi(c.Param("id"))

	svc := service.New(c.Request.Context())
	term, err := svc.GetTaxonomyInfo(uint64(termID))
	if err != nil {
		api.SendResponse(c, errno.ErrTermNotFount, nil)
		return
	}

	api.SendResponse(c, nil, term)
}
