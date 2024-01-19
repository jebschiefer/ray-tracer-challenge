package utilities

import (
	"math"
)

const EPSILON = 1e-4

func FloatsEqual(a float64, b float64) bool {
	return math.Abs(a-b) < EPSILON
}
