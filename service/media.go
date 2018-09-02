package service

import (
	"sync"

	"puti/config"
	"puti/model"
)

// MediaInfo is the media struct for media list
type MediaInfo struct {
	ID         uint64 `json:"id"`
	Title      string `json:"title"`
	GUID       string `json:"url"`
	Type       string `json:"type"`
	UploadTime string `json:"upload_time"`
}

// MediaList media list
type MediaList struct {
	Lock  *sync.Mutex
	IDMap map[uint64]*MediaInfo
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
