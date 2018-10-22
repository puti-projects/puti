package taxonomy

import (
	"fmt"
	Response "puti/handler"
	"puti/model"
	"puti/pkg/errno"
	"puti/util"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// CreateRequest struct to crate taxonomy include category and tag
type CreateRequest struct {
	Name        string `form:"name"`
	Slug        string `form:"slug"`
	Description string `form:"description"`
	ParentID    uint64 `form:"parentId"`
	Taxonomy    string `form:"taxonomy"` // category or tag
}

// Create create txonomy
func Create(c *gin.Context) {
	log.Info("Category Create function called.", lager.Data{"X-request-Id": util.GetReqID(c)})

	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		Response.SendResponse(c, errno.ErrBind, nil)
		return
	}

	// check params
	if err := r.checkParam(); err != nil {
		Response.SendResponse(c, err, nil)
		return
	}

	level, err := model.GetTaxonomyLevel(r.ParentID, r.Taxonomy)
	if err != nil {
		Response.SendResponse(c, errno.ErrTaxonomyParentId, nil)
		return
	}

	t := model.TermModel{
		Name:        r.Name,
		Slug:        r.Slug,
		Description: r.Description,
		Count:       0,
	}
	if err := t.Create(); err != nil {
		Response.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	tm := model.TermTaxonomyModel{
		Term:         t,
		TermID:       t.ID,
		ParentTermID: r.ParentID,
		Level:        level,
		Taxonomy:     r.Taxonomy,
		TermGroup:    0,
	}

	if err := tm.Create(); err != nil {
		fmt.Print(err)
		Response.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	Response.SendResponse(c, nil, nil)
}

func (r *CreateRequest) checkParam() error {
	if r.Name == "" {
		return errno.New(errno.ErrValidation, nil).Add("name is empty.")
	}

	if r.Taxonomy == "" {
		return errno.New(errno.ErrValidation, nil).Add("taxonomy is empty.")
	}

	if ifExist := model.TaxonomyCheckNameExist(r.Name, r.Taxonomy); ifExist == true {
		return errno.New(errno.ErrTaxonomyNameExist, nil).Add(r.Name)
	}

	return nil
}
