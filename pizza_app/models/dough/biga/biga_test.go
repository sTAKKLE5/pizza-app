package biga

import (
	"testing"
)

func TestBigaDough(t *testing.T) {
	biga := BigaDough{
		Flour:           100.0,
		Water:           60.0,
		InstantDryYeast: 0.50,
	}

	if biga.Flour != 100.0 {
		t.Errorf("Expected flour to be 100.0, but got %f", biga.Flour)
	}

	if biga.Water != 60.0 {
		t.Errorf("Expected water to be 60.0, but got %f", biga.Water)
	}

	if biga.InstantDryYeast != 0.50 {
		t.Errorf("Expected instantDryYeast to be 0.50, but got %f", biga.InstantDryYeast)
	}
}

func TestMainDough(t *testing.T) {
	mainDough := MainDough{
		Flour: 120.0,
		Water: 80.0,
		Salt:  1.0,
	}

	if mainDough.Flour != 120.0 {
		t.Errorf("Expected flour to be 120.0, but got %f", mainDough.Flour)
	}

	if mainDough.Water != 80.0 {
		t.Errorf("Expected water to be 80.0, but got %f", mainDough.Water)
	}

	if mainDough.Salt != 1.0 {
		t.Errorf("Expected salt to be 1.0, but got %f", mainDough.Salt)
	}
}

func TestBigaDoughResponse(t *testing.T) {
	bigaResponse := BigaDoughResponse{
		Biga: BigaDough{
			Flour:           100.0,
			Water:           60.0,
			InstantDryYeast: 0.50,
		},
		MainDough: MainDough{
			Flour: 120.0,
			Water: 80.0,
			Salt:  1.0,
		},
		DoughBallAmount:     4,
		Hydration:           70.0,
		BigaHydration:       60.0,
		BigaWaterPercentage: 40.0,
		DoughBallWeight:     270.0,
	}

	if bigaResponse.Biga.Flour != 100.0 {
		t.Errorf("Expected Biga flour to be 100.0, but got %f", bigaResponse.Biga.Flour)
	}

}
