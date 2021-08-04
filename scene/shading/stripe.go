package shading

import (
	"math"
	"strange-secrets.com/mantra/algebra"
)

const StripeMaterialName = "stripe"

var (
	defaultStripeA = algebra.Vector3{0,0,0 }
	defaultStripeB = algebra.Vector3{1,1,1 }
)

type StripeMaterial struct {
	ColorA algebra.Vector3
	ColorB algebra.Vector3
}

func NewStripeMaterial() Material {
	return &StripeMaterial{
		ColorA: defaultStripeA,
		ColorB: defaultStripeB,
	}
}

func (s *StripeMaterial) Evaluate(info ShadeInfo) algebra.Vector3 {
	x := info.Point.Z / 2.5
	if x < 0 {
		x = -x + 1
	}

	x = math.Mod(x, 2.0)
	if x < 1 {
		return s.ColorA
	}

	return s.ColorB
}
