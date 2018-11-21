package user

import (
	"strconv"

	Response "puti/handler"
	"puti/model"
	"puti/pkg/errno"
	"puti/service"
	"puti/utils"

	"github.com/gin-gonic/gin"
)

// savePath defines the use avatar save path
const savePath string = "/uploads/users/"

// Avatar saves the upload avatar for user
func Avatar(c *gin.Context) {
	userID := c.PostForm("userId")
	file, _ := c.FormFile("img")

	fileExt := utils.GetFileExt(file)
	newFileName := "user_" + userID + fileExt

	// Upload the file to specific dst.
	pathName := savePath + newFileName
	dst := "." + pathName
	if err := c.SaveUploadedFile(file, dst); err != nil {
		Response.SendResponse(c, errno.ErrSaveAvatar, nil)
		return
	}

	// update user info for avatar
	ID, err := strconv.Atoi(userID)
	if err != nil {
		Response.SendResponse(c, errno.ErrSaveAvatar, nil)
		return
	}
	user := &model.UserModel{
		Model: model.Model{ID: uint64(ID)},

		Avatar: pathName,
	}

	// Update changed fields.
	if err := service.UpdateUserAvatar(user); err != nil {
		Response.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	Response.SendResponse(c, nil, nil)
}
