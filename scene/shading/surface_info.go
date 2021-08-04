package shading

import "strange-secrets.com/mantra/algebra"

type SurfaceInfo struct {
	Material Material			// The material associated with the surface
	Metallic algebra.MnFloat	// How metallic is the surface [0...1]
	Kr		 algebra.MnFloat	// Coefficient of reflection
	Ka		 algebra.MnFloat	// Coefficient of ambient light
	Kd		 algebra.MnFloat	// Coefficient of diffuse light
	Ks		 algebra.MnFloat	// Coefficient of specular light
}
