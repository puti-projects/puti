package user

import (
	"strconv"

	Response "puti/handler"
	"puti/model"
	"puti/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Delete deletes the user by id
func Delete(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))

	if err := model.DeleteUser(uint64(userID)); err != nil {
		Response.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	Response.SendResponse(c, nil, nil)
}
