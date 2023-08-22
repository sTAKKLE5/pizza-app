package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	bigaDoughHandler "pizza-app/pizza_app/api/handlers/dough/biga"
	directDoughHandler "pizza-app/pizza_app/api/handlers/dough/direct"
	poolishDoughHandler "pizza-app/pizza_app/api/handlers/dough/poolish"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.LoadHTMLGlob("templates/**/*")

	// add favicon
	router.StaticFile("/favicon.ico", "./static/img/favicon.ico")
	router.Static("/css", "./static/css")
	router.Static("/img", "./static/img")
	router.Static("/js", "./static/js")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"currentPath": c.Request.URL.Path})
	})
	router.GET("dough/biga", func(c *gin.Context) {
		c.HTML(http.StatusOK, "biga.html", gin.H{
			"url":         "http://localhost:8080/dough/biga",
			"currentPath": c.Request.URL.Path,
		})
	})
	router.GET("dough/poolish", func(c *gin.Context) {
		c.HTML(http.StatusOK, "poolish.html", gin.H{
			"url":         "http://localhost:8080/dough/poolish",
			"currentPath": c.Request.URL.Path,
		})
	})

	router.GET("/dough/direct", func(c *gin.Context) {
		c.HTML(http.StatusOK, "direct.html", gin.H{
			"url":         "http://localhost:8080/dough/direct",
			"currentPath": c.Request.URL.Path,
		})
	})

	router.POST("/dough/direct", directDoughHandler.HandleDirectDough)
	router.POST("/dough/biga", bigaDoughHandler.HandleBigaDough)
	router.POST("/dough/poolish", poolishDoughHandler.HandlePoolishDough)

	return router
}
