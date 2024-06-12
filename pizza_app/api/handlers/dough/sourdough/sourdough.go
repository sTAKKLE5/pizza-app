package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pizza-app/pizza_app/api/middleware"
	"pizza-app/pizza_app/models/dough/sourdough"
	requestModel "pizza-app/pizza_app/models/request"
)

func HandleSourDough(c *gin.Context) {
	var request requestModel.DoughRequest

	request.DoughBallWeight = middleware.ParseFormFloat64(c, "doughBallWeightSourdough")
	request.Hydration = middleware.ParseFormFloat64(c, "hydrationSourdough")
	request.DoughBallAmount = middleware.ParseFormInt(c, "doughBallAmountSourdough")

	// Process the data and generate the recipe.
	recipe := generateDirectRecipe(request)

	// Respond with the recipe data using an HTML template.
	c.HTML(http.StatusOK, "sourdough_recipe.html", recipe)
}

func generateDirectRecipe(request requestModel.DoughRequest) sourdough.SourDoughResponse {
	// Constants for the base recipe
	const baseFlour = 550.0
	const baseWater = 300.0
	const baseSourdough = 200.0
	const baseSalt = 20.0
	const baseDoughBallWeight = 270.0
	const baseDoughBallAmount = 4

	// Calculate the scaling factor based on the desired dough ball weight and amount
	scale := (request.DoughBallWeight * float64(request.DoughBallAmount)) / (baseDoughBallWeight * baseDoughBallAmount)

	// Calculate the ingredients for the main dough
	totalFlour := baseFlour * scale
	totalSalt := baseSalt * scale
	sourDough := baseSourdough * scale

	// Calculate the water for the sourdough
	sourDoughWater := sourDough / 2

	// Calculate the remaining water for the main dough to achieve the desired hydration
	totalWater := (totalFlour * request.Hydration / 100) - sourDoughWater

	mainDough := sourdough.MainDough{
		Flour:     middleware.RoundToDecimal(totalFlour, 2),
		Water:     middleware.RoundToDecimal(totalWater, 2),
		SourDough: middleware.RoundToDecimal(sourDough, 2),
		Salt:      middleware.RoundToDecimal(totalSalt, 2),
	}

	recipe := sourdough.SourDoughResponse{
		MainDough:       mainDough,
		DoughBallAmount: request.DoughBallAmount,
		Hydration:       request.Hydration,
		DoughBallWeight: request.DoughBallWeight,
	}

	return recipe
}
