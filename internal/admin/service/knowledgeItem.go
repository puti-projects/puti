package service

import (
	"strconv"
	"strings"

	"github.com/puti-projects/puti/internal/admin/dao"
	"github.com/puti-projects/puti/internal/model"
	"github.com/puti-projects/puti/internal/pkg/errno"
	"github.com/puti-projects/puti/internal/pkg/snowflake"
	"github.com/puti-projects/puti/internal/utils"
)

const (
	// KnowledgeItemUpdateTypeSave update type of 'save', only save current version content
	KnowledgeItemUpdateTypeSave = "save"
	// KnowledgeItemUpdateTypePublish update type of 'publish', publish will set now version as the published version
	KnowledgeItemUpdateTypePublish = "publish"
)

// KnowledgeItemCreateRequest struct for binding knowledge item create
type KnowledgeItemCreateRequest struct {
	CreateType  string `json:"create_type" binding:"required,oneof=doc note"`
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
	Index         int64  `json:"index"`
	Content       string `json:"content"`
	Version       uint64 `json:"version"`
	VersionStatus uint8  `json:"version_status"`
}

// KnowledgeItemUpdateInfoRequest struct bind to update knowledge item info
type KnowledgeItemUpdateInfoRequest struct {
	ID               uint64 `json:"knowledge_item_id" binding:"required"`
	Title            string `json:"title" binding:"required"`
	NodeChange       bool   `json:"node_change"`
	ParentID         uint64 `json:"parent_id"`
	IndexChange      string `json:"index_change"`
	IndexRelatedNode uint64 `json:"index_related_node"`
}

// KnowledgeItemUpdateContentRequest struct bind to update knowledge item content
type KnowledgeItemUpdateContentRequest struct {
	ID       uint64 `json:"knowledge_item_id" binding:"required"`
	EditType string `json:"edit_type" binding:"required,oneof=doc note"`     // doc or note
	SaveType string `json:"save_type" binding:"required,oneof=save publish"` // save or publish
	Content  string `json:"content" binding:"required"`
	Version  string `json:"version" binding:"required"`
}

// KnowledgeItemUpdateContentResponse update knowledge item content response
type KnowledgeItemUpdateContentResponse struct {
	ID      uint64 `json:"knowledge_item_id"`
	Version uint64 `json:"version,string"`
}

// KnowledgeItemInfo knowledge info for list tree
type KnowledgeItemInfo struct {
	ID       uint64               `json:"knowledge_item_id"`
	Title    string               `json:"title"`
	ParentID uint64               `json:"parent_id"`
	Level    uint64               `json:"level"`
	Index    int64                `json:"index"`
	Children []*KnowledgeItemInfo `json:"children"`
}

// KnowledgeItemDetail knowledge detail info include content
type KnowledgeItemDetail struct {
	ID                      uint64 `json:"knowledge_item_id"`
	Title                   string `json:"title"`
	Content                 string `json:"content"`
	ContentStatus           uint8  `json:"content_status"`
	ContentPublishedVersion uint64 `json:"content_published_version"`
	ContentNowVersion       uint64 `json:"content_now_version,string"`
	ContentUpdatedTime      string `json:"content_updated_time"`
}

