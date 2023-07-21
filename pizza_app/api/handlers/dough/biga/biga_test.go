package handlers

import (
	"testing"

	bigaDoughModel "pizza-app/pizza_app/models/dough/biga"
	requestModel "pizza-app/pizza_app/models/request"
)

func TestGenerateBigaRecipe(t *testing.T) {
	req := requestModel.DoughRequest{
		DoughBallAmount: 1,
		Hydration:       50,
		DoughBallWeight: 1500,
	}

	expected := bigaDoughModel.BigaDoughResponse{
		Biga: bigaDoughModel.BigaDough{
			Flour:           500,
			Water:           225,
			InstantDryYeast: 0.3,
		},
		MainDough: bigaDoughModel.MainDough{
			Flour: 500,
			Water: 275,
			Salt:  25,
		},
		DoughBallAmount:     req.DoughBallAmount,
		Hydration:           req.Hydration,
		BigaHydration:       45,
		BigaWaterPercentage: 45,
		DoughBallWeight:     req.DoughBallWeight,
	}

	result := generateBigaRecipe(req)

	// you might want to adjust the comparison to tolerate tiny inaccuracies due to float calculation
	if result != expected {
		t.Errorf("Test failed. \nExpected: %+v, \nGot: %+v", expected, result)
	}
}
