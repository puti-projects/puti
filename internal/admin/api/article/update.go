package article

import (
	"strconv"

	"github.com/puti-projects/puti/internal/admin/api"
	"github.com/puti-projects/puti/internal/admin/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Update update article
// Delete and restore article are also in this function and it depends on the 'status'
func Update(c *gin.Context) {
	// Get article id
	ID, _ := strconv.Atoi(c.Param("id"))

	var r service.ArticleUpdateRequest
	if err := c.ShouldBind(&r); err != nil {
		api.SendResponse(c, errno.ErrBind, nil)
		return
	}

	articleID := uint64(ID)

	// check params
	if err := checkUpdateParam(&r, articleID); err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	if r.Status == "deleted" {
		if err := service.TrashPost(articleID); err != nil {
			api.SendResponse(c, err, nil)
			return
		}
	} else if r.Status == "restore" {
		if err := service.RestorePost(articleID); err != nil {
			api.SendResponse(c, err, nil)
			return
		}
	} else {
		if err := service.UpdateArticle(&r); err != nil {
			api.SendResponse(c, err, nil)
			return
		}
	}

	api.SendResponse(c, nil, nil)
	return
}

func checkUpdateParam(r *service.ArticleUpdateRequest, articleID uint64) error {
	if r.ID == 0 {
		return errno.New(errno.ErrValidation, nil).Add("need id.")
	}

	if r.ID != articleID {
		return errno.New(errno.ErrValidation, nil).Add("error id.")
	}

	if r.Status == "" {
		return errno.New(errno.ErrValidation, nil).Add("need status.")
	}

	if r.Status != "publish" && r.Status != "draft" && r.Status != "deleted" && r.Status != "restore" {
		return errno.New(errno.ErrValidation, nil).Add("error status.")
	}

	return nil
}
