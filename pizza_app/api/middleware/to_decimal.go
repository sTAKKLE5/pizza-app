package utils

import "math"

func RoundToDecimal(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Round(f*shift) / shift
}
