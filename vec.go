package physics

import (
	"math"
)

type Vec3 [3]Nubmber // order: x, y, z

/*func NewVec3(x, y, z Nubmber) (*Vec3) {
	v := &Vec3{ x, y, z }
	return v
}*/

func NewVec3() (*Vec3) {
	v := &Vec3{ 0, 0, 0 }
	return v
}

/**
 * Set the vectors' 3 elements
 * @method set
 * @param {Number} x
 * @param {Number} y
 * @param {Number} z
 * @return Vec3
 */
func (v *Vec3) Set(x, y, z Nubmber) (*Vec3) {
	v[0], v[1], v[2] = x, y, z
	return v
}

/**
 * Clone the vector
 * @method clone
 * @return {Vec3}
 */
func (v *Vec3) Clone() (*Vec3) {
	return &Vec3{ v[0], v[1], v[2] }
}

/**
 * Copies value of source to this vector.
 * @method copy
 * @param {Vec3} source
 * @return {Vec3} this
 */
func (v *Vec3) Copy(source *Vec3) (*Vec3) {
	v[0], v[1], v[2] = source[0], source[1], source[2]
	return v
};


/**
 * @method IsZero
 * @return bool
 */
func (v *Vec3) IsZero() (bool) {
	return v[0] == 0 && v[1] == 0 && v[2] == 0
}

/**
 * Check if the vector is anti-parallel to another vector.
 * @method isAntiparallelTo
 * @param  {Vec3}  v
 * @param  {Number}  precision Set to zero for exact comparisons
 * @return {Boolean}
 */
func (v *Vec3) IsAntiparallelTo(v1 *Vec3) (bool) {
	antip_neg := NewVec3()
	v.Negate(antip_neg)
	return antip_neg.AlmostEquals(v1)
}

/**
 * Check if a vector equals is almost equal to another one.
 * @method almostEquals
 * @param {Vec3} v
 * @return bool
 */
func (v *Vec3) AlmostEquals(v1 *Vec3) (bool) {
	if v[0].AlmostEquals(v1[0]) && v[1].AlmostEquals(v1[1]) && v[2].AlmostEquals(v1[2]) {
		return true
	}
	return false
}

/**
 * Check if a vector is almost zero
 * @method almostZero
 * @param {Number} precision
 */
func (v *Vec3) AlmostZero() (bool) {
	if v[0].AlmostZero() && v[1].AlmostZero() && v[2].AlmostZero() {
		return true
	}
	return false
}


/**
 * Calculate dot product
 * @method Dot
 * @param {Vec3} v
 * @return {Number}
 */
func (v *Vec3) Dot(v1 *Vec3) (Nubmber) {
	return v[0] * v1[0] + v[1] * v1[1] + v[2] * v1[2]
}

/**
 * Get the squared length of the vector.
 * @method LengthSquared
 * @return {Number}
 */
func (v *Vec3) LengthSquared(v1 *Vec3) (Nubmber) {
	return v.Dot(v)
}


/**
 * Make the vector point in the opposite direction.
 * @method Negate
 * @param {Vec3} target Optional target to save in
 * @return {Vec3}
 */
func (v *Vec3) Negate(target *Vec3) (*Vec3) {
	if target == nil {
		target = &Vec3{}
	}
	target[0] = -v[0]
	target[1] = -v[1]
	target[2] = -v[2]

	return target
}


/**
 * Vector addition / subtraction
 * @method VAdd
 * @param {Vec3} v
 * @param {Vec3} target Optional.
 * @return {Vec3}
 */
func (v *Vec3) VAdd(v1 *Vec3, target *Vec3) (*Vec3) {
	if target == nil {
		target = &Vec3{}
	}

	target[0] = v[0] + v1[0]
	target[1] = v[1] + v1[1]
	target[2] = v[2] + v1[2]

	return target
}

/**
 * Vector cross product
 * @method Cross
 * @param {Vec3} v
 * @param {Vec3} target Optional. Target to save in.
 * @return {Vec3}
 */
func (v *Vec3) Cross(v1 *Vec3, target *Vec3) (*Vec3) {
	vx, vy, vz := v1[0], v1[1], v1[2]
	x, y, z := v[0], v[1], v[2]

	if target == nil {
		target = &Vec3{}
	}

	target[0] = (y * vz) - (z * vy)
	target[1] = (z * vx) - (x * vz)
	target[2] = (x * vy) - (y * vx)

	return target
}


/**
 * Get the length of the vector
 * @method Norm
 * @return {Number}
 * @deprecated Use .length() instead
 */
func (v *Vec3) Norm() (Nubmber) {
	x, y, z := v[0], v[1], v[2]
	n := math.Sqrt(float64(x*x + y*y + z*z))

	return Nubmber(n)
}

/**
 * Get the length of the vector
 * @method Length
 * @return {Number}
 */
