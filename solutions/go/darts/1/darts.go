package darts

import "math"

func Score(x, y float64) int {

	radius := math.Sqrt((x * x) + (y * y))

	if radius > 5 && radius <= 10 {
		return 1
	} else if radius > 1 && radius <= 5 {
		return 5
	} else if radius >= 0 && radius <= 1 {
		return 10
	}

	return 0
}
