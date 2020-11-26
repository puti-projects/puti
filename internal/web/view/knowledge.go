package view

import (
	"github.com/puti-projects/puti/internal/model"
	"github.com/puti-projects/puti/internal/pkg/config"
	"github.com/puti-projects/puti/internal/web/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowKnowledgeDetail(c *gin.Context) {
	renderData := getRenderData(c)

	// get params
	kType := c.Param("type")
	if kType != model.KnowledgeTypeDoc && kType != model.KnowledgeTypeNote {
		ShowNotFound(c)
		return
	}
	kSlug := c.Param("slug")
	kiSymbol := c.Param("symbol")

	var err error
	var kInfo *service.ShowKnowledgeInfo
	if data, exist := service.SrvEngine.GetCache(config.CacheKnowledgeInfoPrefix + kSlug); exist {
		service.SrvEngine.JSONUnmarshal(data, &kInfo)
	} else {
		kInfo, err = service.GetKnowledgeBySlug(kType, kSlug)
		if err != nil {
			ShowInternalServerError(c)
			return
		}
		service.SrvEngine.MarshalAndSetCache(config.CacheKnowledgeInfoPrefix+kSlug, kInfo)
	}
	renderData["KInfo"] = kInfo

	var treeList []*service.ShowKnowledgeItemTreeNode
	if data, exist := service.SrvEngine.GetCache(config.CacheKnowledgeItemListPrefix + strconv.Itoa(int(kInfo.ID))); exist {
		service.SrvEngine.JSONUnmarshal(data, &treeList)
	} else {
		treeList, err = service.GetKnowledgeItemList(kType, kSlug, kInfo.ID)
		if err != nil {
			ShowInternalServerError(c)
			return
		}
		service.SrvEngine.MarshalAndSetCache(config.CacheKnowledgeItemListPrefix+strconv.Itoa(int(kInfo.ID)), treeList)
	}
	renderData["KiList"] = treeList

	// check symbol
	if kiSymbol != "" {
		var content *service.ShowKnowledgeItemContent
		if data, exist := service.SrvEngine.GetCache(config.CacheKnowledgeItemContentPrefix + kiSymbol); exist {
			service.SrvEngine.JSONUnmarshal(data, &content)
		} else {
			content, err = service.GetKnowledgeItemContentBySymbol(kiSymbol)
			if err != nil {
				ShowInternalServerError(c)
				return
			}
			service.SrvEngine.MarshalAndSetCache(config.CacheKnowledgeItemContentPrefix+kiSymbol, content)
		}
		renderData["KiContent"] = content
	}

	renderData["Title"] = kInfo.Name + " - " + renderData["Setting"].(map[string]interface{})["BlogName"].(string)
	c.HTML(http.StatusOK, getTheme(c)+"/knowledge.html", renderData)
}
