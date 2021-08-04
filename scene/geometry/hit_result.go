package geometry

import (
	"math"
	"strange-secrets.com/mantra/algebra"
)

type HitResult struct {
	Normal algebra.Vector3
	Location algebra.Vector3
	Distance algebra.MnFloat
}

var InvalidHitResult = HitResult{
	Distance: math.Inf(1),
}
