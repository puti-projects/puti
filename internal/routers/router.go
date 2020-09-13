package routers

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/puti-projects/puti/internal/backend/api"
	"github.com/puti-projects/puti/internal/backend/api/article"
	"github.com/puti-projects/puti/internal/backend/api/auth"
	"github.com/puti-projects/puti/internal/backend/api/media"
	"github.com/puti-projects/puti/internal/backend/api/option"
	"github.com/puti-projects/puti/internal/backend/api/page"
	"github.com/puti-projects/puti/internal/backend/api/statistics"
	"github.com/puti-projects/puti/internal/backend/api/subject"
	"github.com/puti-projects/puti/internal/backend/api/taxonomy"
	"github.com/puti-projects/puti/internal/backend/api/user"
	"github.com/puti-projects/puti/internal/frontend/web"
	"github.com/puti-projects/puti/internal/pkg/config"
	"github.com/puti-projects/puti/internal/pkg/logger"
	optionCache "github.com/puti-projects/puti/internal/pkg/option"
	"github.com/puti-projects/puti/internal/pkg/theme"
	"github.com/puti-projects/puti/internal/routers/middleware"
	apiMiddleware "github.com/puti-projects/puti/internal/routers/middleware/api"
	webMiddleware "github.com/puti-projects/puti/internal/routers/middleware/web"
	"github.com/puti-projects/puti/internal/utils"

	"github.com/gin-gonic/gin"
)

func NewRouter(runmode string) *gin.Engine {
	// Set gin mode before initialize the gin router
	if "debug" == runmode {
		gin.SetMode(gin.DebugMode)
	} else if "test" == runmode {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// create the gin engine
	g := gin.New()

	g = setFuncMap(g)

	g.Use(middleware.AccessLogger())
	g.Use(middleware.Recovery())

	currentTheme := optionCache.Options.Get("current_theme")

	if runmode == gin.DebugMode {
		g.Use(apiMiddleware.Options)
	}

	loadHealthTest(g)
	loadAPI(g)
	loadWeb(g, currentTheme)
	loadStatic(g, currentTheme)

	return g
}

func setFuncMap(g *gin.Engine) *gin.Engine {
	g.SetFuncMap(template.FuncMap{
		"minus": func(a, b int) int {
			return a - b
		},
		// TODO this function should remove after output of widget standardization
		"formatNullTime": func(time sql.NullTime, format string) string {
			return utils.GetFormatNullTime(&time, format)
		},
	})

	return g
}

// loadWeb load frontend and backend entrance(SPA web)
func loadWeb(g *gin.Engine, currentTheme string) {
	// Group for backend
	admin := g.Group("/admin")
	admin.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "console.html", gin.H{})
	})

	// Group for frontend
	// notice: page route is handle in NoRoute(), since the wildcard problem in root from httprouter
	webGroup := g.Group("")
	webGroup.Use(webMiddleware.Renderer)
	{
		webGroup.GET("", web.ShowIndex)
		webGroup.GET("/article", web.ShowArticleList)
		webGroup.GET("/category/:slug", web.ShowCategoryArticleList)
		webGroup.GET("/tag/:slug", web.ShowTagArticleList)
		webGroup.GET("/article/:id", web.ShowArticleDetail)
		webGroup.GET("/archive", web.ShowArchive)
		webGroup.GET("/subject", web.ShowTopSubjects)
		webGroup.GET("/subject/:slug", web.ShowSubjects)

	}

	// no route handle
	g.NoRoute(webMiddleware.Renderer, web.ShowNotFound)
}

