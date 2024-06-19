package pide

type MainDough struct {
	Flour           float64 `json:"flour"`
	Water           float64 `json:"water"`
	InstantDryYeast float64 `json:"instantDryYeast"`
	FreshYeast      float64 `json:"instantFreshYeast"`
	Salt            float64 `json:"salt"`
}

type PideDoughResponse struct {
	MainDough       MainDough `json:"mainDough"`
	DoughBallAmount int       `json:"doughBallAmount"`
	Hydration       float64   `json:"hydration"`
	DoughBallWeight float64   `json:"doughBallWeight"`
}
