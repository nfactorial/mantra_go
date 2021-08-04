package render

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"strange-secrets.com/mantra/algebra"
)

const (
	MaximumImageWidth	= 4096
	MaximumImageHeight	= 4096
	DefaultGamma		= 2.2
)

type GammaTarget struct {
	pixelMap	*image.RGBA
	name		string
	gamma		algebra.MnFloat
	width		int
	height		int
}

func NewGammaTarget(name string, width int, height int) (*GammaTarget, error) {
	if width <= 0 || width > MaximumImageWidth {
		return nil, fmt.Errorf("cannot create render target of width %d", width)
	}

	if height <= 0 || height > MaximumImageHeight {
		return nil, fmt.Errorf("cannot create render target of height %d", height)
	}

	return &GammaTarget{
		pixelMap: image.NewRGBA(image.Rect(0, 0, width, height)),
		name: name,
		gamma: DefaultGamma,
		width: width,
		height: height,
	}, nil
}

// Retrieves the name associated with the render target.
func (g *GammaTarget) Name() string {
	return g.name
}

// Retrieves the width of the render target (in pixels)
func (g *GammaTarget) Width() int {
	return g.width
}

// Retrieves the height of the render target (in pixels)
func (g *GammaTarget) Height() int {
	return g.height
}

// Stores a specified colour at a specified pixel within the render target
func (g *GammaTarget) Set(x int, y int, c algebra.Vector3) {
	g.pixelMap.Set(x, y, c.Pow(g.gamma).Saturate().ToColor())
}

func (g *GammaTarget) SavePngImage(fileName string) error {
	if g.pixelMap == nil {
		return fmt.Errorf("cannot save image, missing pixel map")
	}

	f, err := os.Create(fileName)
	if err != nil {
		return err
	}

	if err := png.Encode(f, g.pixelMap); err != nil {
		_ = f.Close()
		return err
	}

	return f.Close()
}
