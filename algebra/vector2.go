package algebra

import (
	"fmt"
	"math"
)

type Vector2 struct {
	X MnFloat
	Y MnFloat
}

var (
	ZeroVector2 = Vector2{0.0, 0.0}
)

func (v Vector2) Dump() Vector2 {
	fmt.Println(fmt.Sprintf("X = %f, Y = %f", v.X, v.Y))
	return v
}

// Calculates the dot product between two vectors.
func (v Vector2) Dot(other Vector2) MnFloat {
	return v.X*other.X + v.Y*other.Y
}

// Calculates the squared magnitude of the vector.
func (v Vector2) LengthSq() MnFloat {
	return v.Dot(v)
}

// Calculates the magnitude of the vector
func (v Vector2) Length() MnFloat {
	return Sqrt(v.LengthSq())
}

// Returns the vector with a magnitude of 1.
func (v Vector2) Normalize() Vector2 {
	length := v.Length()

	return Vector2{
		v.X / length,
		v.Y / length,
	}
}

// Computes the vector result of adding two vectors together.
func (v Vector2) Add(other Vector2) Vector2 {
	return Vector2{
		v.X + other.X,
		v.Y + other.Y,
	}
}

// Computes the vector result of subtracting two vectors together.
func (v Vector2) Subtract(other Vector2) Vector2 {
	return Vector2{
		v.X - other.X,
		v.Y - other.Y,
	}
}

// Computes the vector result by multiplying each component by a scalar value.
func (v Vector2) MultiplyScalar(value MnFloat) Vector2 {
	return Vector2{
		v.X * value,
		v.Y * value,
	}
}

func (v Vector2) Multiply(other Vector2) Vector2 {
	return Vector2{
		v.X * other.X,
		v.Y * other.Y,
	}
}

// Computes the vector consisting of the smallest component values from two vectors.
func (v Vector2) Min(other Vector2) Vector2 {
	return Vector2{
		math.Min(v.X, other.X),
		math.Min(v.Y, other.Y),
	}
}

// Computes the vector consisting of the largest  component values from two vectors.
func (v Vector2) Max(other Vector2) Vector2 {
	return Vector2{
		math.Max(v.X, other.X),
		math.Max(v.Y, other.Y),
	}
}

func (v Vector2) Negate() Vector2 {
	return Vector2{
		-v.X,
		-v.Y,
	}
}
// Performs a linear interpolation between two vectors and returns the result.
func (v Vector2) Lerp(other Vector2, t MnFloat) Vector2 {
	return Vector2{
		Lerp(v.X, other.X, t),
		Lerp(v.Y, other.Y, t),
	}
}
