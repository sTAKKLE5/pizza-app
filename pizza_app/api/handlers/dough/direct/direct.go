package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	directDoughModel "pizza-app/pizza_app/models/dough/direct"
	requestModel "pizza-app/pizza_app/models/request"
)

func HandleDirectDough(c *gin.Context) {
	var request requestModel.DoughRequest

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Process the data and generate the recipe.
	recipe := generateDirectRecipe(request)

	// Respond with the recipe data.
	c.JSON(http.StatusOK, recipe)
}

func generateDirectRecipe(request requestModel.DoughRequest) directDoughModel.DirectDoughResponse {
	saltPerKg := 25.  // 2.5% of 1 kilogram of flour, fixed amount
	yeastPerKg := 0.3 // 0.03% of 1 kilogram of flour, fixed amount

	flour := request.DoughBallWeight / (1 + (request.Hydration / 100))
	water := flour * (request.Hydration / 100)
	salt := (flour / 1000) * saltPerKg   // Salt in relation to the calculated flour amount
	yeast := (flour / 1000) * yeastPerKg // Yeast in relation to the calculated flour amount

	totalFlour := flour * float64(request.DoughBallAmount)
	totalWater := water * float64(request.DoughBallAmount)
	totalSalt := salt * float64(request.DoughBallAmount)
	totalYeast := yeast * float64(request.DoughBallAmount)

	mainDough := directDoughModel.MainDough{
		Flour:           totalFlour,
		Water:           totalWater,
		InstantDryYeast: totalYeast,
		Salt:            totalSalt,
	}

	recipe := directDoughModel.DirectDoughResponse{
		MainDough:       mainDough,
		DoughBallAmount: request.DoughBallAmount,
		Hydration:       request.Hydration,
		DoughBallWeight: request.DoughBallWeight,
	}

	return recipe
}
