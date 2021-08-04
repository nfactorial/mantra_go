package shading

import "strange-secrets.com/mantra/algebra"

type Material interface {
	Evaluate(info ShadeInfo) algebra.Vector3
}
