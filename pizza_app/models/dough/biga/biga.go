package biga

type BigaDough struct {
	Flour           string `json:"flour"`
	Water           string `json:"water"`
	InstantDryYeast string `json:"instantDryYeast"`
}

type MainDough struct {
	Flour string `json:"flour"`
	Water string `json:"water"`
	Salt  string `json:"salt"`
}

type BigaDoughResponse struct {
	Biga                BigaDough `json:"biga"`
	MainDough           MainDough `json:"mainDough"`
	DoughBallAmount     string    `json:"doughBallAmount"`
	Hydration           string    `json:"hydration"`
	BigaHydration       string    `json:"bigaHydration"`
	BigaWaterPercentage string    `json:"bigaWaterPercentage"`
	DoughBallWeight     string    `json:"doughBallWeight"`
}
