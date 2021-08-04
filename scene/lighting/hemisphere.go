package lighting

import (
	"strange-secrets.com/mantra/algebra"
	"strange-secrets.com/mantra/scene/shading"
)

type Hemisphere struct {
	Up			algebra.Vector3
	UpperColor	algebra.Vector3
	LowerColor	algebra.Vector3
}

func (h *Hemisphere) Ambient(info shading.ShadeInfo) algebra.Vector3 {
	return algebra.ZeroVector3
}

func (h *Hemisphere) Diffuse(info shading.ShadeInfo) algebra.Vector3 {
	w := 0.5 * (1.0 + info.Normal.Dot(h.Up))
	return h.UpperColor.MultiplyScalar(w).Add(h.LowerColor.MultiplyScalar(1.0 - w))
}

func (h *Hemisphere) Specular(info shading.ShadeInfo) algebra.Vector3 {
	return algebra.ZeroVector3
}

func (h *Hemisphere) Shadow(info shading.ShadeInfo, shadow TraceShadow) algebra.MnFloat {
	return 0
}

