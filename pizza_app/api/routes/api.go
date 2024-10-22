package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	"net/http"

	bigaDoughHandler "pizza-app/pizza_app/api/handlers/dough/biga"
	directDoughHandler "pizza-app/pizza_app/api/handlers/dough/direct"
	pideDoughHandler "pizza-app/pizza_app/api/handlers/dough/pide"
	poolishDoughHandler "pizza-app/pizza_app/api/handlers/dough/poolish"
	sourDoughHandler "pizza-app/pizza_app/api/handlers/dough/sourdough"
	"pizza-app/pizza_app/api/middleware"
)

func SetupRouter() *gin.Engine {

	// IP RATE LIMITER
	rate, err := limiter.NewRateFromFormatted("5-S")
	middleware.HandleError(err)

	store := memory.NewStore()
	ipRate := limiter.New(store, rate)
	middlewareIpRate := mgin.NewMiddleware(ipRate)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.ForwardedByClientIP = true
	router.Use(middlewareIpRate)

	router.LoadHTMLGlob("templates/**/*")

	// add favicon
	router.StaticFile("/favicon.ico", "./static/img/favicon.ico")
	router.Static("/css", "./static/css")
	router.Static("/js", "./static/js")

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusNoContent, nil)
	})

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "biga.html", nil)
	})
	router.GET("biga", func(c *gin.Context) {
		c.HTML(http.StatusOK, "biga.html", nil)
	})
	router.GET("poolish", func(c *gin.Context) {
		c.HTML(http.StatusOK, "poolish.html", nil)
	})

	router.GET("/direct", func(c *gin.Context) {
		c.HTML(http.StatusOK, "direct.html", nil)
	})

	router.GET("/sourdough", func(c *gin.Context) {
		c.HTML(http.StatusOK, "sourdough.html", nil)
	})

	router.GET("/pide", func(c *gin.Context) {
		c.HTML(http.StatusOK, "pide.html", nil)
	})

	router.POST("/direct", directDoughHandler.HandleDirectDough)
	router.POST("/biga", bigaDoughHandler.HandleBigaDough)
	router.POST("/poolish", poolishDoughHandler.HandlePoolishDough)
	router.POST("/sourdough", sourDoughHandler.HandleSourDough)
	router.POST("/pide", pideDoughHandler.HandlePideDough)

	return router
}
