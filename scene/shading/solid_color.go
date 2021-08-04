package shading

import "strange-secrets.com/mantra/algebra"

const SolidColorMaterialName = "solid_color"

var defaultSolidColor = algebra.Vector3{1.0, 1.0, 1.0}

type SolidColorMaterial struct {
	Color algebra.Vector3
}

func NewSolidColorMaterial() Material {
	return &SolidColorMaterial{
		Color: defaultSolidColor,
	}
}

func (s *SolidColorMaterial) Evaluate(shadeInfo ShadeInfo) algebra.Vector3 {
	return s.Color
}
