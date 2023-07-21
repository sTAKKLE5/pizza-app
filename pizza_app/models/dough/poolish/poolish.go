package poolish

type PoolishDough struct {
	Flour           float64 `json:"flour"`
	Water           float64 `json:"water"`
	InstantDryYeast float64 `json:"instantDryYeast"`
}

type MainDough struct {
	Flour           float64 `json:"flour"`
	Water           float64 `json:"water"`
	Salt            float64 `json:"salt"`
	InstantDryYeast float64 `json:"instantDryYeast"`
}

type PoolishDoughResponse struct {
	Poolish         PoolishDough `json:"poolish"`
	MainDough       MainDough    `json:"mainDough"`
	DoughBallAmount int          `json:"doughBallAmount"`
	Hydration       float64      `json:"hydration"`
	DoughBallWeight float64      `json:"doughBallWeight"`
}
