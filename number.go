package physics

import (
	"math"
)

const (
	PRECISION = 1e-6
)

type Number = float64

func almostZero(n Number) (bool) {
	if math.Abs(float64(n)) > PRECISION {
		return false
	}
	return true
}

func almostEquals(n, n1 Number) (bool) {
	if math.Abs(float64(n - n1)) > PRECISION {
		return false
	}
	return true
}


