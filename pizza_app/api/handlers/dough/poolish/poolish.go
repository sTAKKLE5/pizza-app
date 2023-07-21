package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	// Implement your recipe generating logic here
	var recipe poolishDoughModel.PoolishDoughResponse

	return recipe
}
