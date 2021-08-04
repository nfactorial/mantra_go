package sampling

import "strange-secrets.com/mantra/algebra"

// A pixel sampler provides a single sample passes through the center of a pixel.
type PixelSampler struct {
	CurrentX algebra.MnFloat
	CurrentY algebra.MnFloat
	counter int
}

// Creates a new pixel sampler for the specified pixel
func NewPixelSampler() *PixelSampler {
	return &PixelSampler{
		counter: 0,
	}
}

func (r *PixelSampler) Reset(x int, y int) Sampler2D {
	r.CurrentX = algebra.MnFloat(x)
	r.CurrentY = algebra.MnFloat(y)
	r.counter = 0

	return r
}

// Returns the number of samples the pixel sampler provides
func (p *PixelSampler) Length() int {
	return 1
}

// Returns the X coordinate of the current sample
func (p *PixelSampler) X() algebra.MnFloat {
	return p.CurrentX
}

// Returns the Y coordinate of the current sample
func (p *PixelSampler) Y() algebra.MnFloat {
	return p.CurrentY
}

// Moves to the next sample
func (p *PixelSampler) Next() bool {
	p.counter++

	if p.counter <= 1 {
		return true
	}

	return false
}
