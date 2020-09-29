package article

import (
	"github.com/puti-projects/puti/internal/admin/api"
	"github.com/puti-projects/puti/internal/admin/service"
	"github.com/puti-projects/puti/internal/pkg/errno"
	"github.com/puti-projects/puti/internal/pkg/token"

	"github.com/gin-gonic/gin"
)

// Create create article (published or draft) handler
func Create(c *gin.Context) {
	// get token and parse
	t := c.Query("token")
	userContext, err := token.ParseToken(t)

	var r service.ArticleCreateRequest
	if err := c.Bind(&r); err != nil {
		api.SendResponse(c, errno.ErrBind, nil)
		return
	}

	// check params
	if err := checkCreateParam(&r); err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	rsp, err := service.CreateArticle(&r, userContext.ID)
	if err != nil {
		api.SendResponse(c, errno.ErrArticleCreateFailed, nil)
		return
	}

	api.SendResponse(c, nil, rsp)
}

func checkCreateParam(r *service.ArticleCreateRequest) error {
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

	return nil
}
