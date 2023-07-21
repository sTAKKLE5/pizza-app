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
	// Implement your recipe generating logic here
	var recipe directDoughModel.DirectDoughResponse

	return recipe
}
