package service

import (
	"errors"
	"strings"
	"sync"

	"github.com/puti-projects/puti/internal/admin/dao"
	"github.com/puti-projects/puti/internal/model"
	"github.com/puti-projects/puti/internal/pkg/errno"
	"github.com/puti-projects/puti/internal/utils"

	"gorm.io/gorm"
)

// KnowledgeCreateRequest struct bind to create knowledge
type KnowledgeCreateRequest struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Type        string `json:"type"`
	Description string `json:"description"`
	CoverImage  uint64 `json:"cover_image"`
}

// KnowledgeUpdateRequest struct bind to update knowledge
type KnowledgeUpdateRequest struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Type        string `json:"type"`
	Description string `json:"description"`
	CoverImage  uint64 `json:"cover_image"`
}

// KnowledgeList knowledge list
type KnowledgeList struct {
	Lock    *sync.Mutex
	TypeMap map[string][]*KnowledgeInfo
}

// KnowledgeInfo struct of knowledge list for output
type KnowledgeInfo struct {
	ID             uint64 `json:"id"`
	Name           string `json:"name"`
	Slug           string `json:"slug"`
	Description    string `json:"description"`
	CoverImageID   uint64 `json:"cover_image_id"`
	CoverImageName string `json:"cover_image_name"`
	CoverImageURL  string `json:"cover_image_url"`
	LastUpdated    string `json:"last_updated"`
	CreatedTime    string `json:"created_time"`
}

// CheckKnowledgeType check knowledge type
func (svc *Service) CheckKnowledgeType(knowledgeType string) bool {
	if knowledgeType != model.KnowledgeTypeDoc && knowledgeType != model.KnowledgeTypeNote {
		return false
	}

	return true
}

// CreateKnowledge create knowledge base
func (svc *Service) CreateKnowledge(r *KnowledgeCreateRequest) error {
	k := &model.Knowledge{
		Name:        r.Name,
		Slug:        r.Slug,
		Type:        r.Type,
		Description: r.Description,
		CoverImage:  r.CoverImage,
	}

	if err := svc.dao.CreateKnowledge(k); err != nil {
		return errno.New(errno.ErrDatabase, err)
	}

	return nil
}

// UpdateKnowledge update knowledge base info
func (svc Service) UpdateKnowledge(r *KnowledgeUpdateRequest) error {
	k := &model.Knowledge{
		Model: model.Model{ID: r.ID},

		Name:        strings.TrimSpace(r.Name),
		Slug:        strings.TrimSpace(r.Slug),
		Type:        r.Type,
		Description: strings.TrimSpace(r.Description),
		CoverImage:  r.CoverImage,
	}

	if err := svc.dao.UpdateKnowledge(k); err != nil {
		return errno.New(errno.ErrDatabase, err)
	}

	// update finished. clean cache.
	svc.CleanCacheAfterUpdateKnowledge(k.Slug)
	return nil
}

// GetKnowledgeList get knowledge base list
func (svc *Service) GetKnowledgeList() (map[string][]*KnowledgeInfo, error) {
	ks, err := svc.dao.GetAllKnowledgeList()
	if err != nil {
		return nil, errno.New(errno.ErrDatabase, err)
	}

	wg := sync.WaitGroup{}
	finished := make(chan bool, 1)
	kList := KnowledgeList{
		Lock:    new(sync.Mutex),
		TypeMap: make(map[string][]*KnowledgeInfo, len(ks)),
	}

	for _, v := range ks {
		wg.Add(1)

		go func(ki *dao.KnowledgeInfo) {
			defer wg.Done()

			kList.Lock.Lock()
			defer kList.Lock.Unlock()
			kList.TypeMap[ki.Type] = append(kList.TypeMap[ki.Type], &KnowledgeInfo{
				ID:             ki.ID,
				Name:           ki.Name,
				Slug:           ki.Slug,
				Description:    ki.Description,
				CoverImageID:   ki.CoverImageID,
				CoverImageName: ki.CoverImageName,
				CoverImageURL:  ki.CoverImageURL,
				LastUpdated:    utils.GetFormatNullTime(&ki.LastUpdated, "2006-01-02 15:04:05"),
				CreatedTime:    utils.GetFormatTime(&ki.CreatedTime, "2006-01-02 15:04:05"),
			})
		}(v)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	<-finished

	return kList.TypeMap, nil
}

// GetKnowledgeInfo get knowledge info by knowledge ID
func (svc *Service) GetKnowledgeInfo(kID uint64) (*KnowledgeInfo, error) {
	k, err := svc.dao.GetKnowledgeByID(kID)
	if err != nil {
		return nil, err
	}

	var coverImageID uint64
	var coverImageName, coverImageURL string
	if k.CoverImage != 0 {
		m, err := svc.GetMediaByID(k.CoverImage)
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			coverImageName = "Unknown File"
		}
		coverImageID = m.ID
		coverImageName = m.Title
		coverImageURL = m.GUID
	}

	kInfo := &KnowledgeInfo{
		ID:             k.ID,
		Name:           k.Name,
		Slug:           k.Slug,
		Description:    k.Description,
		CoverImageID:   coverImageID,
		CoverImageName: coverImageName,
		CoverImageURL:  coverImageURL,
		LastUpdated:    utils.GetFormatNullTime(&k.LastUpdated, "2006-01-02 15:04:05"),
		CreatedTime:    utils.GetFormatTime(&k.CreatedAt, "2006-01-02 15:04:05"),
	}

	return kInfo, nil
}
