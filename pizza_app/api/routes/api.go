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
	router.Static("/img", "./static/img")
	router.Static("/js", "./static/js")

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusNoContent, gin.H{})
	})

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

	type ResponsePepper struct {
		Id         string `json:"id"`
		Variety    string `json:"variety"`
		Generation string `json:"generation"`
		Specimen   string `json:"specimen"`
	}

	router.GET("/pepper-id/:id", func(c *gin.Context) {
		peppers := make(map[string]string)
		generation := make(map[string]string)
		specimen := make(map[string]string)

		peppers["1"] = "Aji Charapita"
		peppers["2"] = "Aji Mango Stumpy"
		peppers["3"] = "Habanada Orange"

		peppers["4-1"] = "Pimenta Da Neyda X West Indian Yellow Habanero"
		generation["4-1"] = "F3"
		specimen["4-1"] = "1"

		peppers["4-2"] = "Pimenta Da Neyda X West Indian Yellow Habanero"
		generation["4-2"] = "F3"
		specimen["4-2"] = "2"

		peppers["4-3"] = "Pimenta Da Neyda X West Indian Yellow Habanero"
		generation["4-3"] = "F3"
		specimen["4-3"] = "3"

		peppers["4-4"] = "Pimenta Da Neyda X West Indian Yellow Habanero"
		generation["4-4"] = "F3"
		specimen["4-4"] = "4"

		peppers["4-5"] = "Pimenta Da Neyda X West Indian Yellow Habanero"
		generation["4-5"] = "F3"
		specimen["4-5"] = "5"

		peppers["4-6"] = "Pimenta Da Neyda X West Indian Yellow Habanero"
		generation["4-6"] = "F3"
		specimen["4-6"] = "6"

		peppers["5"] = "Pimenta Da Neyda"

		id := c.Param("id")

		c.JSON(200, gin.H{
			"variety":    peppers[id],
			"generation": generation[id],
			"specimen":   specimen[id],
		})
	})

	return router
}
