package biga

type BigaDough struct {
	Flour           float64 `json:"flour"`
	Water           float64 `json:"water"`
	InstantDryYeast float64 `json:"instantDryYeast"`
	FreshYeast      float64 `json:"freshYeast"`
}

type MainDough struct {
	Flour float64 `json:"flour"`
	Water float64 `json:"water"`
	Salt  float64 `json:"salt"`
}

type BigaDoughResponse struct {
	Biga                BigaDough `json:"biga"`
	MainDough           MainDough `json:"mainDough"`
	DoughBallAmount     int       `json:"doughBallAmount"`
	Hydration           float64   `json:"hydration"`
	BigaHydration       float64   `json:"bigaHydration"`
	BigaWaterPercentage float64   `json:"bigaWaterPercentage"`
	DoughBallWeight     float64   `json:"doughBallWeight"`
}
