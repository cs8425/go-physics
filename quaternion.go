package physics

import (
	"math"
)

type AxisOrder int

const (
	XYZ AxisOrder = iota // 0
	YXZ // 1
	ZXY // 2
	ZYX // 3
	YZX // 4
	XZY // 5
)

type Quat [4]Number // order: x, y, z, w

/*func NewQuat(x, y, z, w Number) (*Quat) {
	q := &Quat{ x, y, z, w }
	return q
}*/

func NewQuat() (*Quat) {
	q := &Quat{ 0, 0, 0, 1 }
	return q
}


/**
 * Set the value of the quaternion.
 * @method set
 * @param {Number} x
 * @param {Number} y
 * @param {Number} z
 * @param {Number} w
 */
func (q *Quat) Set(x, y, z, w Number) (*Quat) {
	q[0], q[1], q[2], q[3] = x, y, z, w
	return q
}

/**
 * @method clone
 * @return {Quaternion}
 */
func (q *Quat) Clone() (*Quat) {
	return &Quat{ q[0], q[1], q[2], q[3] }
}

/**
 * Copies value of source to this quaternion.
 * @method copy
 * @param {Quaternion} source
 * @return {Quaternion} this
 */
func (q *Quat) Copy(source *Quat) (*Quat) {
	q[0], q[1], q[2], q[3] = source[0], source[1], source[2], source[3]
	return q
}


/**
 * Set the quaternion components given an axis and an angle.
 * @method setFromAxisAngle
 * @param {Vec3} axis
 * @param {Number} angle in radians
 */
func (q *Quat) SetFromAxisAngle(axis *Vec3, angle Number) (*Quat) {
	sin, c := math.Sincos(float64(angle) * 0.5)
	s := Number(sin)
	q[0] = axis[0] * s
	q[1] = axis[1] * s
	q[2] = axis[2] * s
	q[3] = Number(c)
	return q
}

/**
 * Converts the quaternion to axis/angle representation.
 * @method toAxisAngle
 * @param {Vec3} [targetAxis] A vector object to reuse for storing the axis.
 * @return {Vec3, Number} first elemnt is the axis and the second is the angle in radians.
 */
func (q *Quat) ToAxisAngle(targetAxis *Vec3) (*Vec3, Number) {
	if targetAxis == nil {
		targetAxis = &Vec3{}
	}

	q.Normalize() // if w>1 acos and sqrt will produce errors, this cant happen if quaternion is normalised

	w := float64(q[3])
	angle := Number(2 * math.Acos(w))
	s := Number(math.Sqrt(1 - w*w)) // assuming quaternion normalised then w is less than 1, so term always positive.

	if (s < 0.001) { // test to avoid divide by zero, s is always positive due to sqrt
		// if s close to zero then direction of axis not important
		targetAxis[0] = q[0] // if it is important that axis is normalised then replace with x=1; y=z=0;
		targetAxis[1] = q[1]
		targetAxis[2] = q[2]
	} else {
		targetAxis[0] = q[0] / s // normalise axis
		targetAxis[1] = q[1] / s
		targetAxis[2] = q[2] / s
	}

	return targetAxis, angle
}

/**
 * Set the quaternion value given two vectors. The resulting rotation will be the needed rotation to rotate u to v.
 * @method setFromVectors
 * @param {Vec3} u
 * @param {Vec3} v
 */
func (q *Quat) SetFromVectors(u *Vec3, v *Vec3) (*Quat) {
	if u.IsAntiparallelTo(v) {
		t1 := NewVec3()
		t2 := NewVec3()

		u.Tangents(t1, t2)
		q.SetFromAxisAngle(t1, math.Pi)
	} else {
		a := u.Cross(v, nil)
		q[0] = a[0]
		q[1] = a[1]
		q[2] = a[2]
		//un := u.Norm()
		//vn := v.Norm()
		//q[3] = Number(math.Sqrt(float64( un*un * vn*vn ))) + u.Dot(v)
		q[3] = Number(math.Sqrt(float64( u.LengthSquared() * v.LengthSquared() ))) + u.Dot(v)
		q.Normalize()
	}
	return q
}

/**
 * Quaternion multiplication
 * @method mult
 * @param {Quaternion} q
 * @param {Quaternion} target Optional.
 * @return {Quaternion}
 */
func (q *Quat) Mult(q1 *Quat, target *Quat) (*Quat) {
	if target == nil {
		target = NewQuat()
	}

	ax, ay, az, aw := q[0], q[1], q[2], q[3]
	bx, by, bz, bw := q1[0], q1[1], q1[2], q1[3]

	target[0] = ax * bw + aw * bx + ay * bz - az * by
	target[1] = ay * bw + aw * by + az * bx - ax * bz
	target[2] = az * bw + aw * bz + ax * by - ay * bx
	target[3] = aw * bw - ax * bx - ay * by - az * bz

	return target
}

/**
 * Get the inverse quaternion rotation.
 * @method inverse
 * @param {Quaternion} target
 * @return {Quaternion}
 */
