package models

type DoughRequest struct {
	DoughBallAmount int     `json:"doughBallAmount"`
	Hydration       float64 `json:"hydration"`
	DoughBallWeight float64 `json:"doughBallWeight"`
}
