package sourdough

type MainDough struct {
	Flour     float64 `json:"flour"`
	Water     float64 `json:"water"`
	SourDough float64 `json:"sourDough"`
	Salt      float64 `json:"salt"`
}

type SourDoughResponse struct {
	MainDough       MainDough `json:"mainDough"`
	DoughBallAmount int       `json:"doughBallAmount"`
	Hydration       float64   `json:"hydration"`
	DoughBallWeight float64   `json:"doughBallWeight"`
}
