package scene

import (
	"strange-secrets.com/mantra/algebra"
	"strange-secrets.com/mantra/scene/geometry"
	"strange-secrets.com/mantra/scene/shading"
)

type SphereFlake struct {
	Origin		algebra.Vector3
	Radius		algebra.MnFloat
	Depth		int
	Scene		*Scene
}

func (s *SphereFlake) Generate() *Scene {
	return s.GenerateSphere(s.Origin, s.Radius, 0)
}

func (s *SphereFlake) GenerateSphere(origin algebra.Vector3, radius algebra.MnFloat, depth int) *Scene {
	n, _ := s.Scene.CreateGeometryNode(s.Scene.CreateUniqueName("sphere_flake"), shading.SolidColorMaterialName, &geometry.SphereGeometry{
		Radius: radius,
	})
	n.Location = origin

	if depth < s.Depth {
		halfRadius := radius / 2.0

		s.GenerateSphere(origin.Add(algebra.Vector3{
			X: radius + halfRadius,
			Y: 0,
			Z: 0,
		}), halfRadius, depth + 1)

		s.GenerateSphere(origin.Add(algebra.Vector3{
			X: -(radius + halfRadius),
			Y: 0,
			Z: 0,
		}), halfRadius, depth + 1)

		s.GenerateSphere(origin.Add(algebra.Vector3{
			X: 0,
			Y: radius + halfRadius,
			Z: 0,
		}), halfRadius, depth + 1)

		s.GenerateSphere(origin.Add(algebra.Vector3{
			X: 0,
			Y: -(radius + halfRadius),
			Z: 0,
		}), halfRadius, depth + 1)
/*
		s.GenerateSphere(origin.Add(algebra.Vector3{
			X: 0,
			Y: 0,
			Z: radius + halfRadius,
		}), halfRadius, depth + 1)

		s.GenerateSphere(origin.Add(algebra.Vector3{
			X: 0,
			Y: 0,
			Z: -(radius + halfRadius),
		}), halfRadius, depth + 1)

 */
	}

	return s.Scene
}
