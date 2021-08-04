package sampling

import (
	"math/rand"
	"strange-secrets.com/mantra/algebra"
	"time"
)

type StochasticSampler struct {
	pixelScale algebra.MnFloat
	xSamples int
	ySamples int
	x int
	y int
	random *rand.Rand
}

func NewStochasticSampler(pixelsWide int, pixelsHigh int, pixelScale float64) StochasticSampler {
	return StochasticSampler{
		pixelScale: pixelScale,
		xSamples: pixelsWide,
		ySamples: pixelsHigh,
		random:   rand.New(rand.NewSource(time.Now().Unix())),
	}
}

func (s *StochasticSampler) Reset() {
	s.x = 0
	s.y = 0
}

func (s *StochasticSampler) Length() int {
	return s.xSamples * s.ySamples
}

func (s *StochasticSampler) NextSample(position algebra.Vector2) (algebra.Vector2, bool) {
	s.x++

	if s.x >= s.xSamples {
		s.x = 0

		s.y++
		if s.y >= s.ySamples {
			s.y = s.ySamples
			return algebra.ZeroVector2, false
		}
	}

	return algebra.Vector2{
		X: position.X + s.random.Float64() * s.pixelScale,
		Y: position.Y + s.random.Float64() * s.pixelScale,
	}, true
}
