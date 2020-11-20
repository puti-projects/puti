package web

import (
	"github.com/puti-projects/puti/internal/pkg/cache"
	"html/template"

	"github.com/gin-gonic/gin"
)

// RenderData renderer data
type RenderData map[string]interface{}

// Renderer  gin handlerFunc for set render data
func Renderer(c *gin.Context) {
	// init the data
	renderData := &RenderData{}
	c.Set("renderData", renderData)

	// session := util.GetSession(c)
	// (*dataModel)["User"] = session

	renderBasicSetting(c, renderData)

	c.Next()
}

func renderBasicConfig(c *gin.Context, renderData *RenderData) *gin.Context {
	configMap := map[string]interface{}{}
	configMap["StaticServer"] = ""

	(*renderData)["Config"] = configMap
	return c
}

func renderBasicSetting(c *gin.Context, renderData *RenderData) *gin.Context {
	settingMap := map[string]interface{}{}

	// load some basic settings
	settingMap["CurrentUrl"] = c.Request.URL.Path
	settingMap["BlogName"] = cache.Options.Get("blog_name")
	settingMap["BlogDescription"] = cache.Options.Get("blog_description")
	settingMap["SiteUrl"] = cache.Options.Get("site_url")
	settingMap["SiteDescription"] = cache.Options.Get("site_description")
	settingMap["SiteKeywords"] = cache.Options.Get("site_keywords")
	settingMap["FooterCopyright"] = template.HTML(cache.Options.Get("footer_copyright"))
	settingMap["SiteLanguage"] = cache.Options.Get("site_language")
	settingMap["CurrentTheme"] = cache.Options.Get("current_theme")

	settingMap["ShowOnFront"] = cache.Options.Get("show_on_front")

	settingMap["ArticleCommentStatus"] = cache.Options.Get("article_comment_status")
	settingMap["pageCommentStatus"] = cache.Options.Get("page_comment_status")

	(*renderData)["Setting"] = settingMap

	return c
}
