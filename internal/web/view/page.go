package view

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/puti-projects/puti/internal/pkg/config"
	"github.com/puti-projects/puti/internal/pkg/counter"
	"github.com/puti-projects/puti/internal/web/service"

	"github.com/gin-gonic/gin"
)

// CheckAndShowPage check whether page exist
func CheckAndShowPage(c *gin.Context) {
	if !strings.HasPrefix(c.Request.RequestURI, "/themes") {
		// no static, and get slug
		slug := strings.TrimLeft(c.Request.RequestURI, "/")
		if !strings.Contains(slug, "/") {
			if pageID := service.GetPageIDBySlug(slug); pageID > 0 {
				ShowPageDetail(c, pageID)
				return
			}
		}
	}

	ShowNotFound(c)
}

// ShowPageDetail handle page info
// TODO get page template setting
func ShowPageDetail(c *gin.Context, pageID uint64) {
	renderData := getRenderData(c)

	var pageDetail *service.ShowPageDetail
	var err error
	// cache handle
	if data, exist := service.SrvEngine.GetCache(config.CachePageDetailPrefix + strconv.Itoa(int(pageID))); exist {
		service.SrvEngine.JSONUnmarshal(data, &pageDetail)
	} else {
		pageDetail, err = service.GetPageDetailByID(pageID)
		if err != nil {
			ShowInternalServerError(c)
			return
		}
		service.SrvEngine.MarshalAndSetCache(config.CachePageDetailPrefix+strconv.Itoa(int(pageID)), pageDetail)
	}
	renderData["Page"] = pageDetail

	counter.CounterCache.CountOne(c.ClientIP(), pageID)

	renderData["Widgets"] = getWidgets()
	renderData["Title"] = pageDetail.Title + " - " + renderData["Setting"].(map[string]interface{})["BlogName"].(string)
	c.HTML(http.StatusOK, getTheme(c)+"/page-detail.html", renderData)
}
