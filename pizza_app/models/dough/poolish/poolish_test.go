package poolish

import (
	"testing"
)

func TestPoolishDough(t *testing.T) {
	poolishDough := PoolishDough{
		Flour:           120.0,
		Water:           80.0,
		InstantDryYeast: 0.50,
	}

	if poolishDough.Flour != 120.0 {
		t.Errorf("Expected flour to be 120.0, but got %f", poolishDough.Flour)
	}

}

func TestMainDough(t *testing.T) {
	mainDough := MainDough{
		Flour:           120.0,
		Water:           80.0,
		InstantDryYeast: 0.50,
		Salt:            1.0,
	}

	if mainDough.Flour != 120.0 {
		t.Errorf("Expected flour to be 120.0, but got %f", mainDough.Flour)
	}

}

func TestPoolishDoughResponse(t *testing.T) {
	poolishResponse := PoolishDoughResponse{
		Poolish: PoolishDough{
			Flour:           120.0,
			Water:           80.0,
			InstantDryYeast: 0.50,
		},
		MainDough: MainDough{
			Flour:           120.0,
			Water:           80.0,
			InstantDryYeast: 0.50,
			Salt:            1.0,
		},
		DoughBallAmount:        4,
		Hydration:              70.0,
		DoughBallWeight:        270.0,
		PoolishHydration:       75.0,
		PoolishWaterPercentage: 80.0,
		PoolishYeastPercentage: 2.0,
	}

	if poolishResponse.Poolish.Flour != 120.0 {
		t.Errorf("Expected Poolish flour to be 120.0, but got %f", poolishResponse.Poolish.Flour)
	}

}
