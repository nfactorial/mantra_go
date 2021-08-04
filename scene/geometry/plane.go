package geometry

import "strange-secrets.com/mantra/algebra"

type PlaneGeometry struct {
	Normal algebra.Vector3
	Distance algebra.MnFloat
}

func (p *PlaneGeometry) HitTest(ray algebra.Ray) HitResult {
	return IntersectRayPlane(ray, p.Normal, p.Distance)
}
