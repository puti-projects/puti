package taxonomy

import (
	"strconv"

	"github.com/puti-projects/puti/internal/admin/api"
	"github.com/puti-projects/puti/internal/admin/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Update update taxonomy include category or tag
func Update(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id")) // term id

	var r service.TaxonomyUpdateRequest
	if err := c.ShouldBind(&r); err != nil {
		api.SendResponse(c, errno.ErrBind, nil)
		return
	}

	termID := uint64(ID)
	if err := checkUpdateParam(&r, termID); err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	if err := service.UpdateTaxonomy(&r, termID); err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	api.SendResponse(c, nil, nil)
	return
}

func checkUpdateParam(r *service.TaxonomyUpdateRequest, termID uint64) error {
	if r.ID != termID {
		return errno.New(errno.ErrValidation, nil).Add("error id.")
	}

	if r.Taxonomy != "category" && r.Taxonomy != "tag" {
		return errno.New(errno.ErrValidation, nil).Add("error taxonomy.")
	}

	if r.Taxonomy == "category" && r.ID == r.ParentID {
		return errno.New(errno.ErrTaxonomyParentCanNotSelf, nil)
	}

	if r.Name == "" {
		return errno.New(errno.ErrValidation, nil).Add("name is empty.")
	}

	if r.Slug == "" {
		r.Slug = r.Name
	}

	return nil
}
