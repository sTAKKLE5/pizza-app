package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	"net/http"

	bigaDoughHandler "pizza-app/pizza_app/api/handlers/dough/biga"
	directDoughHandler "pizza-app/pizza_app/api/handlers/dough/direct"
	poolishDoughHandler "pizza-app/pizza_app/api/handlers/dough/poolish"
	tomato "pizza-app/pizza_app/api/handlers/sauce/tomato"
	white "pizza-app/pizza_app/api/handlers/sauce/white"
	"pizza-app/pizza_app/api/middleware"
)

func SetupRouter() *gin.Engine {

	// IP RATE LIMITER
	rate, err := limiter.NewRateFromFormatted("5-S")
	middleware.HandleError(err)

	store := memory.NewStore()
	ipRate := limiter.New(store, rate)
	middlewareIprate := mgin.NewMiddleware(ipRate)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.ForwardedByClientIP = true
	router.Use(middlewareIprate)

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

	router.GET("/sauce/tomato", func(c *gin.Context) {
		c.HTML(http.StatusOK, "tomato.html", gin.H{
			"url":         "http://localhost:8080/sauce/tomato",
			"currentPath": c.Request.URL.Path,
		})
	})

	router.GET("/sauce/white", func(c *gin.Context) {
		c.HTML(http.StatusOK, "white.html", gin.H{
			"url":         "http://localhost:8080/sauce/white",
			"currentPath": c.Request.URL.Path,
		})
	})

	router.POST("/dough/direct", directDoughHandler.HandleDirectDough)
	router.POST("/dough/biga", bigaDoughHandler.HandleBigaDough)
	router.POST("/dough/poolish", poolishDoughHandler.HandlePoolishDough)
	router.POST("/sauce/tomato", tomato.HandleTomatoSauce)
	router.POST("/sauce/white", white.HandleWhiteSauce)

	return router
}
