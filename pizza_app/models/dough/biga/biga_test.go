package biga

import (
	"testing"
)

func TestBigaDough(t *testing.T) {
	biga := BigaDough{
		Flour:           "100.00",
		Water:           "60.00",
		InstantDryYeast: "0.50",
	}

	if biga.Flour != "100.00" {
		t.Errorf("Expected flour to be '100.00', but got '%s'", biga.Flour)
	}

	if biga.Water != "60.00" {
		t.Errorf("Expected water to be '60.00', but got '%s'", biga.Water)
	}

	if biga.InstantDryYeast != "0.50" {
		t.Errorf("Expected instantDryYeast to be '0.50', but got '%s'", biga.InstantDryYeast)
	}
}

func TestMainDough(t *testing.T) {
	mainDough := MainDough{
		Flour: "120.00",
		Water: "80.00",
		Salt:  "1.00",
	}

	if mainDough.Flour != "120.00" {
		t.Errorf("Expected flour to be '120.00', but got '%s'", mainDough.Flour)
	}

	if mainDough.Water != "80.00" {
		t.Errorf("Expected water to be '80.00', but got '%s'", mainDough.Water)
	}

	if mainDough.Salt != "1.00" {
		t.Errorf("Expected salt to be '1.00', but got '%s'", mainDough.Salt)
	}
}

func TestBigaDoughResponse(t *testing.T) {
	bigaResponse := BigaDoughResponse{
		Biga: BigaDough{
			Flour:           "100.00",
			Water:           "60.00",
			InstantDryYeast: "0.50",
		},
		MainDough: MainDough{
			Flour: "120.00",
			Water: "80.00",
			Salt:  "1.00",
		},
		DoughBallAmount:     "4",
		Hydration:           "70.00",
		BigaHydration:       "60.00",
		BigaWaterPercentage: "40.00",
		DoughBallWeight:     "270.00",
	}

	if bigaResponse.Biga.Flour != "100.00" {
		t.Errorf("Expected Biga flour to be '100.00', but got '%s'", bigaResponse.Biga.Flour)
	}

}
