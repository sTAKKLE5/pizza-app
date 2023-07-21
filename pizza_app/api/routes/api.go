package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	bigaDoughHandler "pizza-app/pizza_app/api/handlers/dough/biga"
	directDoughHandler "pizza-app/pizza_app/api/handlers/dough/direct"
	poolishDoughHandler "pizza-app/pizza_app/api/handlers/dough/poolish"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("/dough/direct", func(c *gin.Context) {
		c.HTML(http.StatusOK, "direct.html", gin.H{
			"url": "http://localhost:8080/dough/direct", // FIXME: This is hardcoded
		})
	})

	router.POST("/dough/direct", directDoughHandler.HandleDirectDough)
	router.POST("/dough/biga", bigaDoughHandler.HandleBigaDough)
	router.POST("/dough/poolish", poolishDoughHandler.HandlePoolishDough)

	return router
}
