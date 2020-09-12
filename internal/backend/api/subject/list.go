package subject

import (
	"github.com/puti-projects/puti/internal/backend/api"
	"github.com/puti-projects/puti/internal/backend/service"

	"github.com/gin-gonic/gin"
)

// List  subject list handler
func List(c *gin.Context) {
	subject, err := service.GetSubjectList()
	if err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	api.SendResponse(c, nil, subject)
}
