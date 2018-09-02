package router

import (
	"net/http"

	"puti/handler/article"
	"puti/handler/auth"
	"puti/handler/media"
	"puti/handler/sd"
	"puti/handler/user"
	"puti/router/middleware"

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

	api.POST("/login", auth.Login)
	api.GET("/token", auth.Info)

	u := api.Group("/user")
	u.Use(middleware.AuthMiddleware())
	{
		u.POST("/:username", user.Create)
		u.GET("/:username", user.Get)
		u.DELETE("/:id", user.Delete)
		u.PUT("/:id", user.Update)
		u.GET("", user.List)
	}

	av := api.Group("/avatar")
	av.Use(middleware.AuthMiddleware())
	{
		av.POST("", user.Avatar)
	}

	a := api.Group("/article")
	a.Use(middleware.AuthMiddleware())
	{
		a.GET("", article.List)
	}

	m := api.Group("/media")
	m.Use(middleware.AuthMiddleware())
	{
		m.GET("/:id", media.Detail)
		m.GET("", media.List)
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
