package direct

type MainDough struct {
	Flour           float64 `json:"flour"`
	Water           float64 `json:"water"`
	InstantDryYeast float64 `json:"instantDryYeast"`
	Salt            float64 `json:"salt"`
}

type DirectDoughResponse struct {
	MainDough       MainDough `json:"mainDough"`
	DoughBallAmount int       `json:"doughBallAmount"`
	Hydration       float64   `json:"hydration"`
	DoughBallWeight float64   `json:"doughBallWeight"`
}
