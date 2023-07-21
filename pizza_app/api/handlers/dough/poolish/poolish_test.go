package handlers

import (
	"github.com/stretchr/testify/assert"
	poolishDoughModel "pizza-app/pizza_app/models/dough/poolish"
	requestModel "pizza-app/pizza_app/models/request"
	"testing"
)

// Test function for generatePoolishRecipe
func TestGeneratePoolishRecipe(t *testing.T) {
	request := requestModel.DoughRequest{
		DoughBallAmount: 1,
		Hydration:       50,
		DoughBallWeight: 1500,
	}

	expectedResult := poolishDoughModel.PoolishDoughResponse{
		Poolish: poolishDoughModel.PoolishDough{
			Flour:           500,
			Water:           500,
			InstantDryYeast: 0.15,
		},
		MainDough: poolishDoughModel.MainDough{
			Flour:           500,
			Water:           0,
			Salt:            25,
			InstantDryYeast: 0.15,
		},
		DoughBallAmount: 1,
		Hydration:       50,
		DoughBallWeight: 1500,
	}

	result := generatePoolishRecipe(request)

	assert.Equal(t, expectedResult, result, "The two values should be the same.")
}
