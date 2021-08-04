package sampling

import (
	"math/rand"
	"strange-secrets.com/mantra/algebra"
	"time"
)

type RandomSampler struct {
	CurrentX	algebra.MnFloat
	CurrentY	algebra.MnFloat
	x			algebra.MnFloat
	y			algebra.MnFloat
	random		*rand.Rand
	samples		int
	counter		int
}

func NewRandomSampler(samples int) *RandomSampler {
	return &RandomSampler{
		random:   rand.New(rand.NewSource(time.Now().Unix())),
		samples:  samples,
		counter:  0,
	}
}

func (r *RandomSampler) Reset(x int, y int) Sampler2D {
	r.CurrentX = algebra.MnFloat(x)
	r.CurrentY = algebra.MnFloat(y)
	r.x = algebra.MnFloat(x)
	r.y = algebra.MnFloat(y)
	r.counter = 0

	return r
}

func (r *RandomSampler) Length() int {
	return r.samples
}

func (r *RandomSampler) X() algebra.MnFloat {
	return r.CurrentX
}

func (r *RandomSampler) Y() algebra.MnFloat {
	return r.CurrentY
}

func (r *RandomSampler) Next() bool {
	r.counter++

	if r.counter <= r.samples {
		r.CurrentX = r.x + r.random.Float64()
		r.CurrentY = r.y + r.random.Float64()
		return true
	}

	return false
}
