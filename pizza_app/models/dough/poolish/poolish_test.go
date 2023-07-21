package poolish

import (
	"testing"
)

func TestPoolishDough(t *testing.T) {
	poolishDough := PoolishDough{
		Flour:           "120.00",
		Water:           "80.00",
		InstantDryYeast: "0.50",
	}

	if poolishDough.Flour != "120.00" {
		t.Errorf("Expected flour to be '120.00', but got '%s'", poolishDough.Flour)
	}

}

func TestMainDough(t *testing.T) {
	mainDough := MainDough{
		Flour:           "120.00",
		Water:           "80.00",
		InstantDryYeast: "0.50",
		Salt:            "1.00",
	}

	if mainDough.Flour != "120.00" {
		t.Errorf("Expected flour to be '120.00', but got '%s'", mainDough.Flour)
	}

}

func TestPoolishDoughResponse(t *testing.T) {
	poolishResponse := PoolishDoughResponse{
		Poolish: PoolishDough{
			Flour:           "120.00",
			Water:           "80.00",
			InstantDryYeast: "0.50",
		},
		MainDough: MainDough{
			Flour:           "120.00",
			Water:           "80.00",
			InstantDryYeast: "0.50",
			Salt:            "1.00",
		},
		DoughBallAmount:        "4",
		Hydration:              "70.00",
		DoughBallWeight:        "270.00",
		PoolishHydration:       "75.00",
		PoolishWaterPercentage: "80.00",
		PoolishYeastPercentage: "2.00",
	}

	if poolishResponse.Poolish.Flour != "120.00" {
		t.Errorf("Expected Poolish flour to be '120.00', but got '%s'", poolishResponse.Poolish.Flour)
	}

}
