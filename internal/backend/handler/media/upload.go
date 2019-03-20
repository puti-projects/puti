package media

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	Response "github.com/puti-projects/puti/internal/backend/handler"
	"github.com/puti-projects/puti/internal/common/model"
	"github.com/puti-projects/puti/internal/common/utils"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// UploadResponse is the upload media request's response struct
type UploadResponse struct {
	ID  uint64 `json:"id"`
	URL string `json:"url"`
}

// savePathURI defines the media file save path uri
const savePathURI string = "/uploads/"

// Upload the function handle the file upload
func Upload(c *gin.Context) {
	// get userId and file
	userID := c.PostForm("userId")
	usage := c.DefaultPostForm("usage", "common")
	file, _ := c.FormFile("file")

	// General the save path by upload time
	savePath, err := getSavePath(usage)
	if err != nil {
		Response.SendResponse(c, errno.ErrUploadFile, nil)
		return
	}

	// set variables
	fileExt := utils.GetFileExt(file)
	fileNameWithoutExt := strings.TrimSuffix(file.Filename, fileExt)
	unixTime := time.Now().Unix()

	// set buf string
	buf := bytes.NewBufferString(fileNameWithoutExt)
	buf.Write([]byte(strconv.FormatInt(unixTime, 10))) // add a time string
	// md5 encode
	h := md5.New()
	h.Write([]byte(buf.String())) // encode the buf.String()
	newFileName := hex.EncodeToString(h.Sum(nil))

	// final save path with file name
	pathName := savePath + newFileName + fileExt
	dst := "." + pathName
	// Upload the file to specific dst.
	if err := c.SaveUploadedFile(file, dst); err != nil {
		Response.SendResponse(c, errno.ErrUploadFile, nil)
		return
	}

	uID, err := strconv.Atoi(userID)
	if err != nil {
		Response.SendResponse(c, errno.ErrUploadFile, nil)
		return
	}

	media := &model.MediaModel{
		UserID:   uint64(uID),
		Title:    file.Filename,
		Slug:     fileNameWithoutExt,
		GUID:     pathName,
		MimeType: utils.GetFileMimeTypeByExt(fileExt),
		Usage:    usage,
	}

	// save file info
	if err := media.Create(); err != nil {
		Response.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := UploadResponse{
		ID:  media.ID,
		URL: media.GUID,
	}

	Response.SendResponse(c, nil, rsp)
}

// getSavePath general the hole uri by upload time
func getSavePath(usage string) (string, error) {
	if usage == "cover" {
		coverPath := fmt.Sprintf(".%s%s", savePathURI, "cover")
		coverPathExist, err := utils.PathExists(coverPath)
		if err != nil {
			return "", err
		}
		if !coverPathExist {
			err := os.Mkdir(coverPath, os.ModePerm)
			if err != nil {
				return "", err
			}
		}

		var savePath string
		savePath = fmt.Sprintf("%s%s/", savePathURI, "cover")
		return savePath, nil
	}

	now := time.Now()

	// handel year path
	year := utils.GetFormatTime(&now, "2006")
	yearPath := fmt.Sprintf(".%s%s", savePathURI, year)
	yearExist, err := utils.PathExists(yearPath)
	if err != nil {
		return "", err
	}
	if !yearExist {
		err := os.Mkdir(yearPath, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	// handle month path
	month := utils.GetFormatTime(&now, "01")
	monthPath := fmt.Sprintf(".%s%s/%s", savePathURI, year, month)
	monthExist, err := utils.PathExists(monthPath)
	if err != nil {
		return "", err
	}
	if !monthExist {
		err := os.Mkdir(monthPath, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	var savePath string
	savePath = fmt.Sprintf("%s%s/%s/", savePathURI, year, month)

	return savePath, nil
}
