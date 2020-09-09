package subject

import (
	Response "github.com/puti-projects/puti/internal/backend/handler"
	"github.com/puti-projects/puti/internal/model"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// CreateRequest struct bind to create subject
type CreateRequest struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	ParentID    uint64 `json:"parent_id"`
	CoverImage  uint64 `json:"cover_image"`
	Description string `json:"description"`
}

// Create create a new subject
func Create(c *gin.Context) {
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

	s := model.SubjectModel{
		ParentID:    r.ParentID,
		Name:        r.Name,
		Slug:        r.Slug,
		Description: r.Description,
		CoverImage:  r.CoverImage,
	}
	if err := s.Create(); err != nil {
		Response.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	Response.SendResponse(c, nil, nil)
}

func (r *CreateRequest) checkParam() error {
	if r.Name == "" {
		return errno.New(errno.ErrValidation, nil).Add("name is empty.")
	}

	if r.Slug == "" {
		r.Slug = r.Name
	}

	if ifExist := model.SubjectCheckNameExistWhileCreate(r.Name); ifExist == true {
		return errno.New(errno.ErrTaxonomyNameExist, nil).Add(r.Name)
	}

	return nil
}
