package taxonomy

import (
	"strconv"

	Response "github.com/puti-projects/puti/internal/backend/handler"
	"github.com/puti-projects/puti/internal/common/model"
	"github.com/puti-projects/puti/internal/pkg/errno"
	"github.com/puti-projects/puti/internal/backend/service"
	"github.com/puti-projects/puti/internal/common/utils"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// UpdateRequest param struct to update taxonomy include category and tag
type UpdateRequest struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	ParentID    uint64 `json:"parentId"`
	Taxonomy    string `json:"taxonomy"` // category or tag
}

// Update update taxonomy include category or tag
func Update(c *gin.Context) {
	log.Info("Update function called.", lager.Data{"X-Request-Id": utils.GetReqID(c)})

	// Get term id
	termID, _ := strconv.Atoi(c.Param("id"))

	var r UpdateRequest

	if err := c.ShouldBind(&r); err != nil {
		Response.SendResponse(c, errno.ErrBind, nil)
		return
	}

	theTermID := uint64(termID)

	// check params
	if err := r.checkParam(theTermID); err != nil {
		Response.SendResponse(c, err, nil)
		return
	}

	term := &model.TermModel{
		ID:          theTermID,
		Name:        r.Name,
		Slug:        r.Slug,
		Description: r.Description,
	}

	termTaxonomy := &model.TermTaxonomyModel{
		Term:         *term,
		TermID:       theTermID,
		ParentTermID: r.ParentID,
	}

	// Update changed fields.
	if err := service.UpdateTaxonomy(term, termTaxonomy, r.Taxonomy); err != nil {
		Response.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	Response.SendResponse(c, nil, nil)
	return
}

// checkUpdateParam check params if correct
func (r *UpdateRequest) checkParam(termID uint64) error {
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
