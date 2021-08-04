package scene

import (
	"strange-secrets.com/mantra/algebra"
	"strange-secrets.com/mantra/scene/shading"
)

const (
	DefaultMaximumShadowDepth = 2
	DefaultMaximumTraceDepth = 4
	DefaultIor = 1.0
)

// Describes a ray being traced through the scene.
type TraceInfo struct {
	Rendering shading.RenderInfo	// Information about the of the pixel being rendered
	View algebra.Ray				// Where did the original view ray come from
	Ray algebra.Ray					// The ray being traced
	Shadow bool						// True if the ray is being used for shadow calculations
	Time algebra.MnFloat			// The time the ray was created, used for animations
	Ior algebra.MnFloat				// Current index of refraction for the volume containing the rays origin
	Depth int						// Current trace depth during recursion
	MaximumDepth int				// Maximum depth for reflections
	MaximumShadowDepth int			// Maximum depth for shadow rays
}

func NewTraceInfo(info shading.RenderInfo, origin algebra.Vector3, direction algebra.Vector3) TraceInfo {
	return TraceInfo{
		View:				algebra.Ray{
			Origin:    origin,
			Direction: direction,
		},
		Ray:                algebra.Ray{
			Origin: origin,
			Direction: direction,
		},
		Rendering: 			info,
		Shadow:             false,
		Time:               0,
		Ior:                DefaultIor,
		Depth:              0,
		MaximumDepth:       DefaultMaximumTraceDepth,
		MaximumShadowDepth: DefaultMaximumShadowDepth,
	}
}

// Creates a TraceInfo instance that can be used to follow a shadow ray through the scene.
func (t *TraceInfo) CreateShadow(ray algebra.Ray) TraceInfo {
	depth := t.Depth + 1
	if !t.Shadow {
		depth = 0
	}

	return TraceInfo{
		Rendering: 			t.Rendering,
		View:				t.View,
		Ray:                ray,
		Shadow:             true,
		Time:               t.Time,
		Ior:                t.Ior,
		Depth:              depth,
		MaximumDepth:       t.MaximumDepth,
		MaximumShadowDepth: t.MaximumShadowDepth,
	}
}

// Creates a new TraceInfo that can be used to follow a reflected path around a specified point. The direction of
// the new TraceInfo is calculated using the supplied surface normal.
func (t *TraceInfo) CreateReflection(point algebra.Vector3, normal algebra.Vector3) TraceInfo {
	return TraceInfo{
		Rendering: 			t.Rendering,
		Ray:                algebra.Ray{
			Origin: point,
			Direction: t.Ray.Direction.Reflect(normal),
		},
		View:				t.View,
		Shadow:             t.Shadow,
		Time:               t.Time,
		Ior:                t.Ior,
		Depth:              t.Depth + 1,
		MaximumDepth:       t.MaximumDepth,
		MaximumShadowDepth: t.MaximumShadowDepth,
	}
}
