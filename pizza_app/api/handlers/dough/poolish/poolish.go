package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pizza-app/pizza_app/api/middleware"
	poolishDoughModel "pizza-app/pizza_app/models/dough/poolish"
	requestModel "pizza-app/pizza_app/models/request"
)

func HandlePoolishDough(c *gin.Context) {
	var request requestModel.PoolishDoughRequest

	request.DoughBallWeight = middleware.ParseFormFloat64(c, "doughBallWeight")
	request.Hydration = middleware.ParseFormFloat64(c, "hydration")
	request.DoughBallAmount = middleware.ParseFormInt(c, "doughBallAmount")
	request.PoolishPercentage = middleware.ParseFormFloat64(c, "poolishPercentage")

	// Process the data and generate the recipe.
	recipe := generatePoolishRecipe(request)

	// Respond with the recipe data using an HTML template.
	c.HTML(http.StatusOK, "poolish_recipe.html", recipe)
}

func generatePoolishRecipe(request requestModel.PoolishDoughRequest) poolishDoughModel.PoolishDoughResponse {
	saltPerKg := 25.  // 2.5% of 1 kilogram of flour, fixed amount
	yeastPerKg := 0.3 // 0.03% of 1 kilogram of flour, fixed amount

	flour := request.DoughBallWeight / (1 + (request.Hydration / 100))
	water := flour * (request.Hydration / 100)
	salt := (flour / 1000) * saltPerKg   // Salt in relation to the calculated flour amount
	yeast := (flour / 1000) * yeastPerKg // Yeast in relation to the calculated flour amount

	poolishWater := water * (request.PoolishPercentage / 100)
	poolishFlour := poolishWater // Poolish is 100% hydration
	poolishYeast := yeast / 2    // Half of yeast is part of poolish

	finalDoughFlour := flour - poolishFlour
	finalDoughWater := water - poolishWater
	finalDoughYeast := yeast / 2 // Other half of yeast is part of the final dough

	poolish := poolishDoughModel.PoolishDough{
		Flour:           middleware.RoundToDecimal(poolishFlour*float64(request.DoughBallAmount), 2),
		Water:           middleware.RoundToDecimal(poolishWater*float64(request.DoughBallAmount), 2),
		InstantDryYeast: middleware.RoundToDecimal(poolishYeast*float64(request.DoughBallAmount), 2),
	}

	mainDough := poolishDoughModel.MainDough{
		Flour:           middleware.RoundToDecimal(finalDoughFlour*float64(request.DoughBallAmount), 2),
		Water:           middleware.RoundToDecimal(finalDoughWater*float64(request.DoughBallAmount), 2),
		Salt:            middleware.RoundToDecimal(salt*float64(request.DoughBallAmount), 2),
		InstantDryYeast: middleware.RoundToDecimal(finalDoughYeast*float64(request.DoughBallAmount), 2),
	}

	recipe := poolishDoughModel.PoolishDoughResponse{
		Poolish:           poolish,
		MainDough:         mainDough,
		DoughBallAmount:   request.DoughBallAmount,
		Hydration:         request.Hydration,
		DoughBallWeight:   request.DoughBallWeight,
		PoolishPercentage: request.PoolishPercentage,
	}

	return recipe
}
