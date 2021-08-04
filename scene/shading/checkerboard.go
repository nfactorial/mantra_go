package shading

import (
	"math"
	"strange-secrets.com/mantra/algebra"
)

const CheckerboardMaterialName = "checkerboard"

var (
	defaultCheckerboardA = algebra.Vector3{0,0,0}
	defaultCheckerboardB = algebra.Vector3{1,1,1}
)

type CheckerboardMaterial struct {
	ColorA algebra.Vector3
	ColorB algebra.Vector3
}

func NewCheckerboardMaterial() Material {
	return &CheckerboardMaterial{
		ColorA: defaultCheckerboardA,
		ColorB: defaultCheckerboardB,
	}
}

func (c *CheckerboardMaterial) Evaluate(info ShadeInfo) algebra.Vector3 {
	x := info.Point.X / 2.5
	y := info.Point.Y / 2.5
	z := info.Point.Z / 2.5

	if x < 0.0 {
		x = -x + 1
	}

	if y < 0.0 {
		y = -y + 1
	}

	if z < 0.0 {
		z = -z + 1
	}

	x = math.Mod(x, 2.0)
	y = math.Mod(y, 2.0)
	z = math.Mod(z, 2.0)

	selected := 0

	if x >= 1.0 {
		selected = selected ^ 1
	}

	if y >= 1.0 {
		selected = selected ^ 1
	}

	if z >= 1.0 {
		selected = selected ^ 1
	}

	if selected == 0 {
		return c.ColorA
	}

	return c.ColorB
}
