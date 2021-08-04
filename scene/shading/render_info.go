package shading

import "strange-secrets.com/mantra/algebra"

type RenderInfo struct {
	Pixel		algebra.Vector2			// Coordinates of the pixel currently being rendered
	Width		int						// Width of the render target (in pixels)
	Height		int						// Height of the render target (in pixels)
	Aspect		algebra.MnFloat			// Aspect ratio of the image being rendered
}
