package subject

import (
	"github.com/puti-projects/puti/internal/admin/api"
	"github.com/puti-projects/puti/internal/admin/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Create create subject handler
func Create(c *gin.Context) {
	var r service.SubjectCreateRequest
	if err := c.Bind(&r); err != nil {
		api.SendResponse(c, errno.ErrBind, nil)
		return
	}

	svc := service.New(c.Request.Context())

	// check params
	if err := checkCreateParam(&svc, &r); err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	if err := svc.CreateSubject(&r); err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	api.SendResponse(c, nil, nil)
}

func checkCreateParam(svc *service.Service, r *service.SubjectCreateRequest) error {
	if r.Name == "" {
		return errno.New(errno.ErrValidation, nil).Add("name is empty.")
	}

	if r.Slug == "" {
		r.Slug = r.Name
	}

	if ifExist := svc.CheckSubjectNameExist(0, r.Name); ifExist == true {
		return errno.New(errno.ErrTaxonomyNameExist, nil).Add(r.Name)
	}

	return nil
}
