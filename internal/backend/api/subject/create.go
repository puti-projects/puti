package subject

import (
	"github.com/puti-projects/puti/internal/backend/api"
	"github.com/puti-projects/puti/internal/backend/service"
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

	// check params
	if err := checkCreateParam(&r); err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	if err := service.CreateSubject(&r); err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	api.SendResponse(c, nil, nil)
}

func checkCreateParam(r *service.SubjectCreateRequest) error {
	if r.Name == "" {
		return errno.New(errno.ErrValidation, nil).Add("name is empty.")
	}

	if r.Slug == "" {
		r.Slug = r.Name
	}

	if ifExist := service.CheckSubjectNameExist(0, r.Name); ifExist == true {
		return errno.New(errno.ErrTaxonomyNameExist, nil).Add(r.Name)
	}

	return nil
}
