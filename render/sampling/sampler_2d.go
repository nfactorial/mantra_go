package sampling

import "strange-secrets.com/mantra/algebra"

type Sampler2D interface {
	// Resets the sampler for generation with a new pixel
	Reset() Sampler2D

	// Retrieves the next sample, if there are no more samples this method returns false.
	Next(pixel algebra.Vector2) (algebra.Vector2, bool)

	// Returns the total number of samples provided by the sampler.
	Length() int
}
