package camera

import (
	"strange-secrets.com/mantra/algebra"
)

const (
	PerspectiveCamera = 0
	OrthographicCamera = 1
)

type Camera struct {
	Type int
	Near algebra.MnFloat
	Far algebra.MnFloat
	FieldOfView algebra.MnFloat
}

func (c Camera) CastRay(x algebra.MnFloat, y algebra.MnFloat, aspect algebra.MnFloat) algebra.Vector3 {
	inverseProjection := algebra.NewPerspectiveMatrix4(c.FieldOfView, aspect, c.Near, c.Far).Invert()

	return algebra.Vector3{
		X: x,
		Y: y,
		Z: 0.0,
	}.Transform(inverseProjection).Normalize()
}
