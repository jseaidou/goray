package math

import "math"

type Vector struct {
	X float64
	Y float64
	Z float64
}

func NewVector(x, y, z float64) Vector {
	return Vector{
		X: x,
		Y: y,
		Z: z,
	}
}

/**
 * Returns the norm of a vector denoted by:
 * sqrt(x^2 + y^2 + z^2) = sqrt(v*v) where * is the dot product
 */
func (v Vector) Norm() float64 {
	return math.Sqrt(v.Dot(v))
}

/**
 * Normalizes a vector: x/|x| where |x| is the norm of the vector
 */
func (v Vector) Normalize() Vector {
	norm := v.Norm()
	if norm == 0 {
		// A 0 norm occurs iff a vector v is the 0 vector
		return v
	}
	return NewVector(v.X/norm, v.Y/norm, v.Z/norm)
}

func (v Vector) Add(w Vector) Vector {
	return NewVector(v.X+w.X, v.Y+w.Y, v.Z+w.Z)
}

func (v Vector) Mult(c float64) Vector {
	return NewVector(c*v.X, c*v.Y, c*v.Z)
}

func (v Vector) Cross(w Vector) Vector {
	return NewVector(v.Y*w.Z-v.Z*w.Y, v.Z*w.X-v.X*w.Z, v.X*w.Y-v.Y*w.Z)
}

func (v Vector) Dot(w Vector) float64 {
	return v.X*w.X + v.Y*w.Y + v.Z*w.Z
}

/**
 * Returns whether a vector V is orthogonal to vector W defined by if the dot product is equal to 0
 * since dot product is equivalent to |V|*|W|cos(theta) where |V| and |W| are the norms of V and W
 * respectively and and theta is the angle between V and W. Hence if V and W are orthogonal
 * cos(theta) = cos(90) = cos(PI/2) = 0 and V and W are orthogonal
 */
func (v Vector) IsOrth(w Vector) bool {
	return v.Dot(w) == 0
}
