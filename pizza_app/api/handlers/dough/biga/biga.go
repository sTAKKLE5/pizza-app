package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pizza-app/pizza_app/api/middleware"
	bigaDoughModel "pizza-app/pizza_app/models/dough/biga"
	requestModel "pizza-app/pizza_app/models/request"
)

func HandleBigaDough(c *gin.Context) {
	var request requestModel.DoughRequest

	request.DoughBallWeight = middleware.ParseFormFloat64(c, "doughBallWeight")
	request.Hydration = middleware.ParseFormFloat64(c, "hydration")
	request.DoughBallAmount = middleware.ParseFormInt(c, "doughBallAmount")

	// Process the data and generate the recipe.
	recipe := generateBigaRecipe(request)

	// Respond with the recipe data using an HTML template.
	c.HTML(http.StatusOK, "biga_recipe.html", recipe)
}

func generateBigaRecipe(request requestModel.DoughRequest) bigaDoughModel.BigaDoughResponse {
	saltPerKg := 25.  // 2.5% of 1 kilogram of flour, fixed amount
	yeastPerKg := 0.3 // 0.03% of 1 kilogram of flour, fixed amount

	flour := request.DoughBallWeight / (1 + (request.Hydration / 100))
	water := flour * (request.Hydration / 100)
	salt := (flour / 1000) * saltPerKg   // Salt in relation to the calculated flour amount
	yeast := (flour / 1000) * yeastPerKg // Yeast in relation to the calculated flour amount

	bigaFlour := flour * 0.50
	bigaWater := bigaFlour * 0.45
	bigaYeast := yeast // All yeast is part of biga

	finalDoughFlour := flour - bigaFlour
	finalDoughWater := water - bigaWater

	biga := bigaDoughModel.BigaDough{
		Flour:           middleware.RoundToDecimal(bigaFlour*float64(request.DoughBallAmount), 2),
		Water:           middleware.RoundToDecimal(bigaWater*float64(request.DoughBallAmount), 2),
		InstantDryYeast: bigaYeast * float64(request.DoughBallAmount),
	}

	mainDough := bigaDoughModel.MainDough{
		Flour: middleware.RoundToDecimal(finalDoughFlour*float64(request.DoughBallAmount), 2),
		Water: middleware.RoundToDecimal(finalDoughWater*float64(request.DoughBallAmount), 2),
		Salt:  middleware.RoundToDecimal(salt*float64(request.DoughBallAmount), 2),
	}

	recipe := bigaDoughModel.BigaDoughResponse{
		Biga:                biga,
		MainDough:           mainDough,
		DoughBallAmount:     request.DoughBallAmount,
		Hydration:           request.Hydration,
		BigaHydration:       bigaWater / bigaFlour * 100,
		BigaWaterPercentage: bigaWater / water * 100,
		DoughBallWeight:     request.DoughBallWeight,
	}

	return recipe
}
