package router

import (
	"net/http"

	"gingob/handler/sd"
	"gingob/handler/user"
	"gingob/router/middleware"

	"github.com/gin-gonic/gin"
)

// Load loads the middlewares, routes, handles.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	// 404 handle
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// static resource
	g.Static("/static", "backend/dist/static")
	g.Static("/upload", "upload/")
	// g.StaticFile("/favicon.ico", "./resources/favicon.ico")

	// Group for api
	api := g.Group("/api")

	api.POST("/login", user.Login)
	api.GET("/token", user.Info)

	u := api.Group("/user")
	u.Use(middleware.AuthMiddleware())
	{
		u.POST("/:username", user.Create)
		u.GET("/:username", user.Get)
		u.DELETE("/:id", user.Delete)
		u.PUT("/:id", user.Update)
		u.GET("", user.List)
	}

	// the health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	// backend index
	g.LoadHTMLFiles("backend/dist/index.html")
	g.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	return g
}
