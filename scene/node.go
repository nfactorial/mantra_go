package scene

import (
	"strange-secrets.com/mantra/algebra"
	"strange-secrets.com/mantra/scene/camera"
	"strange-secrets.com/mantra/scene/geometry"
	"strange-secrets.com/mantra/scene/lighting"
	"strange-secrets.com/mantra/scene/shading"
)

const (
	DefaultCastShadows = true
	DefaultReceiveShadows = true

	nodeTypeGeometry = 0
	nodeTypeCamera = 1
	nodeTypeLight = 1
)

type Node struct {
	name           	string
	nodeType		int
	light			lighting.Light
	CastShadows    	bool
	ReceiveShadows 	bool
	Camera			*camera.Camera
	Color          	algebra.Vector3
	Location       	algebra.Vector3
	Geometry       	geometry.Geometry
	Transform      	algebra.Matrix4		// TODO: Transform should be a quaternion
	Surface		   	shading.SurfaceInfo
}

func NewNode(name string, nodeType int) *Node {
	return &Node{
		name:     		name,
		nodeType:		nodeType,
		light:			nil,
		Camera:			nil,
		CastShadows:	DefaultCastShadows,
		ReceiveShadows: DefaultReceiveShadows,
		Color:			algebra.ZeroVector3,
		Location:		algebra.ZeroVector3,
		Transform:		algebra.IdentityMatrix4,
		Surface:		shading.SurfaceInfo{
			Material: &shading.SolidColorMaterial{
				Color: algebra.Vector3{1.0, 1.0, 1.0},
			},
			Metallic: 0,
			Kr:		  0.2,
			Ka:       0.2,
			Kd:       0.6,
			Ks:       0.3,
		},
		Geometry: nil,
	}
}

// Retrieves the name of the node
func (n *Node) Name() string {
	return n.name
}

func (n *Node) CastRay(x algebra.MnFloat, y algebra.MnFloat, aspect algebra.MnFloat) algebra.Vector3 {
	if n.Camera == nil {
		return algebra.ForwardVector3	// Maybe retrieve from our transform
	}

	direction := n.Camera.CastRay(x, y, aspect)

	return direction.Transform(n.Transform)
}

func (n *Node) LookAt(target algebra.Vector3) {
	forward := n.Location.Subtract(target).Normalize()
	right := algebra.UpVector3.Cross(forward).Normalize()
	up := forward.Cross(right).Normalize()

	n.Transform = algebra.NewRotationMatrix4(right, up, forward)
}

// Determines whether or not the specified ray intersects the nodes geometry
func (n *Node) HitTest(ray algebra.Ray, r algebra.Range) geometry.HitResult {
	if n.Geometry != nil {
		// Transform incoming ray to local space
		localRay := algebra.Ray{
			Origin:    ray.Origin.Subtract(n.Location),
			Direction: ray.Direction.Transform(n.Transform.Invert()),
		}

		// TODO: Check ray intersects bounding box and is within the range

		result := n.Geometry.HitTest(localRay)
		if result != geometry.InvalidHitResult && r.Contains(result.Distance) {
			// Transform result back to world space
			result.Normal = result.Normal.Transform(n.Transform)
			result.Location = result.Location.Transform(n.Transform).Add(n.Location)
			return result
		}
	}

	return geometry.InvalidHitResult
}
