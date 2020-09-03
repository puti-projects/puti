package subject

import (
	Response "github.com/puti-projects/puti/internal/backend/handler"
	"github.com/puti-projects/puti/internal/backend/service"
	"github.com/puti-projects/puti/internal/model"
	"github.com/puti-projects/puti/internal/pkg/errno"
	"github.com/puti-projects/puti/internal/pkg/logger"
	"github.com/puti-projects/puti/internal/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// UpdateRequest struct bind to update subject
type UpdateRequest struct {
	ID          uint64 `json:"ID"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	ParentID    uint64 `json:"parent_id"`
	CoverImage  uint64 `json:"cover_image"`
	Description string `json:"description"`
}

// Update update the subject by ID
func Update(c *gin.Context) {
	logger.Info("Subject update function called.", zap.String("X-request-Id", utils.GetReqID(c)))

	var r UpdateRequest
	if err := c.Bind(&r); err != nil {
		Response.SendResponse(c, errno.ErrBind, nil)
		return
	}

	// check params
	if err := r.checkParam(); err != nil {
		Response.SendResponse(c, err, nil)
		return
	}

	subject := &model.SubjectModel{
		Model: model.Model{ID: r.ID},

		ParentID:    r.ParentID,
		Name:        r.Name,
		Slug:        r.Slug,
		Description: r.Description,
		CoverImage:  r.CoverImage,
	}

	// Update changed fields.
	if err := service.UpdateSubject(subject); err != nil {
		Response.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	Response.SendResponse(c, nil, nil)
}

func (r *UpdateRequest) checkParam() error {
	if r.Name == "" {
		return errno.New(errno.ErrValidation, nil).Add("name is empty.")
	}

	if r.Slug == "" {
		r.Slug = r.Name
	}

	if ifExist := model.SubjectCheckNameExistWhileUpdate(r.ID, r.Name); ifExist == true {
		return errno.New(errno.ErrTaxonomyNameExist, nil).Add(r.Name)
	}

	return nil
}
