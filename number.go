package physics

import (
	"math"
)

const (
	PRECISION = 1e-6
)

type Nubmber = float64

func almostZero(n Nubmber) (bool) {
	if math.Abs(float64(n)) > PRECISION {
		return false
	}
	return true
}

func almostEquals(n, n1 Nubmber) (bool) {
	if math.Abs(float64(n - n1)) > PRECISION {
		return false
	}
	return true
}


