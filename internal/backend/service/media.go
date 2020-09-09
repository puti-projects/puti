package service

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"mime/multipart"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/puti-projects/puti/internal/dao"
	"github.com/puti-projects/puti/internal/model"
	"github.com/puti-projects/puti/internal/pkg/config"
	"github.com/puti-projects/puti/internal/pkg/errno"
	"github.com/puti-projects/puti/internal/utils"
)

// MediaUploadResponse is the upload media request's response struct
type MediaUploadResponse struct {
	ID  uint64 `json:"id"`
	URL string `json:"url"`
}

// MediaUpdateRequest is the update media request params struct
type MediaUpdateRequest struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
}

// MediaListRequest is the media list request struct
type MediaListRequest struct {
	Limit int `form:"limit"`
	Page  int `form:"page"`
}

// MediaListResponse returns total number of media and current page of media
type MediaListResponse struct {
	TotalCount int64        `json:"totalCount"`
	MediaList  []*MediaInfo `json:"mediaList"`
}

// MediaInfo is the media struct for media list
type MediaInfo struct {
	ID         uint64 `json:"id"`
	Title      string `json:"title"`
	GUID       string `json:"url"`
	Type       string `json:"type"`
	UploadTime string `json:"upload_time"`
}

// MediaDetail media info in detail inclue Mediainfo struct
type MediaDetail struct {
	MediaInfo
	Slug        string `json:"slug"`
	Description string `json:"description"`
}

// MediaList media list
type MediaList struct {
	Lock  *sync.Mutex
	IDMap map[uint64]*MediaInfo
}

// UploadMedia upload media and save record
func UploadMedia(c *gin.Context, userID, usage string, file *multipart.FileHeader) (ID uint64, GUID string, err error) {
	fileNameWithoutExt, fileExt, pathName, dst, err := getFileSavePath(usage, file)
	if err != nil {
		return
	}

	// Upload the file to specific dst.
	if err = c.SaveUploadedFile(file, dst); err != nil {
		err = errno.New(errno.ErrUploadFile, err)
		return
	}

	uID, err := strconv.Atoi(userID)
	if err != nil {
		return
	}

	ID, GUID, err = dao.Engine.CreateMedia(uID, file.Filename, fileNameWithoutExt, fileExt, pathName, usage)
	return
}

// getFileSavePath general the hole uri for the file
func getFileSavePath(usage string, file *multipart.FileHeader) (
	fileNameWithoutExt string,
	fileExt string,
	pathName string,
	dst string,
	err error,
) {
	// General the save path by upload time
	savePath, err := getSavePath(usage)
	if err != nil {
		return
	}

	// set variables
	fileExt = utils.GetFileExt(file)
	fileNameWithoutExt = strings.TrimSuffix(file.Filename, fileExt)
	unixTime := time.Now().Unix()

	// set buf string
	buf := bytes.NewBufferString(fileNameWithoutExt)
	buf.Write([]byte(strconv.FormatInt(unixTime, 10))) // add a time string
	// md5 encode
	h := md5.New()
	h.Write([]byte(buf.String())) // encode the buf.String()
	newFileName := hex.EncodeToString(h.Sum(nil))

	// final save path with file name
	pathName = savePath + newFileName + fileExt
	dst = "." + pathName

	return fileNameWithoutExt, fileExt, pathName, dst, nil
}

// getSavePath general the hole uri by upload time
func getSavePath(usage string) (string, error) {
	savePathURI := config.UploadPath

	// for cover picture
	if usage == "cover" {
		// check cover path exist
		coverPath := fmt.Sprintf(".%s%s", savePathURI, "cover")
		if err := utils.CheckPathAndCreate(coverPath); err != nil {
			return "", err
		}

		savePath := fmt.Sprintf("%s%s/", savePathURI, "cover")
		return savePath, nil
	}

	// for common picture
	now := time.Now()
	// handel year path
	year := utils.GetFormatTime(&now, "2006")
	yearPath := fmt.Sprintf(".%s%s", savePathURI, year)
	if err := utils.CheckPathAndCreate(yearPath); err != nil {
		return "", err
	}

	// handle month path
	month := utils.GetFormatTime(&now, "01")
	monthPath := fmt.Sprintf(".%s%s/%s", savePathURI, year, month)
	if err := utils.CheckPathAndCreate(monthPath); err != nil {
		return "", err
	}

	savePath := fmt.Sprintf("%s%s/%s/", savePathURI, year, month)
	return savePath, nil
}

// GetMediaByID get media by ID
func GetMediaByID(ID uint64) (*model.Media, error) {
	media, err := dao.Engine.GetMediaByID(ID)
	if err != nil {
		return nil, err
	}

	return media, nil
}

// GetMedia return media info if database select success
func GetMediaDetail(ID uint64) (*MediaDetail, error) {
	media, err := dao.Engine.GetMediaByID(ID)
	if err != nil {
		return nil, err
	}

	mediaInfo := &MediaDetail{
		MediaInfo: MediaInfo{
			ID:         media.ID,
			Title:      media.Title,
			GUID:       media.GUID,
			Type:       media.Type,
			UploadTime: utils.GetFormatTime(&media.CreatedAt, "2006-01-02 15:04:05"),
		},
		Slug:        media.Slug,
		Description: media.Description,
	}

	return mediaInfo, nil
}

// ListMedia returns current page media list and the total number of media
func ListMedia(limit, page int) ([]*MediaInfo, int64, error) {
	infos := make([]*MediaInfo, 0)

	medias, count, err := dao.Engine.ListMedia(limit, page)
	if err != nil {
		return nil, count, err
	}

	ids := []uint64{}
	for _, media := range medias {
		ids = append(ids, media.ID)
	}

	wg := sync.WaitGroup{}
	mediaList := MediaList{
		Lock:  new(sync.Mutex),
		IDMap: make(map[uint64]*MediaInfo, len(medias)),
	}

	finished := make(chan bool, 1)

	// Improve query efficiency in parallel
	for _, u := range medias {
		wg.Add(1)
		go func(u *model.Media) {
			defer wg.Done()

			mediaList.Lock.Lock()
			defer mediaList.Lock.Unlock()
			mediaList.IDMap[u.ID] = &MediaInfo{
				ID:         u.ID,
				Title:      u.Title,
				GUID:       u.GUID,
				Type:       u.Type,
				UploadTime: u.CreatedAt.In(config.TimeLoc()).Format("2006-01-02 15:04"),
			}
		}(u)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	<-finished

	for _, id := range ids {
		infos = append(infos, mediaList.IDMap[id])
	}

	return infos, count, nil
}

// UpdateMedia update media info
func UpdateMedia(r *MediaUpdateRequest, userID int) (err error) {
	if err := dao.Engine.UpdateMedia(uint64(userID), r.Title, r.Slug, r.Description); err != nil {
		return errno.New(errno.ErrDatabase, err)
	}

	return nil
}

// DeleteMedia delete media
func DeleteMedia(userID uint64) error {
	if err := dao.Engine.DeleteMediaByID(userID); err != nil {
		return errno.New(errno.ErrDatabase, err)
	}

	return nil
}
