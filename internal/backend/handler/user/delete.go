package user

import (
	"strconv"

	Response "github.com/puti-projects/puti/internal/backend/handler"
	"github.com/puti-projects/puti/internal/backend/service"

	"github.com/gin-gonic/gin"
)

// Delete user delete handler
func Delete(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))

	if err := service.DeleteUser(userID); err != nil {
		Response.SendResponse(c, err, nil)
		return
	}

	Response.SendResponse(c, nil, nil)
}
