package taxonomy

import (
	"strconv"

	Response "puti/handler"
	"puti/model"
	"puti/pkg/errno"
	"puti/service"
	"puti/util"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

type UpdateRequest struct {
	ID          uint64 `form:"id"`
	Name        string `form:"name"`
	Slug        string `form:"slug"`
	Description string `form:"description"`
	ParentID    uint64 `form:"parentId"`
	Taxonomy    string `form:"taxonomy"` // category or tag
}

// Update update taxonomy include category or tag
func Update(c *gin.Context) {
	log.Info("Update function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

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

	return nil
}
