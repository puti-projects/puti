package taxonomy

import (
	"strconv"

	"github.com/puti-projects/puti/internal/admin/api"
	"github.com/puti-projects/puti/internal/admin/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Delete delete the taxonomy directly without soft delete
func Delete(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	taxonomyType := c.Query("taxonomy") // TODO

	termID := uint64(ID)

	// check
	if err := checkIfCanDelete(termID, taxonomyType); err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	if err := service.DeleteTaxonomy(termID, taxonomyType); err != nil {
		api.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	api.SendResponse(c, nil, nil)
}

func checkIfCanDelete(termID uint64, taxonomyType string) error {
	if taxonomyType != "category" && taxonomyType != "tag" {
		return errno.New(errno.ErrValidation, nil).Add("error taxonomy.")
	}

	ifHasChild, err := service.IfTaxonomyHasChild(termID, taxonomyType)
	if err != nil {
		return err
	}
	if ifHasChild == true {
		return errno.New(errno.ErrValidation, nil).Add("taxonomy has children and can not be deleted")

	}

	return nil
}
