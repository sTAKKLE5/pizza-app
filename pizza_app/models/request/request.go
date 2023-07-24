package models

type DoughRequest struct {
	DoughBallAmount int     `json:"doughBallAmount"` // In number of dough balls
	Hydration       float64 `json:"hydration"`       // In percentage from 0 to 100
	DoughBallWeight float64 `json:"doughBallWeight"` // In grams
}

type PoolishDoughRequest struct {
	DoughBallWeight   float64 `json:"doughBallWeight"`
	DoughBallAmount   int     `json:"doughBallAmount"`
	Hydration         float64 `json:"hydration"`
	PoolishPercentage float64 `json:"poolishPercentage"`
}
