package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pizza-app/pizza_app/api/middleware"
	bigaDoughModel "pizza-app/pizza_app/models/dough/biga"
	requestModel "pizza-app/pizza_app/models/request"
)

func HandleBigaDough(c *gin.Context) {
	var request requestModel.BigaDoughRequest

	request.DoughBallWeight = middleware.ParseFormFloat64(c, "doughBallWeight")
	request.Hydration = middleware.ParseFormFloat64(c, "hydration")
	request.DoughBallAmount = middleware.ParseFormInt(c, "doughBallAmount")
	request.FermentationHours = middleware.ParseFormInt(c, "fermentationHours")
	request.FlourPercentage = middleware.ParseFormFloat64(c, "percentageBiga")
	request.BigaHydration = middleware.ParseFormFloat64(c, "percentageBigaHydration")
	request.FermentationTemperature = middleware.ParseFormInt(c, "fermentationTemperature")

	// Process the data and generate the recipe.
	recipe := generateBigaRecipe(request)

	// Respond with the recipe data using an HTML template.
	c.HTML(http.StatusOK, "biga_recipe.html", recipe)
}

func generateBigaRecipe(request requestModel.BigaDoughRequest) bigaDoughModel.BigaDoughResponse {
	saltPerKg := 25. // 2.5% of 1 kilogram of flour, fixed amount

	// yeast depends on the fermentation time
	var yeastPerKg float64
	if request.FermentationHours <= 12 {
		yeastPerKg = 3.0
	} else if request.FermentationHours <= 24 {
		yeastPerKg = 1.5
	} else {
		yeastPerKg = 0.75
	}

	flour := request.DoughBallWeight / (1 + (request.Hydration / 100))
	water := flour * (request.Hydration / 100)
	salt := (flour / 1000) * saltPerKg   // Salt in relation to the calculated flour amount
	yeast := (flour / 1000) * yeastPerKg // Yeast in relation to the calculated flour amount

	bigaFlour := flour * (request.FlourPercentage / 100) // Use the specified percentage for biga flour
	bigaWater := bigaFlour * request.BigaHydration / 100
	bigaYeast := yeast // All yeast is part of biga

	finalDoughFlour := flour - bigaFlour
	finalDoughWater := water - bigaWater

	biga := bigaDoughModel.BigaDough{
		Flour:           middleware.RoundToDecimal(bigaFlour*float64(request.DoughBallAmount), 2),
		Water:           middleware.RoundToDecimal(bigaWater*float64(request.DoughBallAmount), 2),
		InstantDryYeast: middleware.RoundToDecimal(bigaYeast*float64(request.DoughBallAmount)/2, 2), // half of the yeast for instant dry yeast
		FreshYeast:      middleware.RoundToDecimal(bigaYeast*float64(request.DoughBallAmount), 2),
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
