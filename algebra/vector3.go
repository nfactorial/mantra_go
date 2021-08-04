package algebra

import (
	"fmt"
	"image/color"
	"math"
)

type Vector3 struct {
	X MnFloat
	Y MnFloat
	Z MnFloat
}

var (
	ZeroVector3		= Vector3{0, 0, 0}
	UpVector3		= Vector3{0, 1, 0}
	RightVector3	= Vector3{1, 0, 0}
	ForwardVector3	= Vector3{0, 0, -1}
)

func (v Vector3) Dump() Vector3 {
	fmt.Println(fmt.Sprintf("X = %f, Y = %f, Z = %f", v.X, v.Y, v.Z))
	return v
}

// Calculates the dot product between two vectors.
func (v Vector3) Dot(other Vector3) MnFloat {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

// Calculates the squared magnitude of the vector.
func (v Vector3) LengthSq() MnFloat {
	return v.Dot(v)
}

// Calculates the magnitude of the vector
func (v Vector3) Length() MnFloat {
	return Sqrt(v.LengthSq())
}

// Returns the vector with a magnitude of 1.
func (v Vector3) Normalize() Vector3 {
	length := v.Length()

	return Vector3{
		v.X / length,
		v.Y / length,
		v.Z / length,
	}
}

func (v Vector3) Pow(t MnFloat) Vector3 {
	return Vector3{
		math.Pow(v.X, t),
		math.Pow(v.Y, t),
		math.Pow(v.Z, t),
	}
}

// Computes the vector result of adding two vectors together.
func (v Vector3) Add(other Vector3) Vector3 {
	return Vector3{
		v.X + other.X,
		v.Y + other.Y,
		v.Z + other.Z,
	}
}

// Computes the vector result of subtracting two vectors together.
func (v Vector3) Subtract(other Vector3) Vector3 {
	return Vector3{
		v.X - other.X,
		v.Y - other.Y,
		v.Z - other.Z,
	}
}

func (v Vector3) DivideScalar(value MnFloat) Vector3 {
	return Vector3{
		v.X / value,
		v.Y / value,
		v.Z / value,
	}
}

// Computes the vector result by multiplying each component by a scalar value.
func (v Vector3) MultiplyScalar(value MnFloat) Vector3 {
	return Vector3{
		v.X * value,
		v.Y * value,
		v.Z * value,
	}
}

func (v Vector3) Multiply(other Vector3) Vector3 {
	return Vector3{
		v.X * other.X,
		v.Y * other.Y,
		v.Z * other.Z,
	}
}

// Computes the cross product of two vectors.
func (v Vector3) Cross(other Vector3) Vector3 {
	return Vector3{
		(v.Y * other.Z) - (v.Z * other.Y),
		(v.Z * other.X) - (v.X * other.Z),
		(v.X * other.Y) - (v.Y * other.X),
	}
}

func (v Vector3) Saturate() Vector3 {
	return Vector3{
		math.Max(MnZero, math.Min(MnOne, v.X)),
		math.Max(MnZero, math.Min(MnOne, v.Y)),
		math.Max(MnZero, math.Min(MnOne, v.Z)),
	}
}

// Computes the vector consisting of the smallest component values from two vectors.
func (v Vector3) Min(other Vector3) Vector3 {
	return Vector3{
		math.Min(v.X, other.X),
		math.Min(v.Y, other.Y),
		math.Min(v.Z, other.Z),
	}
}

// Computes the vector consisting of the largest  component values from two vectors.
func (v Vector3) Max(other Vector3) Vector3 {
	return Vector3{
		math.Max(v.X, other.X),
		math.Max(v.Y, other.Y),
		math.Max(v.Z, other.Z),
	}
}

func (v Vector3) Negate() Vector3 {
	return Vector3{
		-v.X,
		-v.Y,
		-v.Z,
	}
}
// Performs a linear interpolation between two vectors and returns the result.
func (v Vector3) Lerp(other Vector3, t MnFloat) Vector3 {
	return Vector3{
		Lerp(v.X, other.X, t),
		Lerp(v.Y, other.Y, t),
		Lerp(v.Z, other.Z, t),
	}
}

// Computes a reflection vector, given the surface normal.
func (v Vector3) Reflect(normal Vector3) Vector3 {
	return v.Subtract(normal.MultiplyScalar(2.0 * v.Dot(normal)))
}

// Computes the refracted vector given the surface normal and index of refraction (ior).
func (v Vector3) Refract(normal Vector3, ior MnFloat) Vector3 {
	a := v.Dot(normal)
	k := 1.0 - ior * ior * (1.0 - a * a)
	if k < 0.0 {
		return Vector3{0.0, 0.0,0.0 }
	}

	return v.MultiplyScalar(ior).Subtract(normal.MultiplyScalar(ior * a + Sqrt(k)))
}

func (v Vector3) ToColor() color.NRGBA {
	return color.NRGBA{
		R: (uint8)(v.X * 255),
		G: (uint8)(v.Y * 255),
		B: (uint8)(v.Z * 255),
		A: 255,
	}
}

func (v Vector3) Transform(m Matrix4) Vector3 {
	w := m.M03 * v.X + m.M13 * v.Y + m.M23 * v.Z + m.M33
	if w == MnZero {
		w = MnOne
	}

	return Vector3{
		X: (m.M00 * v.X + m.M10 * v.Y + m.M20 * v.Z + m.M30) / w,
		Y: (m.M01 * v.X + m.M11 * v.Y + m.M21 * v.Z + m.M31) / w,
		Z: (m.M02 * v.X + m.M12 * v.Y + m.M22 * v.Z + m.M32) / w,
	}
}