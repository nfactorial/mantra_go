package geometry

import (
	"strange-secrets.com/mantra/algebra"
)

type Geometry interface {
	HitTest(ray algebra.Ray) HitResult
}
