package direct

import (
	"testing"
)

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

	if mainDough.Water != 80.0 {
		t.Errorf("Expected water to be 80.0, but got %f", mainDough.Water)
	}

	if mainDough.InstantDryYeast != 0.50 {
		t.Errorf("Expected instantDryYeast to be 0.50, but got %f", mainDough.InstantDryYeast)
	}

	if mainDough.Salt != 1.0 {
		t.Errorf("Expected salt to be 1.0, but got %f", mainDough.Salt)
	}
}

func TestDirectDoughResponse(t *testing.T) {
	directResponse := DirectDoughResponse{
		MainDough: MainDough{
			Flour:           120.0,
			Water:           80.0,
			InstantDryYeast: 0.50,
			Salt:            1.0,
		},
		DoughBallAmount: 4,
		Hydration:       70.0,
		DoughBallWeight: 270.0,
	}

	if directResponse.MainDough.Flour != 120.0 {
		t.Errorf("Expected MainDough flour to be 120.0, but got %f", directResponse.MainDough.Flour)
	}

}
