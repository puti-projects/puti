package taxonomy

import (
	"github.com/puti-projects/puti/internal/admin/api"
	"github.com/puti-projects/puti/internal/admin/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Create create term taxonomy handler
func Create(c *gin.Context) {
	var r service.TaxonomyCreateRequest
	if err := c.Bind(&r); err != nil {
		api.SendResponse(c, errno.ErrBind, nil)
		return
	}

	svc := service.New(c.Request.Context())
	// check params
	if err := checkCreateParam(&r, &svc); err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	if err := svc.CreateTaxonomy(&r); err != nil {
		api.SendResponse(c, err, nil)
	}

	api.SendResponse(c, nil, nil)
}

func checkCreateParam(r *service.TaxonomyCreateRequest, svc *service.Service) error {
	if r.Taxonomy != "category" && r.Taxonomy != "tag" {
		return errno.New(errno.ErrValidation, nil).Add("error taxonomy.")
	}

	if r.Name == "" {
		return errno.New(errno.ErrValidation, nil).Add("name is empty.")
	}

	if ifExist := svc.CheckTaxonomyNameExist(r.Name, r.Taxonomy); ifExist == true {
		return errno.New(errno.ErrTaxonomyNameExist, nil).Add(r.Name)
	}

	return nil
}
