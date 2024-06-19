package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pizza-app/pizza_app/api/middleware"
	pideDoughModel "pizza-app/pizza_app/models/dough/pide"
	requestModel "pizza-app/pizza_app/models/request"
)

func HandlePideDough(c *gin.Context) {
	var request requestModel.DoughRequestPide

	request.DoughBallAmount = middleware.ParseFormInt(c, "doughBallAmountPide")

	// Process the data and generate the recipe.
	recipe := generatePideRecipe(request)

	// Respond with the recipe data using an HTML template.
	c.HTML(http.StatusOK, "pide_recipe.html", recipe)
}

func generatePideRecipe(request requestModel.DoughRequestPide) pideDoughModel.PideDoughResponse {
	// Constants for the recipe
	const flourPerDoughBall = 62.5
	const yeastPerDoughBall = 1.2
	const waterPerDoughBall = 16.3
	const oilPerDoughBall = 0.1
	const saltPerKiloFlour = 25.0

	// Calculate the total amount of each ingredient
	totalFlour := flourPerDoughBall * float64(request.DoughBallAmount)
	totalYeast := yeastPerDoughBall * float64(request.DoughBallAmount)
	totalWater := waterPerDoughBall * float64(request.DoughBallAmount)
	totalOil := oilPerDoughBall * float64(request.DoughBallAmount)
	totalSalt := (totalFlour / 1000.0) * saltPerKiloFlour // Convert flour to kilos before calculating salt

	// Round the amounts to 2 decimal places
	totalFlour = middleware.RoundToDecimal(totalFlour, 2)
	totalYeast = middleware.RoundToDecimal(totalYeast, 2)
	totalWater = middleware.RoundToDecimal(totalWater, 2)
	totalOil = middleware.RoundToDecimal(totalOil, 2)
	totalSalt = middleware.RoundToDecimal(totalSalt, 2)

	// Create the MainDough object
	mainDough := pideDoughModel.MainDough{
		Flour:      totalFlour,
		Water:      totalWater,
		FreshYeast: totalYeast,
		Salt:       totalSalt,
		OliveOil:   totalOil,
	}

	// Create the PideDoughResponse object
	recipe := pideDoughModel.PideDoughResponse{
		MainDough:       mainDough,
		DoughBallAmount: request.DoughBallAmount,
		DoughBallWeight: 100, // Each dough ball weighs 100 grams
	}

	return recipe
}
