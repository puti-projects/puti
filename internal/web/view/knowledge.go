package view

import (
	"net/http"
	"strconv"

	"github.com/puti-projects/puti/internal/model"
	"github.com/puti-projects/puti/internal/pkg/config"
	"github.com/puti-projects/puti/internal/web/service"

	"github.com/gin-gonic/gin"
)

// ShowKnowledgeDetail show knowledge detail
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
	content := c.Query("content")

	var err error
	if content == "1" {
		if kiSymbol != "" {
			var content *service.ShowKnowledgeItemContent
			if data, exist := service.SrvEngine.GetCache(config.CacheKnowledgeItemContentPrefix + kiSymbol); exist {
				service.SrvEngine.JSONUnmarshal(data, &content)
			} else {
				content, err = service.SrvEngine.GetKnowledgeItemContentBySymbol(kiSymbol)
				if err != nil {
					ShowInternalServerError(c)
					return
				}
				service.SrvEngine.MarshalAndSetCache(config.CacheKnowledgeItemContentPrefix+kiSymbol, content)
			}
			renderData["KiContent"] = content
		}

		c.HTML(http.StatusOK, getTheme(c)+"/knowledge-content.html", renderData)
	} else {
		var kInfo *service.ShowKnowledgeInfo
		if data, exist := service.SrvEngine.GetCache(config.CacheKnowledgeInfoPrefix + kSlug); exist {
			service.SrvEngine.JSONUnmarshal(data, &kInfo)
		} else {
			kInfo, err = service.SrvEngine.GetKnowledgeBySlug(kType, kSlug)
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
			treeList, err = service.SrvEngine.GetKnowledgeItemList(kType, kSlug, kInfo.ID)
			if err != nil {
				ShowInternalServerError(c)
				return
			}
			service.SrvEngine.MarshalAndSetCache(config.CacheKnowledgeItemListPrefix+strconv.Itoa(int(kInfo.ID)), treeList)
		}
		renderData["KiList"] = treeList

		// check symbol
		if kiSymbol == "" && len(treeList) > 0 {
			// give a default symbol; first page
			kiSymbol = strconv.Itoa(int(treeList[0].Symbol))
		}

		// kiSymbol still can be empty
		if kiSymbol != "" {
			var content *service.ShowKnowledgeItemContent
			if data, exist := service.SrvEngine.GetCache(config.CacheKnowledgeItemContentPrefix + kiSymbol); exist {
				service.SrvEngine.JSONUnmarshal(data, &content)
			} else {
				content, err = service.SrvEngine.GetKnowledgeItemContentBySymbol(kiSymbol)
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
}
