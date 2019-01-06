package router

import (
	"net/http"

	"github.com/puti-projects/puti/internal/backend/handler/article"
	"github.com/puti-projects/puti/internal/backend/handler/auth"
	"github.com/puti-projects/puti/internal/backend/handler/media"
	"github.com/puti-projects/puti/internal/backend/handler/option"
	"github.com/puti-projects/puti/internal/backend/handler/page"
	"github.com/puti-projects/puti/internal/backend/handler/sd"
	"github.com/puti-projects/puti/internal/backend/handler/taxonomy"
	"github.com/puti-projects/puti/internal/backend/handler/user"
	"github.com/puti-projects/puti/internal/common/router/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Load loads the middlewares, routes, handles.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	if viper.GetString("runmode") == gin.DebugMode {
		g.Use(middleware.Options)
	}

	loadHealthTest(g)
	loadAPI(g)
	loadWeb(g)
	loadStatic(g)

	return g
}

// loadWeb load frontend and backend entrance(SPA web)
func loadWeb(g *gin.Engine) *gin.Engine {
	// backend index
	g.LoadHTMLFiles("web/backend/index.html")
	g.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	// 404 handle
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect route.")
	})

	return g
}

// loadStatic load static resource
func loadStatic(g *gin.Engine) *gin.Engine {
	g.Static("/static", "web/backend/static")
	g.Static("/uploads", "uploads/")
	// g.StaticFile("/favicon.ico", "./resources/favicon.ico")

	return g
}

// loadAPI load api part
func loadAPI(g *gin.Engine) *gin.Engine {
	// Group for api
	api := g.Group("/api")

	api.Use(middleware.NoCache)
	api.Use(middleware.Secure)
	api.Use(middleware.RequestID())

	api.POST("/login", auth.Login)
	api.GET("/token", auth.Info)

	api.Use(middleware.AuthMiddleware())
	{
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
	}

	return g
}

// loadHelthTest the health check handlers
func loadHealthTest(g *gin.Engine) *gin.Engine {
	// Group for health check
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
