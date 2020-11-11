package service

import (
	"github.com/puti-projects/puti/internal/admin/dao"
	"github.com/puti-projects/puti/internal/model"
	"github.com/puti-projects/puti/internal/pkg/errno"
	"github.com/puti-projects/puti/internal/pkg/snowflake"
	"strings"
)

// KnowledgeItemCreateRequest struct for binding knowledge item create
type KnowledgeItemCreateRequest struct {
	KnowledgeID uint64 `json:"knowledge_id"`
	Title       string `json:"title"`
	ParentID    uint64 `json:"parent_id"`
}

// KnowledgeItemCreateResponse struct for return after created knowledge item
type KnowledgeItemCreateResponse struct {
	ID            uint64 `json:"knowledge_item_id"`
	Title         string `json:"title"`
	ParentID      uint64 `json:"parent_id"`
	Level         uint64 `json:"level"`
	Index         uint64 `json:"index"`
	Content       string `json:"content"`
	Version       int64  `json:"version"`
	VersionStatus uint8  `json:"version_status"`
}

//type KnowledgeItemList struct {
//	Lock     *sync.Mutex
//	ItemTree []*KnowledgeItemInfo
//}

type KnowledgeItemInfo struct {
	ID       uint64 `json:"knowledge_item_id"`
	Title    string `json:"title"`
	ParentID uint64 `json:"parent_id"`
	Level    uint64 `json:"level"`
	Index    uint64 `json:"index"`
	Children []*KnowledgeItemInfo
}

// CreateKnowledgeItem create knowledge item
func CreateKnowledgeItem(r *KnowledgeItemCreateRequest, userID uint64) (*KnowledgeItemCreateResponse, error) {
	// knowledge item
	kItem := &model.KnowledgeItem{
		KnowledgeID:    r.KnowledgeID,
		UserID:         userID,
		Title:          strings.TrimSpace(r.Title),
		ContentVersion: 0,
		GUID:           "",
		CommentCount:   0,
		ViewCount:      0,
	}

	if r.ParentID == 0 {
		kItem.ParentID = 0
		kItem.Level = 0
		kItem.Index = 0
	}

	// knowledge item content
	content := "# " + strings.TrimSpace(r.Title)
	contentVersion := snowflake.GenerateSnowflakeID()
	var contentStatus uint8 = 0
	kItem.ItemContents = []model.KnowledgeItemContent{
		{
			Version: contentVersion,
			Status:  contentStatus,
			Content: content,
		},
	}

	k, err := dao.Engine.CreateKnowledgeItem(kItem)
	if err != nil {
		return nil, errno.New(errno.ErrDatabase, err)
	}

	rsp := &KnowledgeItemCreateResponse{
		ID:            k.ID,
		Title:         k.Title,
		ParentID:      k.ParentID,
		Content:       content,
		Version:       contentVersion,
		VersionStatus: contentStatus,
	}

	return rsp, nil
}

// GetKnowledgeItemList get knowledge item list as a tree
func GetKnowledgeItemList(kID int) ([]*KnowledgeItemInfo, error) {
	kItems, err := dao.Engine.GetKnowledgeItemListByKnowledgeID(uint64(kID))
	if err != nil {
		return nil, errno.New(errno.ErrDatabase, err)
	}

	tree := getKnowledgeItemTree(kItems, 0)
	return tree, nil
}

func getKnowledgeItemTree(kItems []*model.KnowledgeItem, pID uint64) []*KnowledgeItemInfo {
	var tree []*KnowledgeItemInfo

	for _, kItem := range kItems {
		if pID == kItem.ParentID {
			treeNode := &KnowledgeItemInfo{
				ID:       kItem.ID,
				Title:    kItem.Title,
				ParentID: kItem.ParentID,
				Level:    kItem.Level,
				Index:    kItem.Index,
			}
			treeNode.Children = getKnowledgeItemTree(kItems, kItem.ID)
			tree = append(tree, treeNode) // TODO 是否根据index排序
		}
	}
	return tree
}
