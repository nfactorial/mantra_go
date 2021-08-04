package geometry

import (
	"strange-secrets.com/mantra/algebra"
)

type SphereGeometry struct {
	Radius algebra.MnFloat
}

func (s *SphereGeometry) HitTest(ray algebra.Ray) HitResult {
	return IntersectSphere(ray, algebra.ZeroVector3, s.Radius)
}
