package scene

import (
	"strange-secrets.com/mantra/algebra"
	"strange-secrets.com/mantra/scene/geometry"
)

type TraceResult struct {
	Node *Node
	Color algebra.Vector3
	HitResult geometry.HitResult
}
