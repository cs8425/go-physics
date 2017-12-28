package physics

import (
	"testing"
	"math"
)

func TestQuatEquals(t *testing.T) {

	var v = NewQuat().Set(1, 2, 3, 4)
	var u = NewQuat().Set(4, 5, 6, 7)

	if v.IsEquals(u) {
		t.Error("should not equal:", v, u)
	}
	if v.AlmostEquals(u) {
		t.Error("should not almost equal:", v, u)
	}

	u.Set(1 + 0.5*PRECISION, 2 + 0.5*PRECISION, 3 + 0.5*PRECISION, 4 + 0.5*PRECISION)
	if v.IsEquals(u) {
		t.Error("should not equal:", v, u)
	}
	if !v.AlmostEquals(u) {
		t.Error("should almost equal:", v, u)
	}

}


func TestQuatConjugate(t *testing.T) {

	var qok = NewQuat().Set(-1, -2, -3, 4)

	var q = NewQuat().Set(1, 2, 3, 4)
	q = q.Conjugate(nil)

	if !q.IsEquals(qok) {
		t.Error("Error Calculating Conjugate, got ", q, qok)
	}


	q.Set(1, 2, 3, 4)
	q.Conjugate(q)

	if !q.IsEquals(qok) {
		t.Error("Error Calculating Conjugate, got ", q, qok)
	}

}

func TestQuatInverse(t *testing.T) {

	var denominator Number = 1*1 + 2*2 + 3*3 + 4*4
	var qok = NewQuat().Set(-1/denominator, -2/denominator, -3/denominator, 4/denominator)

	var q = NewQuat().Set(1, 2, 3, 4)
	q = q.Inverse(nil)

	if !q.IsEquals(qok) {
		t.Error("Error Calculating Inverse, got ", q, qok)
	}


	q.Set(1, 2, 3, 4)
	q.Inverse(q)

	if !q.IsEquals(qok) {
		t.Error("Error Calculating Inverse, got ", q, qok)
	}

}

func TestQuatSlerp(t *testing.T) {

	var qa = NewQuat()
	var qb = NewQuat()
	qa.Slerp(qb, 0.5, qb)

	if !qa.IsEquals(qb) {
		t.Error("Error Calculating Slerp, got ", qa, qb)
	}

	var qok = NewQuat()
	var axis = NewVec3().Set(0, 0, 1)
	qa.SetFromAxisAngle(axis, math.Pi / 4)
	qb.SetFromAxisAngle(axis, -math.Pi / 4)

	qa.Slerp(qb, 0.5, qb)

	if !qb.IsEquals(qok) {
		t.Error("Error Calculating Slerp, got ", qb, qok, qa)
	}

}

func TestQuatSlerp2(t *testing.T) {

	var qa = NewQuat()
	var qb = NewQuat()
	qb = qa.Slerp(qb, 0.5, nil)

	if !qa.IsEquals(qb) {
		t.Error("Error Calculating Slerp, got ", qa, qb)
	}

	var qok = NewQuat()
	var axis = NewVec3().Set(0, 0, 1)
	qa.SetFromAxisAngle(axis, math.Pi / 4)
	qb.SetFromAxisAngle(axis, -math.Pi / 4)

	qb = qa.Slerp(qb, 0.5, nil)

	if !qb.IsEquals(qok) {
		t.Error("Error Calculating Slerp, got ", qb, qok, qa)
	}

}

func TestQuatSetFromVectors(t *testing.T) {

	var va = NewVec3().Set(1, 0, 0)
	var vb = NewVec3().Set(-1, 0, 0)
	var q = NewQuat()
	q.SetFromVectors(va, vb)

	ret := q.VMult(va, nil)
	if !ret.AlmostEquals(vb) {
		t.Error("Error Calculating SetFromVectors, got ", va, vb)
	}


	va = NewVec3().Set(0, 1, 0)
	vb = NewVec3().Set(0, -1, 0)
	q.SetFromVectors(va, vb)

	ret = q.VMult(va, nil)
	if !ret.AlmostEquals(vb) {
		t.Error("Error Calculating SetFromVectors, got ", va, vb)
	}

	va = NewVec3().Set(0, 0, 1)
	vb = NewVec3().Set(0, 0, -1)
	q.SetFromVectors(va, vb)

	ret = q.VMult(va, nil)
	if !ret.AlmostEquals(vb) {
		t.Error("Error Calculating SetFromVectors, got ", va, vb)
	}

}

func TestQuatToEuler(t *testing.T) {

	axis := NewVec3().Set(0, 0, 1)
	q := NewQuat().SetFromAxisAngle(axis, math.Pi / 4)

	euler := q.ToEuler(nil, YZX)

	// we should expect (0,0,pi/4)
	vok := NewVec3().Set(0, 0, math.Pi / 4)
	if !vok.IsEquals(euler) {
		t.Error("Error Calculating ToEuler, got ", euler, vok, q)
	}
}


