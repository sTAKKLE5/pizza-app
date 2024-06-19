package models

type DoughRequest struct {
	DoughBallAmount int     `json:"doughBallAmount"` // In number of dough balls
	Hydration       float64 `json:"hydration"`       // In percentage from 0 to 100
	DoughBallWeight float64 `json:"doughBallWeight"` // In grams
}

type DoughRequestPide struct {
	DoughBallAmount int `json:"doughBallAmount"` // In number of dough balls
}

type BigaDoughRequest struct {
	DoughBallAmount         int     `json:"doughBallAmount"`         // In number of dough balls
	Hydration               float64 `json:"hydration"`               // In percentage from 0 to 100
	DoughBallWeight         float64 `json:"doughBallWeight"`         // In grams
	FlourPercentage         float64 `json:"percentageBiga"`          // In percentage from 0 to 100
	BigaHydration           float64 `json:"percentageBigaHydration"` // In percentage from 0 to 100
	FermentationHours       int     `json:"fermentationHours"`       // In hours
	FermentationTemperature int     `json:"fermentationTemperature"` // In Celsius
}

type PoolishDoughRequest struct {
	DoughBallWeight   float64 `json:"doughBallWeight"`
	DoughBallAmount   int     `json:"doughBallAmount"`
	Hydration         float64 `json:"hydration"`
	PoolishPercentage float64 `json:"poolishPercentage"`
}
