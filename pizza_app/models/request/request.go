package models

type DoughRequest struct {
	DoughBallAmount int     `json:"doughBallAmount"` // In number of dough balls
	Hydration       float64 `json:"hydration"`       // In percentage from 0 to 100
	DoughBallWeight float64 `json:"doughBallWeight"` // In grams
}

type BigaDoughRequest struct {
	DoughBallAmount int     `json:"doughBallAmount"`         // In number of dough balls
	Hydration       float64 `json:"hydration"`               // In percentage from 0 to 100
	DoughBallWeight float64 `json:"doughBallWeight"`         // In grams
	FlourPercentage float64 `json:"percentageBiga"`          // In percentage from 0 to 100
	BigaHydration   float64 `json:"percentageBigaHydration"` // In percentage from 0 to 100
}

type PoolishDoughRequest struct {
	DoughBallWeight   float64 `json:"doughBallWeight"`
	DoughBallAmount   int     `json:"doughBallAmount"`
	Hydration         float64 `json:"hydration"`
	PoolishPercentage float64 `json:"poolishPercentage"`
}
