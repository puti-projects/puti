package taxonomy

import (
	"fmt"

	Response "github.com/puti-projects/puti/internal/backend/handler"
	"github.com/puti-projects/puti/internal/model"
	"github.com/puti-projects/puti/internal/pkg/errno"
	"github.com/puti-projects/puti/internal/pkg/logger"
	"github.com/puti-projects/puti/internal/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateRequest struct to crate taxonomy include category and tag
type CreateRequest struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	ParentID    uint64 `json:"parentId"`
	Taxonomy    string `json:"taxonomy"` // category or tag
}

// Create create txonomy
func Create(c *gin.Context) {
	logger.Info("Category Create function called.", zap.String("X-request-Id", utils.GetReqID(c)))

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
		Response.SendResponse(c, errno.ErrTaxonomyParentID, nil)
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
	if r.Taxonomy != "category" && r.Taxonomy != "tag" {
		return errno.New(errno.ErrValidation, nil).Add("error taxonomy.")
	}

	if r.Name == "" {
		return errno.New(errno.ErrValidation, nil).Add("name is empty.")
	}

	if r.Slug == "" {
		r.Slug = r.Name
	}

	if ifExist := model.TaxonomyCheckNameExist(r.Name, r.Taxonomy); ifExist == true {
		return errno.New(errno.ErrTaxonomyNameExist, nil).Add(r.Name)
	}

	return nil
}
