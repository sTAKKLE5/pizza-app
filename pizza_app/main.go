package main

import (
	utils "pizza-app/pizza_app/api/middleware"
	"pizza-app/pizza_app/api/routes"
)

func main() {
	r := routes.SetupRouter()
	err := r.Run()
	utils.HandleError(err)
}