// loadStatic load static resource
func loadStatic(g *gin.Engine, currentTheme string) {
	// resource
	g.Static("/static", config.StaticPath("console/static"))
	g.Static("/uploads", config.StaticPath("uploads/"))
	g.Static("/assets", config.StaticPath("assets/"))
	g.StaticFile("/favicon.ico", config.StaticPath("assets/favicon.ico"))

	// load theme templates file
	var themeTemplates []string
	for _, theme := range theme.Themes {
		themePath := config.StaticPath("theme/" + theme)
		g.Static("/theme/"+theme+"/public", themePath+"/public")
		g.StaticFile("/theme/"+theme+"/thumbnail.jpg", themePath+"/thumbnail.jpg")
		themeTemplate, err := filepath.Glob(themePath + "/*.html")
		if nil != err {
			logger.Fatalf("load theme %s templates failed: %s", theme, err.Error())
		}
		themeTemplates = append(themeTemplates, themeTemplate...)
	}
	commentTemplates, err := filepath.Glob(config.StaticPath("theme/common/comment/*.html"))
	if nil != err {
		logger.Fatal("load comment templates failed: " + err.Error())
	}
	headTemplates, err := filepath.Glob(config.StaticPath("theme/common/head/*.html"))
	if nil != err {
		logger.Fatal("load head templates failed: " + err.Error())
	}
	templates := append(themeTemplates, commentTemplates...)
	templates = append(templates, headTemplates...)
	// load backend console html
	templates = append(templates, config.StaticPath("console/console.html"))
	// load all files
	g.LoadHTMLFiles(templates...)
}

// loadAPI load api part
func loadAPI(g *gin.Engine) {
	// Group for api
	apiGroup := g.Group("/api")

	apiGroup.Use(apiMiddleware.NoCache)
	apiGroup.Use(apiMiddleware.Secure)
	apiGroup.Use(apiMiddleware.RequestID())

	apiGroup.POST("/login", auth.Login)
	apiGroup.GET("/token", auth.Info)

	apiGroup.Use(apiMiddleware.AuthMiddleware())
	{
		apiGroup.GET("/statistics/dashboard", statistics.Dashboard)
		apiGroup.GET("/statistics/system", statistics.System)
		apiGroup.POST("/user/:username", user.Create)
		apiGroup.GET("/user/:username", user.Get)
		apiGroup.DELETE("/user/:id", user.Delete)
		apiGroup.PUT("/user/:id", user.Update)
		apiGroup.GET("/user", user.List)
		apiGroup.POST("/avatar", user.Avatar)
		apiGroup.GET("/article", article.List)
		apiGroup.GET("/article/:id", article.Get)
		apiGroup.POST("/article", article.Create)
		apiGroup.PUT("/article/:id", article.Update)
		apiGroup.DELETE("/article/:id", article.Delete)
		apiGroup.GET("/page", page.List)
		apiGroup.GET("/page/:id", page.Get)
		apiGroup.POST("/page", page.Create)
		apiGroup.PUT("/page/:id", page.Update)
		apiGroup.DELETE("/page/:id", page.Delete)
		apiGroup.POST("/taxonomy/:name", taxonomy.Create)
		apiGroup.GET("/taxonomy/:id", taxonomy.Get)
		apiGroup.DELETE("/taxonomy/:id", taxonomy.Delete)
		apiGroup.PUT("/taxonomy/:id", taxonomy.Update)
		apiGroup.GET("/taxonomy", taxonomy.List)
		apiGroup.GET("/media/:id", media.Detail)
		apiGroup.GET("/media", media.List)
		apiGroup.POST("/media", media.Upload)
		apiGroup.DELETE("/media/:id", media.Delete)
		apiGroup.PUT("/media/:id", media.Update)
		apiGroup.GET("/option", option.List)
		apiGroup.PUT("/option", option.Update)
		apiGroup.GET("/subject", subject.List)
		apiGroup.GET("/subject/:id", subject.Detail)
		apiGroup.POST("/subject/:name", subject.Create)
		apiGroup.PUT("/subject/:id", subject.Update)
		apiGroup.DELETE("/subject/:id", subject.Delete)
	}
}

// loadHelthTest the health check handlers
func loadHealthTest(g *gin.Engine) {
	// Group for health check
	svcd := g.Group("/check")
	{
		svcd.GET("/health", api.HealthCheck)
	}
}
