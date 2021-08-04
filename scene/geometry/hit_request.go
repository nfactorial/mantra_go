package geometry

import "strange-secrets.com/mantra/algebra"

type HitRequest struct {
	Ray algebra.Ray					// The ray to be tested against
	Range algebra.Range				// The valid hit range
	Position algebra.Vector3		// The position of the object being tested
	Transform algebra.Matrix4		// The transform of the object being tested
}
