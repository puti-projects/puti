package subject

import (
	"github.com/puti-projects/puti/internal/admin/api"
	"github.com/puti-projects/puti/internal/admin/service"
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
	svc := service.New(c.Request.Context())

	// check params
	if err := checkUpdateParam(&svc, &r); err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	// Update changed fields.
	if err := svc.UpdateSubject(&r); err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	api.SendResponse(c, nil, nil)
}

func checkUpdateParam(svc *service.Service, r *service.SubjectUpdateRequest) error {
	if r.Name == "" {
		return errno.New(errno.ErrValidation, nil).Add("name is empty.")
	}

	if r.Slug == "" {
		r.Slug = r.Name
	}

	if ifExist := svc.CheckSubjectNameExist(r.ID, r.Name); ifExist == true {
		return errno.New(errno.ErrTaxonomyNameExist, nil).Add(r.Name)
	}

	return nil
}
