package subject

import (
	"github.com/puti-projects/puti/internal/backend/api"
	"github.com/puti-projects/puti/internal/backend/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Update update the subject by ID
func Update(c *gin.Context) {
	var r service.SubjectUpdateRequest
	if err := c.Bind(&r); err != nil {
		api.SendResponse(c, errno.ErrBind, nil)
		return
	}

	// check params
	if err := checkUpdateParam(&r); err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	// Update changed fields.
	if err := service.UpdateSubject(&r); err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	api.SendResponse(c, nil, nil)
}

func checkUpdateParam(r *service.SubjectUpdateRequest) error {
	if r.Name == "" {
		return errno.New(errno.ErrValidation, nil).Add("name is empty.")
	}

	if r.Slug == "" {
		r.Slug = r.Name
	}

	if ifExist := service.CheckSubjectNameExist(r.ID, r.Name); ifExist == true {
		return errno.New(errno.ErrTaxonomyNameExist, nil).Add(r.Name)
	}

	return nil
}
