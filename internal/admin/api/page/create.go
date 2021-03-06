package page

import (
	"github.com/puti-projects/puti/internal/admin/api"
	"github.com/puti-projects/puti/internal/admin/service"
	"github.com/puti-projects/puti/internal/pkg/errno"
	"github.com/puti-projects/puti/internal/pkg/token"

	"github.com/gin-gonic/gin"
)

// Create add new page
func Create(c *gin.Context) {
	// get token and parse
	t := c.Query("token")
	userContext, err := token.ParseToken(t)

	var r service.PageCreateRequest
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

	rsp, err := svc.CreatePage(&r, userContext.ID)
	if err != nil {
		api.SendResponse(c, errno.ErrPageCreateFailed, nil)
		return
	}

	api.SendResponse(c, nil, rsp)
}

func checkCreateParam(svc *service.Service, r *service.PageCreateRequest) error {
	if r.Title == "" {
		return errno.New(errno.ErrValidation, nil).Add("Title can not be empty.")
	}

	if r.Content == "" {
		return errno.New(errno.ErrValidation, nil).Add("Content can not be empty.")
	}

	if r.Status == "" {
		return errno.New(errno.ErrValidation, nil).Add("Status can not be empty.")
	}

	if r.Status != "publish" && r.Status != "draft" {
		return errno.New(errno.ErrValidation, nil).Add("Status is incorrect.")
	}

	if isExist := svc.CheckPageSlugExist(0, r.Slug); isExist == true {
		return errno.New(errno.ErrSlugExist, nil)
	}

	return nil
}
