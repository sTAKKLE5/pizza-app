package handlers

import (
	"github.com/stretchr/testify/assert"
	requestModel "pizza-app/pizza_app/models/request"
	"testing"
)

func TestGenerateDirectRecipe(t *testing.T) {
	// Define test request
	request := requestModel.DoughRequest{
		DoughBallAmount: 5,
		Hydration:       50,
		DoughBallWeight: 1500,
	}

	// Generate recipe
	recipe := generateDirectRecipe(request)

	// Assertions
	assert.Equal(t, float64(5000), recipe.MainDough.Flour, "The calculated flour amount should be 1000g")
	assert.Equal(t, float64(2500), recipe.MainDough.Water, "The calculated water amount should be 500g")
	assert.Equal(t, float64(125), recipe.MainDough.Salt, "The calculated salt amount should be 25g")
	assert.Equal(t, float64(1.5), recipe.MainDough.InstantDryYeast, "The calculated yeast amount should be 0.3g")
	assert.Equal(t, 5, recipe.DoughBallAmount, "The dough ball amount should be 1")
	assert.Equal(t, float64(50), recipe.Hydration, "The hydration should be 50")
	assert.Equal(t, float64(1500), recipe.DoughBallWeight, "The dough ball weight should be 1000g")
}
