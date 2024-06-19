package pide

type MainDough struct {
	Flour      float64 `json:"flour"`
	Water      float64 `json:"water"`
	FreshYeast float64 `json:"instantFreshYeast"`
	Salt       float64 `json:"salt"`
	OliveOil   float64 `json:"oliveOil"`
}

type PideDoughResponse struct {
	MainDough       MainDough `json:"mainDough"`
	DoughBallAmount int       `json:"doughBallAmount"`
	DoughBallWeight float64   `json:"doughBallWeight"`
}
