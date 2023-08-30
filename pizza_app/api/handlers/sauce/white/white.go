package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pizza-app/pizza_app/api/middleware"
	requestModel "pizza-app/pizza_app/models/request"
	whiteSauceModel "pizza-app/pizza_app/models/sauce/white"
)

func HandleWhiteSauce(c *gin.Context) {
	var request requestModel.SauceRequest
	request.PizzaAmount = middleware.ParseFormInt(c, "pizzaAmount")

	recipe := generateWhiteRecipe(request)
	// Respond with the recipe data using an HTML template.
	c.HTML(http.StatusOK, "white_recipe.html", recipe)
}

func generateWhiteRecipe(request requestModel.SauceRequest) whiteSauceModel.WhiteSauceResponse {
	multiplier := float64(request.PizzaAmount) / 3.0 // This will give 1/3 for 1 pizza, 2/3 for 2 pizzas, 1 for 3 pizzas, and so on.

	recipe := whiteSauceModel.WhiteSauceResponse{
		Butter:         middleware.RoundToDecimal(28*multiplier, 2),
		Flour:          middleware.RoundToDecimal(16*multiplier, 2),
		Milk:           middleware.RoundToDecimal(300*multiplier, 2),
		SeaSalt:        middleware.RoundToDecimal(1.5*multiplier, 2),
		BlackPepper:    middleware.RoundToDecimal(0.3*multiplier, 2),
		GarlicCloves:   middleware.RoundToDecimal(8*multiplier, 2),
		ParmesanCheese: middleware.RoundToDecimal(25*multiplier, 2),
	}

	return recipe
}
