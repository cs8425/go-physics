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

func clamp(value, min, max Number) (Number) {
	val := math.Max( float64(min), math.Min( float64(max), float64(value) ) )
	return Number(val)
}

func clampF(value Number, min float64, max float64) (float64) {
	val := math.Max( min, math.Min( max, float64(value) ) )
	return val
}

