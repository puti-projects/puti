package taxonomy

import (
	"github.com/puti-projects/puti/internal/backend/api"
	"github.com/puti-projects/puti/internal/backend/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Create create term txonomy handler
func Create(c *gin.Context) {
	var r service.TaxonomyCreateRequest
	if err := c.Bind(&r); err != nil {
		api.SendResponse(c, errno.ErrBind, nil)
		return
	}

	// check params
	if err := checkCreateParam(&r); err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	if err := service.CreateTaxonomy(&r); err != nil {
		api.SendResponse(c, err, nil)
	}

	api.SendResponse(c, nil, nil)
}

func checkCreateParam(r *service.TaxonomyCreateRequest) error {
	if r.Taxonomy != "category" && r.Taxonomy != "tag" {
		return errno.New(errno.ErrValidation, nil).Add("error taxonomy.")
	}

	if r.Name == "" {
		return errno.New(errno.ErrValidation, nil).Add("name is empty.")
	}

	if ifExist := service.CheckTaxonomyNameExist(r.Name, r.Taxonomy); ifExist == true {
		return errno.New(errno.ErrTaxonomyNameExist, nil).Add(r.Name)
	}

	return nil
}
