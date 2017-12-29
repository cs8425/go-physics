package physics


type Transform struct {
	Pos *Vec3
	Rot *Quat
}

/**
 * Get a global point in local transform coordinates.
 * @method pointToLocal
 * @param  {Vec3} point
 * @param  {Vec3} result
 * @return {Vec3} The "result" vector object
 */
func (tf *Transform) PointToLocal(worldPoint *Vec3, result *Vec3) (*Vec3) {
	return TransformPointToLocalFrame(tf.Pos, tf.Rot, worldPoint, result)
}


/**
 * Get a local point in global transform coordinates.
 * @method pointToWorld
 * @param  {Vec3} point
 * @param  {Vec3} result
 * @return {Vec3} The "result" vector object
 */
func (tf *Transform) PointToWorld(localPoint *Vec3, result *Vec3) (*Vec3) {
	return TransformPointToWorldFrame(tf.Pos, tf.Rot, localPoint, result)
}


func (tf *Transform) VectorToWorldFrame(localVector *Vec3, result *Vec3) (*Vec3) {
	if result == nil {
		result = &Vec3{}
	}
	tf.Rot.VMult(localVector, result)
	return result
}


/**
 * @static
 * @method pointToLocaFrame
 * @param {Vec3} position
 * @param {Quaternion} quaternion
 * @param {Vec3} worldPoint
 * @param {Vec3} result
 */
func TransformPointToLocalFrame(position *Vec3, quaternion *Quat, worldPoint *Vec3, result *Vec3) (*Vec3) {
	if result == nil {
		result = &Vec3{}
	}
	worldPoint.VSub(position, result)
	tmpQuat := &Quat{}
	quaternion.Conjugate(tmpQuat)
	tmpQuat.VMult(result, result)
	return result
}


/**
 * @static
 * @method pointToWorldFrame
 * @param {Vec3} position
 * @param {Vec3} quaternion
 * @param {Vec3} localPoint
 * @param {Vec3} result
 */
func TransformPointToWorldFrame(position *Vec3, quaternion *Quat, localPoint *Vec3, result *Vec3) (*Vec3) {
	if result == nil {
		result = &Vec3{}
	}
	quaternion.VMult(localPoint, result)
	result.VAdd(position, result)
	return result
}


func TransformVectorToWorldFrame(quaternion *Quat, localVector *Vec3, result *Vec3) (*Vec3) {
	if result == nil {
		result = &Vec3{}
	}
	quaternion.VMult(localVector, result)
	return result
}


func TransformVectorToLocalFrame(position *Vec3, quaternion *Quat, worldVector *Vec3, result *Vec3) (*Vec3) {
	if result == nil {
		result = &Vec3{}
	}
	quaternion[3] *= -1
	quaternion.VMult(worldVector, result)
	quaternion[3] *= -1
	return result
}