// CreateKnowledgeItem create knowledge item
func CreateKnowledgeItem(r *KnowledgeItemCreateRequest, userID uint64) (*KnowledgeItemCreateResponse, error) {
	// knowledge item
	kItem := &model.KnowledgeItem{
		KnowledgeID:    r.KnowledgeID,
		Symbol:         snowflake.GenerateSnowflakeID(),
		UserID:         userID,
		Title:          strings.TrimSpace(r.Title),
		ContentVersion: 0,
		CommentCount:   0,
		ViewCount:      0,
	}

	// TODO now directly set index to 0 if create a new item. should set it to the first one always
	kItem.ParentID = r.ParentID
	if r.ParentID == 0 {
		kItem.Level = 1
	} else {
		// check parent's level
		parent, err := dao.Engine.GetKnowledgeItemByID(r.ParentID)
		if err != nil {
			return nil, errno.New(errno.ErrDatabase, err)
		}
		kItem.Level = parent.Level + 1
	}
	kItem.Index = 0

	// init a knowledge item content record
	content := "# " + strings.TrimSpace(r.Title)
	contentVersion := snowflake.GenerateSnowflakeID()
	var contentStatus uint8
	switch r.CreateType {
	case model.KnowledgeTypeDoc:
		contentStatus = 0
	case model.KnowledgeTypeNote:
		contentStatus = 1
	}
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

// GetKnowledgeItemInfo get knowledge item info
func GetKnowledgeItemInfo(kiID int) (*KnowledgeItemDetail, error) {
	ki, err := dao.Engine.GetKnowledgeItemByID(uint64(kiID))
	if err != nil {
		return nil, err
	}

	if ki != nil {
		kiDetail := &KnowledgeItemDetail{
			ID:                      ki.ID,
			Title:                   ki.Title,
			ContentPublishedVersion: ki.ContentVersion,
		}

		kic, err := dao.Engine.GetKnowledgeItemLastContent(ki)
		if err != nil {
			return nil, err
		}

		kiDetail.Content = kic.Content
		kiDetail.ContentStatus = kic.Status
		kiDetail.ContentNowVersion = kic.Version
		kiDetail.ContentUpdatedTime = utils.GetFormatTime(&kic.UpdatedAt, "2006-01-02 15:04:05")

		return kiDetail, nil
	}

	return nil, err
}

// UpdateKnowledgeItemInfo update knowledge item info
func UpdateKnowledgeItemInfo(ir *KnowledgeItemUpdateInfoRequest, kItemID uint64) error {
	// if change node
	if ir.NodeChange {
		if err := dao.Engine.UpdateKnowledgeItemWithNodeChange(kItemID, ir.ParentID, ir.IndexChange, ir.IndexRelatedNode); err != nil {
			return err
		}
	} else {
		//  do not change node, only update title
		if err := dao.Engine.UpdateKnowledgeItemTitle(kItemID, ir.Title); err != nil {
			return err
		}
	}

	// clean list cache
	kID, _ := dao.Engine.GetKnowledgeIDByItemID(kItemID)
	SrvEngine.CleanCacheKnowledgeItemList(kID)
	return nil
}

// UpdateKnowledgeItemContent update knowledge item content
func UpdateKnowledgeItemContent(cr *KnowledgeItemUpdateContentRequest, kItemID uint64) (*KnowledgeItemUpdateContentResponse, error) {
	var err error

	version, _ := strconv.Atoi(cr.Version)

	kItemContent, err := dao.Engine.GetKnowledgeItemContentByVersion(kItemID, uint64(version))
	if err != nil {
		return nil, err
	}

	rsp := &KnowledgeItemUpdateContentResponse{
		ID: kItemContent.KnowledgeItemID,
	}

	switch cr.EditType {
	case model.KnowledgeTypeNote:
		// only one version for note; directly update
		kItemContent.Content = cr.Content
		err = dao.Engine.UpdateKnowledgeItemContent(kItemContent)
		rsp.Version = kItemContent.Version
	case model.KnowledgeTypeDoc:
		if cr.SaveType == KnowledgeItemUpdateTypeSave {
			if kItemContent.Status == model.KnowledgeItemContentStatusCurrent {
				// create a new version
				newVersionContent := &model.KnowledgeItemContent{
					KnowledgeItemID: kItemID,
					Version:         snowflake.GenerateSnowflakeID(),
					Status:          model.KnowledgeItemContentStatusCommon,
					Content:         cr.Content,
				}
				err = dao.Engine.CreateKnowledgeItemContent(newVersionContent)
				rsp.Version = newVersionContent.Version
			} else if kItemContent.Status == model.KnowledgeItemContentStatusCommon {
				// directly update
				kItemContent.Content = cr.Content
				err = dao.Engine.UpdateKnowledgeItemContent(kItemContent)
				rsp.Version = kItemContent.Version
			}
		} else if cr.SaveType == KnowledgeItemUpdateTypePublish {
			if kItemContent.Status == model.KnowledgeItemContentStatusCurrent {
				// directly update
				kItemContent.Content = cr.Content
				err = dao.Engine.UpdateKnowledgeItemContent(kItemContent)
				rsp.Version = kItemContent.Version
			} else if kItemContent.Status == model.KnowledgeItemContentStatusCommon {
				// set 1 to this version; update old version to 0
				kItemContent.Status = model.KnowledgeItemContentStatusCurrent
				kItemContent.Content = cr.Content
				err = dao.Engine.ChangePublishedKnowledgeItemContent(kItemContent)
				rsp.Version = kItemContent.Version
			}
		}
	}

	if err != nil {
		return nil, errno.New(errno.ErrUpdateKnowledgeItemContent, err)
	}

	// update finished. clean cache.
	symbol, _ := dao.Engine.GetKnowledgeItemSymbolByID(kItemID)
	SrvEngine.CleanCacheAfterUpdateKnowledgeItemContent(symbol)
	return rsp, nil
}

// DeleteKnowledgeItem delete knowledge item
// Note: content will not be deleted
func DeleteKnowledgeItem(kItemID uint64) error {
	kItem, err := dao.Engine.GetKnowledgeItemByID(kItemID)
	if err != nil {
		return err
	}

	// check if can delete
	hasChild, err := dao.Engine.CheckKnowledgeItemHasChildren(kItemID, kItem.KnowledgeID)
	if err != nil {
		return err
	}

	if hasChild {
		return errno.ErrKnowledgeItemCanNotBeDeleted
	}

	if err := dao.Engine.DeleteKnowledgeItem(kItemID); err != nil {
		return errno.New(errno.ErrDatabase, err)
	}

	// clean list cache
	SrvEngine.CleanCacheKnowledgeItemList(kItem.KnowledgeID)
	return nil
}
