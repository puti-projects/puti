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
	g.Static("/static", "web/backend/static")
	g.Static("/uploads", "uploads/")
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
		a.GET("/:id", article.Get)
		a.POST("", article.Create)
		a.PUT("/:id", article.Update)
		a.DELETE("/:id", article.Delete)
	}

	p := api.Group("/page")
	p.Use(middleware.AuthMiddleware())
	{
		p.GET("", page.List)
		p.GET("/:id", page.Get)
		p.POST("", page.Create)
		p.PUT("/:id", page.Update)
		p.DELETE("/:id", page.Delete)
	}

	t := api.Group("/taxonomy")
	t.Use(middleware.AuthMiddleware())
	{
		t.POST("/:name", taxonomy.Create)
		t.GET("/:id", taxonomy.Get)
		t.DELETE("/:id", taxonomy.Delete)
		t.PUT("/:id", taxonomy.Update)
		t.GET("", taxonomy.List)
	}

	m := api.Group("/media")
	m.Use(middleware.AuthMiddleware())
	{
		m.GET("/:id", media.Detail)
		m.GET("", media.List)
		m.POST("", media.Upload)
		m.DELETE("/:id", media.Delete)
		m.PUT("/:id", media.Update)
	}

	o := api.Group("/option")
	o.Use(middleware.AuthMiddleware())
	{
		o.GET("", option.List)
		o.PUT("", option.Update)
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
	g.LoadHTMLFiles("web/backend/index.html")
	g.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	return g
}
