package geometry

import (
	"math"
	"strange-secrets.com/mantra/algebra"
)

// NOTE: It maybe more suitable to move these methods into another package (algebra/intersections etc).

// Calculates the intersection between a ray and a plane in 3D space.
func IntersectRayPlane(ray algebra.Ray, normal algebra.Vector3, d algebra.MnFloat) HitResult {
	denom := normal.Dot(ray.Direction.Negate())
	if math.Abs(denom) > algebra.Epsilon {
		t := (normal.Dot(ray.Origin) - d) / denom
		if t >= 0 {
			return HitResult{
				Normal: normal,
				Location: ray.GetPosition(t),
				Distance: t,
			}
		}
	}

	return InvalidHitResult
}

// Calculates the intersection between a ray and a sphere in 3D space.
func IntersectSphere(ray algebra.Ray, spherePosition algebra.Vector3, radius algebra.MnFloat) HitResult {
	l := spherePosition.Subtract(ray.Origin)

	radius2 := radius * radius

	tca := l.Dot(ray.Direction)
	if tca >= 0.0 {
		d2 := l.Dot(l) - (tca*tca)
		if d2 <= radius2 {
			thc := math.Sqrt(radius2 - d2)
			t0 := tca - thc
			t1 := tca + thc

			if t0 > t1 {
				temp := t0
				t0 = t1
				t1 = temp
			}

			if t0 < 0 {
				t0 = t1
			}

			if t0 >= 0 {
				position := ray.GetPosition(t0)

				return HitResult{
					Normal:   position.Subtract(spherePosition).Normalize(),
					Location: position,
					Distance: t0,
				}
			}
		}
	}

	return InvalidHitResult
}
