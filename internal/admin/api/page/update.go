package page

import (
	"strconv"

	"github.com/puti-projects/puti/internal/admin/api"
	"github.com/puti-projects/puti/internal/admin/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Update update page info handler
// Delete and restore info are also in this function and it depends on the 'status'
func Update(c *gin.Context) {
	// Get page id
	ID, _ := strconv.Atoi(c.Param("id"))

	var r service.PageUpdateRequest
	if err := c.ShouldBind(&r); err != nil {
		api.SendResponse(c, errno.ErrBind, nil)
		return
	}

	pageID := uint64(ID)

	svc := service.New(c.Request.Context())

	// check params
	if err := checkUpdateParam(&svc, &r, pageID); err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	if r.Status == "deleted" {
		if err := svc.TrashPost(pageID); err != nil {
			api.SendResponse(c, err, nil)
			return
		}
	} else if r.Status == "restore" {
		if err := svc.RestorePost(pageID); err != nil {
			api.SendResponse(c, err, nil)
			return
		}
	} else {
		if err := svc.UpdatePage(&r); err != nil {
			api.SendResponse(c, err, nil)
			return
		}
	}

	api.SendResponse(c, nil, nil)
	return
}

func checkUpdateParam(svc *service.Service, r *service.PageUpdateRequest, pageID uint64) error {
	if r.ID == 0 {
		return errno.New(errno.ErrValidation, nil).Add("need id.")
	}

	if r.ID != pageID {
		return errno.New(errno.ErrValidation, nil).Add("error id.")
	}

	if r.Status == "" {
		return errno.New(errno.ErrValidation, nil).Add("need status.")
	}

	if r.Status != "publish" && r.Status != "draft" && r.Status != "deleted" && r.Status != "restore" {
		return errno.New(errno.ErrValidation, nil).Add("error status.")
	}

	if r.Status == "publish" || r.Status == "draft" {
		if isExist := svc.CheckPageSlugExist(r.ID, r.Slug); isExist == true {
			return errno.New(errno.ErrSlugExist, nil)
		}
	}

	return nil
}
