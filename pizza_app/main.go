package main

import (
	"pizza-app/pizza_app/api/middleware"
	"pizza-app/pizza_app/api/routes"
)

func main() {
	r := routes.SetupRouter()
	err := r.Run()
	middleware.HandleError(err)
}
