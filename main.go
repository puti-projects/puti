package main

import (
	"GdPHP/router"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// create the gin engine
	g := gin.New()

	// gin middlewares
	middlewares := []gin.HandlerFunc{}

	// routes
	router.Load(g, middlewares...)

	log.Printf("Start to listening the incoming requests on http address: %s", ":8000")
	log.Printf(http.ListenAndServe(":8000", g).Error())

}
