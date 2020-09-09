package subject

import (
	"strconv"

	Response "github.com/puti-projects/puti/internal/backend/handler"
	"github.com/puti-projects/puti/internal/backend/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Delete delete the taxonomy directly without soft delete
func Delete(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	subjectID := uint64(ID)

	// check
	if err := checkIfCanDelete(subjectID); err != nil {
		Response.SendResponse(c, err, nil)
		return
	}

	if err := service.DeleteSubject(subjectID); err != nil {
		Response.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	Response.SendResponse(c, nil, nil)
}

func checkIfCanDelete(subjectID uint64) error {
	if ifHasChild := service.IfSubjectHasChild(subjectID); ifHasChild == true {
		return errno.New(errno.ErrValidation, nil).Add("subject has children and can not be deleted")
	}

	return nil
}
