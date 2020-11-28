package web

import (
	"github.com/puti-projects/puti/internal/pkg/theme"
	"html/template"

	"github.com/puti-projects/puti/internal/pkg/cache"

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
	renderThemeSetting(c, renderData)
	c.Next()
}

func renderBasicConfig(c *gin.Context, renderData *RenderData) *gin.Context {
	configMap := map[string]interface{}{
		"StaticServer": "",
	}

	(*renderData)["Config"] = configMap
	return c
}

func renderBasicSetting(c *gin.Context, renderData *RenderData) *gin.Context {
	// load some basic settings
	settingMap := map[string]interface{}{
		"CurrentUrl":      c.Request.URL.Path,
		"BlogName":        cache.Options.Get("blog_name"),
		"BlogDescription": cache.Options.Get("blog_description"),
		"SiteUrl":         cache.Options.Get("site_url"),
		"SiteDescription": cache.Options.Get("site_description"),
		"SiteKeywords":    cache.Options.Get("site_keywords"),
		"FooterCopyright": template.HTML(cache.Options.Get("footer_copyright")),
		"SiteLanguage":    cache.Options.Get("site_language"),
		"CurrentTheme":    cache.Options.Get("current_theme"),

		"ShowOnFront":          cache.Options.Get("show_on_front"),
		"ArticleCommentStatus": cache.Options.Get("article_comment_status"),
		"pageCommentStatus":    cache.Options.Get("page_comment_status"),
	}

	(*renderData)["Setting"] = settingMap
	return c
}

func renderThemeSetting(c *gin.Context, renderData *RenderData) *gin.Context {
	themeMap := map[string]interface{}{}

	if t, ok := theme.Themes[cache.Options.Get("current_theme")]; ok && t.FaviconExist {
		themeMap["Favicon"] = "/theme/" + cache.Options.Get("current_theme") + "/favicon.ico"
	} else {
		themeMap["Favicon"] = "/favicon.ico"
	}

	(*renderData)["Theme"] = themeMap
	return c
}
