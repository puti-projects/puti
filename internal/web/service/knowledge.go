package service

import (
	"errors"
	"html/template"
	"strconv"
	"sync"

	"github.com/puti-projects/puti/internal/model"
	"github.com/puti-projects/puti/internal/pkg/markdown"
	"github.com/puti-projects/puti/internal/utils"
	"github.com/puti-projects/puti/internal/web/dao"

	"gorm.io/gorm"
)

// KnowledgeList knowledge list
type KnowledgeList struct {
	Lock    *sync.Mutex
	DocMap  map[uint64]*ShowKnowledgeList
	NoteMap map[uint64]*ShowKnowledgeList
}

// GetKnowledgeList get knowledge list, include note and doc
func GetKnowledgeList() (map[string]map[uint64]*ShowKnowledgeList, error) {
	res, err := dao.GetKnowledgeList()
	if err != nil {
		return nil, err
	}

	wg := sync.WaitGroup{}
	finished := make(chan bool, 1)

	knowledgeList := KnowledgeList{
		Lock:    new(sync.Mutex),
		DocMap:  make(map[uint64]*ShowKnowledgeList, 0),
		NoteMap: make(map[uint64]*ShowKnowledgeList, 0),
	}

	for _, k := range res {
		wg.Add(1)
		go func(k *dao.KnowledgeResult) {
			defer wg.Done()

			knowledgeList.Lock.Lock()
			defer knowledgeList.Lock.Unlock()

			var mapBody map[uint64]*ShowKnowledgeList
			if k.Type == model.KnowledgeTypeDoc {
				mapBody = knowledgeList.DocMap
			} else if k.Type == model.KnowledgeTypeNote {
				mapBody = knowledgeList.NoteMap
			}
			if mapBody != nil {
				mapBody[k.ID] = &ShowKnowledgeList{
					ID:            k.ID,
					URL:           "/knowledge/" + k.Type + "/" + k.Slug,
					Name:          k.Name,
					Slug:          k.Slug,
					Description:   k.Description,
					CoverImageURL: k.CoverImageURL,
					UpdatedTime:   utils.GetFormatTime(&k.UpdatedTime, "2006-01-02 15:04"),
				}
			}
		}(k)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	<-finished

	list := make(map[string]map[uint64]*ShowKnowledgeList, 0)
	list["Doc"] = knowledgeList.DocMap
	list["Note"] = knowledgeList.NoteMap
	return list, nil
}

// GetKnowledgeBySlug get knowledge by slug
func GetKnowledgeBySlug(kType, kSlug string) (*ShowKnowledgeInfo, error) {
	k, err := dao.GetKnowledgeBySlug(kType, kSlug)
	if err != nil {
		return nil, err
	}

	info := &ShowKnowledgeInfo{
		ID:          k.ID,
		Name:        k.Name,
		LastUpdated: utils.GetFormatNullTime(&k.LastUpdated, "2006-01-02 15:04:05"),
	}
	return info, nil
}

// GetKnowledgeItemList get knowledge item list in tree
func GetKnowledgeItemList(kType, kSlug string, kID uint64) ([]*ShowKnowledgeItemTreeNode, error) {
	res, err := dao.GetKnowledgeItemList(kID)
	if err != nil {
		return nil, err
	}

	urlPrefix := "/knowledge/" + kType + "/" + kSlug + "/"
	tree := generateKnowledgeItemTree(res, 0, urlPrefix)
	return tree, nil
}

func generateKnowledgeItemTree(ki []*dao.KnowledgeItemResult, pid uint64, urlPrefix string) []*ShowKnowledgeItemTreeNode {
	var tree []*ShowKnowledgeItemTreeNode

	for _, v := range ki {
		if v.ParentID == pid {
			treeNode := ShowKnowledgeItemTreeNode{
				ID:     v.ID,
				Symbol: v.Symbol,
				Title:  v.Title,
				URL:    urlPrefix + strconv.Itoa(int(v.Symbol)),
				Level:  v.Level,
				Index:  v.Index,
			}
			treeNode.Children = generateKnowledgeItemTree(ki, v.ID, urlPrefix)
			tree = append(tree, &treeNode)
		}
	}
	return tree
}

// GetKnowledgeItemContentBySymbol get knowledge item content by knowledge item symbol
func GetKnowledgeItemContentBySymbol(kiSymbol string) (*ShowKnowledgeItemContent, error) {
	result, err := dao.GetKnowledgeItemContentBySymbol(kiSymbol)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if result != nil {
		output := &ShowKnowledgeItemContent{
			Symbol:  result.Symbol,
			Title:   template.HTMLEscapeString(result.Title),
			Content: template.HTML(markdown.Markdown2HTML(result.Title, result.Content)),
		}
		return output, nil
	}

	return nil, nil
}
