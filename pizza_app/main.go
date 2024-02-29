package main

import (
	"os"
	"pizza-app/pizza_app/api/middleware"
	"pizza-app/pizza_app/api/routes"
)

func main() {
	r := routes.SetupRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := r.Run(":" + port)
	middleware.HandleError(err)
}
