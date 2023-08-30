package handlers

import (
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"pizza-app/pizza_app/api/middleware"
	requestModel "pizza-app/pizza_app/models/request"
	tomatoSauceModel "pizza-app/pizza_app/models/sauce/tomato"
)

func HandleTomatoSauce(c *gin.Context) {
	var request requestModel.SauceRequest
	request.PizzaAmount = middleware.ParseFormInt(c, "pizzaAmount")

	recipe := generateTomatoRecipe(request)
	// Respond with the recipe data using an HTML template.
	c.HTML(http.StatusOK, "tomato_recipe.html", recipe)
}

func generateTomatoRecipe(request requestModel.SauceRequest) tomatoSauceModel.TomatoSauceResponse {
	// Calculate number of cans and basil based on the amount of pizzas
	canCount := int(math.Ceil(float64(request.PizzaAmount) / 2.0))
	basilCount := canCount * 5 // 5 basil leaves per 2 pizzas

	recipe := tomatoSauceModel.TomatoSauceResponse{
		TomatoCan: canCount,
		Basil:     basilCount,
	}

	return recipe
}
