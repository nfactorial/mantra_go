package shading

import (
	"strange-secrets.com/mantra/algebra"
)

type ShadeInfo struct {
	Point	 algebra.Vector3	// Point in 3D space (world coordinates)
	Normal	 algebra.Vector3	// Surface normal at point
	Incident algebra.Vector3	// Incident vector at surface point
	View	 algebra.Ray		// Original viewer
	Surface  SurfaceInfo		// Properties of surface
	Shadow	 bool				// True if it is a shadow otherwise false
}
