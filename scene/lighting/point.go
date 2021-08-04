package lighting

import (
	"fmt"
	"math"
	"strange-secrets.com/mantra/algebra"
	"strange-secrets.com/mantra/scene/shading"
)

type Point struct {
	Radius algebra.MnFloat
	Color algebra.Vector3
}

func (p *Point) Ambient(info shading.ShadeInfo) algebra.Vector3 {
	return algebra.ZeroVector3
}

var lightd = 0
func (p *Point) Diffuse(info shading.ShadeInfo) algebra.Vector3 {
	d := math.Max(p.Radius - info.Point.Length(), 0.0) / p.Radius
	l := info.Point.Negate().Normalize()

	i := math.Max(0.0, l.Dot(info.Normal)) * d

	if lightd == 0 && d > 0 {
		lightd++
		fmt.Println(fmt.Sprintf("d = %f, i = %f", d, i))
	}

	return algebra.Vector3{i, i, i}
}

func (p *Point) Specular(info shading.ShadeInfo) algebra.Vector3 {
	/*
	d := math.Max(p.Radius - info.Point.Length(), 0.0) / p.Radius
	l := info.Point.Negate().Normalize()

	r := l.Reflect(info.Normal)
	i := math.Pow(r.Dot(info.View.Direction), SpecularPower) * d

	return algebra.Vector3{i, i, i}
	*/
	return algebra.ZeroVector3
}

func (p *Point) Shadow(info shading.ShadeInfo, shadow TraceShadow) algebra.MnFloat {
	return 0.0
}