func (q *Quat) Inverse(target *Quat) (*Quat) {
	if target == nil {
		target = NewQuat()
	}

	x, y, z, w := q[0], q[1], q[2], q[3]

	q.Conjugate(target)
	inorm2 := 1 / (x*x + y*y + z*z + w*w)

	target[0] *= inorm2
	target[1] *= inorm2
	target[2] *= inorm2
	target[3] *= inorm2

	return target
}

/**
 * Get the quaternion conjugate
 * @method conjugate
 * @param {Quaternion} target
 * @return {Quaternion}
 */
func (q *Quat) Conjugate(target *Quat) (*Quat) {
	if target == nil {
		target = &Quat{}
	}

	target[0] = -q[0]
	target[1] = -q[1]
	target[2] = -q[2]
	target[3] = q[3]

	return target
}

/**
 * Normalize the quaternion. Note that this changes the values of the quaternion.
 * @method normalize
 */
func (q *Quat) Normalize() (*Quat) {
	x, y, z, w := q[0], q[1], q[2], q[3]
	l := math.Sqrt(float64(x*x + y*y + z*z + w*w))

	if l == 0 {
		q[0], q[1], q[2], q[3] = 0, 0, 0, 0
	} else {
		s := Number(1 / l)
		q[0] *= s
		q[1] *= s
		q[2] *= s
		q[3] *= s
	}
	return q
}

/**
 * Approximation of quaternion normalization. Works best when quat is already almost-normalized.
 * @method normalizeFast
 * @see http://jsperf.com/fast-quaternion-normalization
 * @author unphased, https://github.com/unphased
 */
func (q *Quat) NormalizeFast() (*Quat) {
	x, y, z, w := q[0], q[1], q[2], q[3]
	f := (3.0 - (x*x + y*y + z*z + w*w) ) / 2.0

	if f == 0 {
		q[0], q[1], q[2], q[3] = 0, 0, 0, 0
	} else {
		q[0] *= f
		q[1] *= f
		q[2] *= f
		q[3] *= f
	}
	return q
}


/**
 * Multiply the quaternion by a vector
 * @method vmult
 * @param {Vec3} v
 * @param {Vec3} target Optional
 * @return {Vec3}
 */
func (q *Quat) VMult(v *Vec3, target *Vec3) (*Vec3) {
	if target == nil {
		target = NewVec3()
	}

	x, y, z := v[0], v[1], v[2]
	qx, qy, qz, qw := q[0], q[1], q[2], q[3]

	// q*v
	ix :=  qw * x + qy * z - qz * y
	iy :=  qw * y + qz * x - qx * z
	iz :=  qw * z + qx * y - qy * x
	iw := -qx * x - qy * y - qz * z

	target[0] = ix * qw + iw * -qx + iy * -qz - iz * -qy
	target[1] = iy * qw + iw * -qy + iz * -qx - ix * -qz
	target[2] = iz * qw + iw * -qz + ix * -qy - iy * -qx

	return target
}


/**
 * Convert the quaternion to euler angle representation. Order: YZX, as this page describes: http://www.euclideanspace.com/maths/standards/index.htm
 * @method toEuler
 * @param {Vec3} target
 * @param string order Three-character string e.g. "YZX", which also is default.
 */
func (q *Quat) ToEuler(target *Vec3, order AxisOrder) (*Vec3) {
	if target == nil {
		target = NewVec3()
	}

	var heading, attitude, bank float64
	x, y, z, w := q[0], q[1], q[2], q[3]

	switch order {
	default:
		fallthrough // TODO: not supported yet. :(
	case YZX:
		test := x*y + z*w
		if test > 0.499 { // singularity at north pole
			heading = 2 * math.Atan2(float64(x), float64(w))
			attitude = math.Pi / 2
			bank = 0
		}
		if test < -0.499 { // singularity at south pole
			heading = -2 * math.Atan2(float64(x), float64(w))
			attitude = - math.Pi / 2
			bank = 0
		}
		if math.IsNaN(heading) {
			sqx := x*x
			sqy := y*y
			sqz := z*z
			heading = math.Atan2(float64(2*y*w - 2*x*z) , float64(1 - 2*sqy - 2*sqz)) // Heading
			attitude = math.Asin(float64(2*test)) // attitude
			bank = math.Atan2(float64(2*x*w - 2*y*z) , float64(1 - 2*sqx - 2*sqz) ) // bank
		}
	}

	target[0] = Number(heading) // yaw, theta
	target[1] = Number(attitude) // pitch, phi
	target[2] = Number(bank) // roll, psi

	return target
}


/**
 * See http://www.mathworks.com/matlabcentral/fileexchange/20696-function-to-convert-between-dcm-euler-angles-quaternions-and-euler-vectors/content/SpinCalc.m
 * @method setFromEuler
 * @param {Number} x
 * @param {Number} y
 * @param {Number} z
 * @param {String} order The order to apply angles: 'XYZ' or 'YXZ' or any other combination
 */