func (v *Vec3) Length() (Nubmber) {
	return v.Norm()
}

/**
 * Normalize the vector. Note that this changes the values in the vector.
 * @method Normalize
 * @return {Number} Returns the norm of the vector
 */
func (v *Vec3) Normalize() (Nubmber) {

	n := v.Norm()
	if n > 0.0 {
		invN := 1 / n
		v[0] *= invN
		v[1] *= invN
		v[2] *= invN
	} else {
		// Make something up
		v[0], v[1], v[2] = 0, 0, 0
	}

	return n
}


/**
 * Get the version of this vector that is of length 1.
 * @method Unit
 * @param {Vec3} target Optional target to save in
 * @return {Vec3} Returns the unit vector
 */
func (v *Vec3) Unit(target *Vec3) (*Vec3) {
	if target == nil {
		target = &Vec3{1, 0, 0}
	}

	n := v.Norm()
	if n > 0.0 {
		invN := 1 / n
		target[0] = v[0] * invN
		target[1] = v[1] * invN
		target[2] = v[2] * invN
	}

	return target
}


/**
 * Get squared distance from this point to another point
 * @method DistanceSquared
 * @param  {Vec3} p
 * @return {Number}
 */
func (v *Vec3) DistanceSquared(p *Vec3) (Nubmber) {
	x, y, z := v[0], v[1], v[2]
	px, py, pz := p[0], p[1], p[2]

	return (px-x)*(px-x) + (py-y)*(py-y) + (pz-z)*(pz-z)
}

/**
 * Get distance from this point to another point
 * @method DistanceTo
 * @param  {Vec3} p
 * @return {Number}
 */
func (v *Vec3) DistanceTo(p *Vec3) (Nubmber) {
	diSq := v.DistanceSquared(p)
	return Nubmber(math.Sqrt(diSq))
}


/**
 * Multiply all the components of the vector with a scalar.
 * @method Scale
 * @param {Number} scalar
 * @param {Vec3} target The vector to save the result in.
 * @return {Vec3}
 */
func (v *Vec3) Scale(scalar Nubmber, target *Vec3) (*Vec3) {
	if target == nil {
		target = &Vec3{}
	}

	target[0] = v[0] * scalar
	target[1] = v[1] * scalar
	target[2] = v[2] * scalar

	return target
}


/**
 * Multiply the vector with an other vector, component-wise.
 * @method mult
 * @param {Number} vector
 * @param {Vec3} target The vector to save the result in.
 * @return {Vec3}
 */
func (v *Vec3) VMul(vector *Vec3, target *Vec3) (*Vec3) {
	if target == nil {
		target = &Vec3{}
	}

	target[0] = vector[0] * v[0]
	target[1] = vector[1] * v[1]
	target[2] = vector[2] * v[2]

	return target
}

/**
 * Scale a vector and add it to this vector. Save the result in "target". (target = this + vector * scalar)
 * @method AddScaledVector
 * @param {Number} scalar
 * @param {Vec3} vector
 * @param {Vec3} target The vector to save the result in.
 * @return {Vec3}
 */
func (v *Vec3) AddScaledVector(scalar Nubmber, vector *Vec3, target *Vec3) (*Vec3) {
	if target == nil {
		target = &Vec3{}
	}

	// target = this + scalar * vector
	scaled := vector.Scale(scalar, target)
	v.VAdd(scaled, target)

	return target
}

/**
 * Do a linear interpolation between two vectors
 * @method lerp
 * @param {Vec3} v
 * @param {Number} t A number between 0 and 1. 0 will make this function return u, and 1 will make it return v. Numbers in between will generate a vector in between them.
 * @param {Vec3} target
 */
func (v *Vec3) Lerp(v1 *Vec3, t Nubmber, target *Vec3) (*Vec3) {
	if target == nil {
		target = &Vec3{}
	}

	x, y, z := v[0], v[1], v[2]
	target[0] = x + (v1[0]-x)*t
	target[1] = y + (v1[1]-y)*t
	target[2] = z + (v1[2]-z)*t

	return target
}

/**
 * Compute two artificial tangents to the vector
 * @method tangents
 * @param {Vec3} t1 Vector object to save the first tangent in
 * @param {Vec3} t2 Vector object to save the second tangent in
 */
func (v *Vec3) Tangents(t1 *Vec3, t2 *Vec3) {
	norm := v.Norm()
	if norm > 0.0 {
		inorm := 1 / norm
		n := v.Scale(inorm, nil)
		randVec := NewVec3()
		if(math.Abs(n[0]) < 0.9){
			randVec.Set(1, 0, 0)
		} else {
			randVec.Set(0, 1, 0)
		}
		n.Cross(randVec, t1)
		n.Cross(t1, t2)
	} else {
		// The normal length is zero, make something up
		t1.Set(1, 0, 0)
		t2.Set(0, 1, 0)
	}
}


