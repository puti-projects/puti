package user

import (
	"strconv"

	"github.com/puti-projects/puti/internal/backend/api"
	"github.com/puti-projects/puti/internal/backend/service"

	"github.com/gin-gonic/gin"
)

// Delete user delete handler
func Delete(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))

	if err := service.DeleteUser(userID); err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	api.SendResponse(c, nil, nil)
}