func (q *Quat) SetFromEuler(x Number,y Number,z Number, order AxisOrder) (*Quat) {

	c, s := math.Sincos(float64( x / 2))
	c1, s1 := Number(c), Number(s)

	c, s = math.Sincos(float64( y / 2))
	c2, s2 := Number(c), Number(s)

	c, s = math.Sincos(float64( z / 2))
	c3, s3 := Number(c), Number(s)

	switch order {
	default:
		fallthrough
	case XYZ:
		q[0] = s1 * c2 * c3 + c1 * s2 * s3
		q[1] = c1 * s2 * c3 - s1 * c2 * s3
		q[2] = c1 * c2 * s3 + s1 * s2 * c3
		q[3] = c1 * c2 * c3 - s1 * s2 * s3
	case YXZ:
		q[0] = s1 * c2 * c3 + c1 * s2 * s3
		q[1] = c1 * s2 * c3 - s1 * c2 * s3
		q[2] = c1 * c2 * s3 - s1 * s2 * c3
		q[3] = c1 * c2 * c3 + s1 * s2 * s3
	case ZXY:
		q[0] = s1 * c2 * c3 - c1 * s2 * s3
		q[1] = c1 * s2 * c3 + s1 * c2 * s3
		q[2] = c1 * c2 * s3 + s1 * s2 * c3
		q[3] = c1 * c2 * c3 - s1 * s2 * s3
	case ZYX:
		q[0] = s1 * c2 * c3 - c1 * s2 * s3
		q[1] = c1 * s2 * c3 + s1 * c2 * s3
		q[2] = c1 * c2 * s3 - s1 * s2 * c3
		q[3] = c1 * c2 * c3 + s1 * s2 * s3
	case YZX:
		q[0] = s1 * c2 * c3 + c1 * s2 * s3
		q[1] = c1 * s2 * c3 + s1 * c2 * s3
		q[2] = c1 * c2 * s3 - s1 * s2 * c3
		q[3] = c1 * c2 * c3 - s1 * s2 * s3
	case XZY:
		q[0] = s1 * c2 * c3 - c1 * s2 * s3
		q[1] = c1 * s2 * c3 - s1 * c2 * s3
		q[2] = c1 * c2 * s3 + s1 * s2 * c3
		q[3] = c1 * c2 * c3 + s1 * s2 * s3
	}

	return q
}


/**
 * Performs a spherical linear interpolation between two quat
 *
 * @method slerp
 * @param {Quaternion} toQuat second operand
 * @param {Number} t interpolation amount between the self quaternion and toQuat
 * @param {Quaternion} [target] A quaternion to store the result in. If not provided, a new one will be created.
 * @returns {Quaternion} The "target" object
 */
func (q *Quat) Slerp(toQuat *Quat, t Number, target *Quat) (*Quat) {
	if target == nil {
		target = &Quat{}
	}

	ax, ay, az, aw := q[0], q[1], q[2], q[3]
	bx, by, bz, bw := toQuat[0], toQuat[1], toQuat[2], toQuat[3]

	// calc cosine
	cosom := ax * bx + ay * by + az * bz + aw * bw

	// adjust signs (if necessary)
	if ( cosom < 0.0 ) {
		cosom = -cosom
		bx = - bx
		by = - by
		bz = - bz
		bw = - bw
	}

	var scale0 Number
	var scale1 Number

	// calculate coefficients
	if ( (1.0 - cosom) > 0.000001 ) {
		// standard case (slerp)
		omega  := math.Acos(float64(cosom))
		sinom  := Number(math.Sin(float64(omega)))

		t0 := math.Sin(float64(1.0 - t) * omega)
		scale0 = Number(t0) / sinom

		t0 = math.Sin(float64(t) * omega)
		scale1 = Number(t0) / sinom
	} else {
		// "from" and "to" quaternions are very close
		//  ... so we can do a linear interpolation
		scale0 = 1.0 - t
		scale1 = t
	}

	// calculate final values
	target[0] = scale0 * ax + scale1 * bx
	target[1] = scale0 * ay + scale1 * by
	target[2] = scale0 * az + scale1 * bz
	target[3] = scale0 * aw + scale1 * bw

	return target
}


/**
 * Rotate an absolute orientation quaternion given an angular velocity and a time step.
 * @param  {Vec3} angularVelocity
 * @param  {number} dt
 * @param  {Vec3} angularFactor
 * @param  {Quaternion} target
 * @return {Quaternion} The "target" object
 */
func (q *Quat) Integrate(angularVelocity *Vec3, dt Number, angularFactor *Vec3, target *Quat) (*Quat) {
	if target == nil {
		target = NewQuat()
	}

	ax := angularVelocity[0] * angularFactor[0]
	ay := angularVelocity[1] * angularFactor[1]
	az := angularVelocity[2] * angularFactor[2]

	bx, by, bz, bw := q[0], q[1], q[2], q[3]

	half_dt := dt * 0.5

	target[0] += half_dt * (ax * bw + ay * bz - az * by)
	target[1] += half_dt * (ay * bw + az * bx - ax * bz)
	target[2] += half_dt * (az * bw + ax * by - ay * bx)
	target[3] += half_dt * (- ax * bx - ay * by - az * bz)

	return target
}

