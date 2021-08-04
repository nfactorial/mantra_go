package algebra

import "math"

type MnFloat = float64

const (
	//Epsilon MnFloat = 1e-6
	Epsilon MnFloat = 0.001
	MnOne = 1.0
	MnZero = 0.0
)

var (
	NegativeInfinity = math.Inf(-1)
	PositiveInfinity = math.Inf(1)
)

func Lerp(a MnFloat, b MnFloat, t MnFloat) MnFloat {
	return a + t*(b - a)
}

func Sqrt(x MnFloat) MnFloat {
	return MnFloat(math.Sqrt(x))
}

func ToRadians(degrees MnFloat) MnFloat {
	return degrees * (math.Pi / 180.0)
}

func ToDegrees(radians MnFloat) MnFloat {
	return radians * (180.0 / math.Pi)
}

func Step(x MnFloat, y MnFloat) MnFloat {
	if x < y {
		return 0.0
	}

	return 1.0
}

func Swap(x MnFloat, y MnFloat) (MnFloat, MnFloat) {
	return y, x
}

func SmoothStep(min MnFloat, max MnFloat, x MnFloat) MnFloat {
	if x < min {
		return 0.0
	}

	if x > max {
		return 1.0
	}

	return (x - min) / (max - min)
}
