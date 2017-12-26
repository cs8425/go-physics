package physics

import (
	"math"
)

const (
	PRECISION = 1e-6
)

type Nubmber = float64

func (n *Nubmber) AlmostZero() (bool) {
	if math.Abs(n) > PRECISION {
		return false
	}
	return true
}

func (n *Nubmber) AlmostEquals(n1 *Nubmber) (bool) {
	if math.Abs(n - n1) > PRECISION {
		return false
	}
	return true
}


