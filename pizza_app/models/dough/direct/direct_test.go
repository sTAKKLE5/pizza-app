package direct

import (
	"testing"
)

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

	if mainDough.Water != "80.00" {
		t.Errorf("Expected water to be '80.00', but got '%s'", mainDough.Water)
	}

	if mainDough.InstantDryYeast != "0.50" {
		t.Errorf("Expected instantDryYeast to be '0.50', but got '%s'", mainDough.InstantDryYeast)
	}

	if mainDough.Salt != "1.00" {
		t.Errorf("Expected salt to be '1.00', but got '%s'", mainDough.Salt)
	}
}

func TestDirectDoughResponse(t *testing.T) {
	directResponse := DirectDoughResponse{
		MainDough: MainDough{
			Flour:           "120.00",
			Water:           "80.00",
			InstantDryYeast: "0.50",
			Salt:            "1.00",
		},
		DoughBallAmount: "4",
		Hydration:       "70.00",
		DoughBallWeight: "270.00",
	}

	if directResponse.MainDough.Flour != "120.00" {
		t.Errorf("Expected MainDough flour to be '120.00', but got '%s'", directResponse.MainDough.Flour)
	}

}
