package web

import (
	"html/template"

	optionCache "github.com/puti-projects/puti/internal/pkg/option"

	"github.com/gin-gonic/gin"
)

// RenderData renderer data
type RenderData map[string]interface{}

// Renderer  gin handlerfunc for set rederer data
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
	settingMap["BlogName"] = optionCache.Options.Get("blog_name")
	settingMap["BlogDescription"] = optionCache.Options.Get("blog_description")
	settingMap["SiteUrl"] = optionCache.Options.Get("site_url")
	settingMap["SiteDescription"] = optionCache.Options.Get("site_description")
	settingMap["SiteKeywords"] = optionCache.Options.Get("site_keywords")
	settingMap["FooterCopyright"] = template.HTML(optionCache.Options.Get("footer_copyright"))
	settingMap["SiteLanguage"] = optionCache.Options.Get("site_language")
	settingMap["CurrentTheme"] = optionCache.Options.Get("current_theme")

	settingMap["ShowOnFront"] = optionCache.Options.Get("show_on_front")

	settingMap["ArticleCommentStatus"] = optionCache.Options.Get("article_comment_status")
	settingMap["pageCommentStatus"] = optionCache.Options.Get("page_comment_status")

	(*renderData)["Setting"] = settingMap

	return c
}
