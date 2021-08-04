package lighting

import (
	"math"
	"strange-secrets.com/mantra/algebra"
	"strange-secrets.com/mantra/scene/shading"
)

type Directional struct {
	Direction algebra.Vector3
	Color algebra.Vector3
}

const (
	SpecularPower = 6.0
)

func (d *Directional) Ambient(info shading.ShadeInfo) algebra.Vector3 {
	return algebra.ZeroVector3
}

func (d *Directional) Diffuse(info shading.ShadeInfo) algebra.Vector3 {
	i := math.Max(0.0, d.Direction.Negate().Dot(info.Normal))
	return algebra.Vector3{i, i, i}
}

func (d *Directional) Specular(info shading.ShadeInfo) algebra.Vector3 {
	l := d.Direction.Negate()
	r := l.Reflect(info.Normal)
	i := math.Pow(r.Dot(info.View.Direction), SpecularPower)

	return algebra.Vector3{i, i, i}
}

func (d *Directional) Shadow(info shading.ShadeInfo, shadow TraceShadow) algebra.MnFloat {
	result := 0.0

	if shadow(algebra.Ray{
		Origin:    info.Point,
		Direction: d.Direction.Negate(),
	}, algebra.FrontRange) {
		return 1.0
	}

	return result
}
