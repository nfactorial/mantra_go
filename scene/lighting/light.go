package lighting

import (
	"strange-secrets.com/mantra/algebra"
	"strange-secrets.com/mantra/scene/shading"
)

type TraceShadow func(ray algebra.Ray, valid algebra.Range) bool

type Light interface {
	Ambient(info shading.ShadeInfo) algebra.Vector3
	Diffuse(info shading.ShadeInfo) algebra.Vector3
	Specular(info shading.ShadeInfo) algebra.Vector3
	Shadow(info shading.ShadeInfo, shadow TraceShadow) algebra.MnFloat
}
