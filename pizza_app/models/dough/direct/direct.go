package direct

type MainDough struct {
	Flour           string `json:"flour"`
	Water           string `json:"water"`
	InstantDryYeast string `json:"instantDryYeast"`
	Salt            string `json:"salt"`
}

type DirectDoughResponse struct {
	MainDough       MainDough `json:"mainDough"`
	DoughBallAmount string    `json:"doughBallAmount"`
	Hydration       string    `json:"hydration"`
	DoughBallWeight string    `json:"doughBallWeight"`
}
