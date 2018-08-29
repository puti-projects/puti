package user

import (
	"mime"
	"mime/multipart"
	"path/filepath"
	"strconv"
	"strings"

	Response "puti/handler"
	"puti/model"
	"puti/pkg/errno"
	"puti/service"

	"github.com/gin-gonic/gin"
)

// savePath defines the use avatar save path
const savePath string = "/upload/users/"

// Avatar saves the upload avatar for user
func Avatar(c *gin.Context) {
	userID := c.PostForm("userId")
	file, _ := c.FormFile("img")

	fileExt := getFileExt(file)
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

// getFileExt returns the file name extension
// The extension is the suffix beginning at the final dot
func getFileExt(file *multipart.FileHeader) string {
	var ext string
	// get by Ext func first
	ext = filepath.Ext(file.Filename)
	if ext == "" {
		// get by content-type
		typ := file.Header.Get("Content-Type")
		exts, _ := mime.ExtensionsByType(typ)
		if 0 < len(exts) {
			ext = exts[0]
		} else {
			ext = "." + strings.Split(typ, "/")[1]
		}
	}

	return ext
}
