package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	bigaDoughModel "pizza-app/pizza_app/models/dough/biga"
	requestModel "pizza-app/pizza_app/models/request"
)

func HandleBigaDough(c *gin.Context) {
	var request requestModel.DoughRequest

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Process the data and generate the recipe.
	recipe := generateBigaRecipe(request)

	// Respond with the recipe data.
	c.JSON(http.StatusOK, recipe)
}

func generateBigaRecipe(request requestModel.DoughRequest) bigaDoughModel.BigaDoughResponse {
	// Implement your recipe generating logic here
	var recipe bigaDoughModel.BigaDoughResponse

	return recipe
}
