package service

import (
	"sync"

	"github.com/puti-projects/puti/internal/common/config"
	"github.com/puti-projects/puti/internal/common/model"
)

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

// GetMedia return media info if database select success
func GetMedia(id uint64) (*MediaDetail, error) {
	m, err := model.GetMediaByID(id)
	if err != nil {
		return nil, err
	}

	mediaInfo := &MediaDetail{
		MediaInfo: MediaInfo{
			ID:         m.ID,
			Title:      m.Title,
			GUID:       m.GUID,
			Type:       m.Type,
			UploadTime: m.CreatedAt.In(config.TimeLoc()).Format("2006-01-02 15:04:05"),
		},
		Slug:        m.Slug,
		Description: m.Description,
	}

	return mediaInfo, nil
}

// ListMedia returns media list and media count
func ListMedia(limit, page int) ([]*MediaInfo, uint64, error) {
	infos := make([]*MediaInfo, 0)
	medias, count, err := model.ListMedia(limit, page)
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

	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	// Improve query efficiency in parallel
	for _, u := range medias {
		wg.Add(1)
		go func(u *model.MediaModel) {
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

	select {
	case <-finished:
	case err := <-errChan:
		return nil, count, err
	}

	for _, id := range ids {
		infos = append(infos, mediaList.IDMap[id])
	}

	return infos, count, nil
}

// UpdateMedia update media info
func UpdateMedia(media *model.MediaModel) (err error) {
	// Get old media info
	oldMedia, err := model.GetMediaByID(media.ID)
	if err != nil {
		return err
	}

	// Set new status values
	oldMedia.Title = media.Title
	oldMedia.Slug = media.Slug
	oldMedia.Description = media.Description

	if err = oldMedia.Update(); err != nil {
		return err
	}

	return nil
}
