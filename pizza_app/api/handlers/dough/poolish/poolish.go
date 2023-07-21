package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	utils "pizza-app/pizza_app/api/middleware"
	poolishDoughModel "pizza-app/pizza_app/models/dough/poolish"
	requestModel "pizza-app/pizza_app/models/request"
)

func HandlePoolishDough(c *gin.Context) {
	var request requestModel.DoughRequest

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Process the data and generate the recipe.
	recipe := generatePoolishRecipe(request)

	// Respond with the recipe data.
	c.JSON(http.StatusOK, recipe)
}

func generatePoolishRecipe(request requestModel.DoughRequest) poolishDoughModel.PoolishDoughResponse {
	fmt.Printf("request: %v\n", request)
	saltPerKg := 25.  // 2.5% of 1 kilogram of flour, fixed amount
	yeastPerKg := 0.3 // 0.03% of 1 kilogram of flour, fixed amount

	flour := request.DoughBallWeight / (1 + (request.Hydration / 100))
	water := flour * (request.Hydration / 100)
	salt := (flour / 1000) * saltPerKg   // Salt in relation to the calculated flour amount
	yeast := (flour / 1000) * yeastPerKg // Yeast in relation to the calculated flour amount

	poolishFlour := flour / 2
	poolishWater := poolishFlour // Poolish is 100% hydration
	poolishYeast := yeast / 2    // Half of yeast is part of poolish

	finalDoughFlour := flour - poolishFlour
	finalDoughWater := water - poolishWater
	finalDoughYeast := yeast / 2 // Other half of yeast is part of the final dough

	poolish := poolishDoughModel.PoolishDough{
		Flour:           utils.RoundToDecimal(poolishFlour, 2) * float64(request.DoughBallAmount),
		Water:           utils.RoundToDecimal(poolishWater, 2) * float64(request.DoughBallAmount),
		InstantDryYeast: poolishYeast * float64(request.DoughBallAmount),
	}

	mainDough := poolishDoughModel.MainDough{
		Flour:           utils.RoundToDecimal(finalDoughFlour, 2) * float64(request.DoughBallAmount),
		Water:           utils.RoundToDecimal(finalDoughWater, 2) * float64(request.DoughBallAmount),
		Salt:            utils.RoundToDecimal(salt, 2) * float64(request.DoughBallAmount),
		InstantDryYeast: finalDoughYeast * float64(request.DoughBallAmount),
	}

	recipe := poolishDoughModel.PoolishDoughResponse{
		Poolish:         poolish,
		MainDough:       mainDough,
		DoughBallAmount: request.DoughBallAmount,
		Hydration:       utils.RoundToDecimal(request.Hydration, 2),
		DoughBallWeight: utils.RoundToDecimal(request.DoughBallWeight, 2),
	}

	fmt.Printf("recipe: %v\n", recipe)
	return recipe
}
