package poolish

type PoolishDough struct {
	Flour           string `json:"flour"`
	Water           string `json:"water"`
	InstantDryYeast string `json:"instantDryYeast"`
}

type MainDough struct {
	Flour           string `json:"flour"`
	Water           string `json:"water"`
	Salt            string `json:"salt"`
	InstantDryYeast string `json:"instantDryYeast"`
}

type PoolishDoughResponse struct {
	Poolish                PoolishDough `json:"poolish"`
	MainDough              MainDough    `json:"mainDough"`
	DoughBallAmount        string       `json:"doughBallAmount"`
	Hydration              string       `json:"hydration"`
	DoughBallWeight        string       `json:"doughBallWeight"`
	PoolishHydration       string       `json:"poolishHydration"`
	PoolishWaterPercentage string       `json:"poolishWaterPercentage"`
	PoolishYeastPercentage string       `json:"poolishYeastPercentage"`
}
