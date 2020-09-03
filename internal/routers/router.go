package routers

import (
	"html/template"
	"net/http"
	"path/filepath"

	apiHandler "github.com/puti-projects/puti/internal/backend/handler"
	"github.com/puti-projects/puti/internal/backend/handler/article"
	"github.com/puti-projects/puti/internal/backend/handler/auth"
	"github.com/puti-projects/puti/internal/backend/handler/media"
	"github.com/puti-projects/puti/internal/backend/handler/option"
	"github.com/puti-projects/puti/internal/backend/handler/page"
	"github.com/puti-projects/puti/internal/backend/handler/statistics"
	"github.com/puti-projects/puti/internal/backend/handler/subject"
	"github.com/puti-projects/puti/internal/backend/handler/taxonomy"
	"github.com/puti-projects/puti/internal/backend/handler/user"
	apiMiddleware "github.com/puti-projects/puti/internal/backend/middleware"
	webHandler "github.com/puti-projects/puti/internal/frontend/handler"
	webMiddleware "github.com/puti-projects/puti/internal/frontend/middleware"
	"github.com/puti-projects/puti/internal/pkg/logger"
	optionCache "github.com/puti-projects/puti/internal/pkg/option"
	"github.com/puti-projects/puti/internal/pkg/theme"
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/puti-projects/puti/internal/pkg/config"
	"github.com/puti-projects/puti/internal/utils"
)

func NewRouter() *gin.Engine {
	// create the gin engine
	g := gin.New()

	g = setFuncMap(g)

	// r.Use(gin.Logger())   // default logger middleware
	g.Use(gin.Recovery()) // default recovery middleware

	currentTheme := optionCache.Options.Get("current_theme")

	if viper.GetString("runmode") == gin.DebugMode {
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
		"formatNullTime": func(time *mysql.NullTime, format string) string {
			return utils.GetFormatNullTime(time, format)
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
	web := g.Group("")
	web.Use(webMiddleware.Renderer)
	{
		web.GET("", webHandler.ShowIndex)
		web.GET("/article", webHandler.ShowArticleList)
		web.GET("/category/:slug", webHandler.ShowCategoryArticleList)
		web.GET("/tag/:slug", webHandler.ShowTagArticleList)
		web.GET("/article/:id", webHandler.ShowArticleDetail)
		web.GET("/archive", webHandler.ShowArchive)
		web.GET("/subject", webHandler.ShowTopSubjects)
		web.GET("/subject/:slug", webHandler.ShowSubjects)

	}

	// no route handle
	g.NoRoute(webMiddleware.Renderer, webHandler.ShowNotFound)
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
	api := g.Group("/api")

	api.Use(apiMiddleware.NoCache)
	api.Use(apiMiddleware.Secure)
	api.Use(apiMiddleware.RequestID())

	api.POST("/login", auth.Login)
	api.GET("/token", auth.Info)

	api.Use(apiMiddleware.AuthMiddleware())
	{
		api.GET("/statistics/dashboard", statistics.Dashboard)
		api.GET("/statistics/system", statistics.System)
		api.POST("/user/:username", user.Create)
		api.GET("/user/:username", user.Get)
		api.DELETE("/user/:id", user.Delete)
		api.PUT("/user/:id", user.Update)
		api.GET("/user", user.List)
		api.POST("/avatar", user.Avatar)
		api.GET("/article", article.List)
		api.GET("/article/:id", article.Get)
		api.POST("/article", article.Create)
		api.PUT("/article/:id", article.Update)
		api.DELETE("/article/:id", article.Delete)
		api.GET("/page", page.List)
		api.GET("/page/:id", page.Get)
		api.POST("/page", page.Create)
		api.PUT("/page/:id", page.Update)
		api.DELETE("/page/:id", page.Delete)
		api.POST("/taxonomy/:name", taxonomy.Create)
		api.GET("/taxonomy/:id", taxonomy.Get)
		api.DELETE("/taxonomy/:id", taxonomy.Delete)
		api.PUT("/taxonomy/:id", taxonomy.Update)
		api.GET("/taxonomy", taxonomy.List)
		api.GET("/media/:id", media.Detail)
		api.GET("/media", media.List)
		api.POST("/media", media.Upload)
		api.DELETE("/media/:id", media.Delete)
		api.PUT("/media/:id", media.Update)
		api.GET("/option", option.List)
		api.PUT("/option", option.Update)
		api.GET("/subject", subject.List)
		api.GET("/subject/:id", subject.Detail)
		api.POST("/subject/:name", subject.Create)
		api.PUT("/subject/:id", subject.Update)
		api.DELETE("/subject/:id", subject.Delete)
	}
}

// loadHelthTest the health check handlers
func loadHealthTest(g *gin.Engine) {
	// Group for health check
	svcd := g.Group("/check")
	{
		svcd.GET("/health", apiHandler.HealthCheck)
	}
}
