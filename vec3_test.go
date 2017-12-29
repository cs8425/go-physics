package physics

import (
	"testing"
)

func TestVec3Equals(t *testing.T) {

	var v = NewVec3().Set(1, 2, 3)
	var u = NewVec3().Set(4, 5, 6)

	if v.IsEquals(u) {
		t.Error("should not equal:", v, u)
	}
	if v.AlmostEquals(u) {
		t.Error("should not almost equal:", v, u)
	}

	u.Set(1 + 0.5*PRECISION, 2 + 0.5*PRECISION, 3 + 0.5*PRECISION)
	if v.IsEquals(u) {
		t.Error("should not equal:", v, u)
	}
	if !v.AlmostEquals(u) {
		t.Error("should almost equal:", v, u)
	}

}

func TestVec3Cross(t *testing.T) {

	var vok = NewVec3().Set(-3, 6, -3)

	var v = NewVec3().Set(1, 2, 3)
	var u = NewVec3().Set(4, 5, 6)
	v = v.Cross(u, nil)

	if !v.IsEquals(vok) {
		t.Error("Error Calculating cross product, got ", v, vok)
	}

	v.Set(1, 2, 3)
	v.Cross(u, v)

	if !v.IsEquals(vok) {
		t.Error("Error Calculating cross product, got ", v, vok)
	}
}

func TestVec3Dot(t *testing.T) {

	var v = NewVec3().Set(1, 2, 3)
	var u = NewVec3().Set(4, 5, 6)
	var dot = v.Dot(u)

	if dot != (4 + 10 + 18) {
		t.Error("Error Calculating dot product, got ", dot)
	}


	v.Set(3, 2, 1)
	u.Set(4, 5, 6)
	dot = v.Dot(u)

	if dot != (12 + 10 + 6) {
		t.Error("Error Calculating dot product, got ", dot)
	}
}

func TestVec3VAdd(t *testing.T) {

	var vok = NewVec3().Set(5, 7, 9)

	var v = NewVec3().Set(1, 2, 3)
	var u = NewVec3().Set(4, 5, 6)
	v = v.VAdd(u, nil)

	if !v.IsEquals(vok) {
		t.Error("Error Calculating VAdd, got ", v, vok)
	}


	v.Set(1, 2, 3)
	v.VAdd(u, v)

	if !v.IsEquals(vok) {
		t.Error("Error Calculating VAdd, got ", v, vok)
	}

}

func TestVec3IsAntiparallelTo(t *testing.T) {

	var v = NewVec3().Set(1, 0, 0)
	var u = NewVec3().Set(-1, 0, 0)

	if !v.IsAntiparallelTo(u) {
		t.Error("Error IsAntiparallelTo, got ", v.IsAntiparallelTo(u))
	}

}


