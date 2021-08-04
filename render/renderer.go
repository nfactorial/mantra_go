package render

import (
	"errors"
	"fmt"
	"strange-secrets.com/mantra/algebra"
	"strange-secrets.com/mantra/render/sampling"
	"strange-secrets.com/mantra/scene"
	"strange-secrets.com/mantra/scene/shading"
	"time"
)

const (
	TestImageWidth = 1024 * 2		// 640
	TestImageHeight = 1024			// 480
)

type Renderer struct {
	Images map[string] Target
	Time algebra.MnFloat
}

func NewRenderer() *Renderer {
	return &Renderer{
		Images: make(map[string] Target),
	}
}

// Creates a new image within the renderer which may be used as a target for a render operation.
func (r *Renderer) CreateRenderTarget(name string, width int, height int) (Target, error) {
	if _, ok := r.Images[name]; ok {
		return nil, fmt.Errorf("cannot create render target \"%s\", name already in use", name)
	}

	renderTarget, err := NewGammaTarget(name, width, height)
	if err != nil {
		return nil, err
	}

	r.Images[name] = renderTarget

	return renderTarget, nil
}

func (r *Renderer) Render(imageName string, cameraName string, world *scene.Scene) error {
	cameraNode := world.GetNode(cameraName)
	if cameraNode == nil {
		return fmt.Errorf("cannot find camera \"%s\" for rendering", cameraName)
	}

	if target, ok := r.Images[imageName]; ok {
		return r.renderImage(cameraNode, target, world)
	}

	return fmt.Errorf("cannot render image \"%s\", image could not be found", imageName)
}

func (r *Renderer) renderImage(cameraNode *scene.Node, target Target, world *scene.Scene) error {
	if cameraNode == nil {
		return errors.New("cannot render image without camera node")
	}

	if target == nil {
		return errors.New("cannot render nil image")
	}

	start := time.Now()

	// We intend to use multi-threading for these blocks of pixels
	block := Block{
		X:      0,
		Y:      0,
		Width:  TestImageWidth,
		Height: TestImageHeight,
	}

	r.renderBlock(block, cameraNode, target, world)

	fmt.Println(fmt.Sprintf("render completed %dms", time.Since(start).Milliseconds()))
	return nil
}

func (r *Renderer) renderImageAsync(cameraNode *scene.Node, target Target, world *scene.Scene) error {
	if cameraNode == nil {
		return errors.New("cannot render image without camera node")
	}

	if target == nil {
		return errors.New("cannot render nil image")
	}

	lbc := NewLinearBlockChannel(target, 16, 16)

	start := time.Now()

	for loop := 0; loop < 4; loop++ {
		go r.processBlocks(lbc, cameraNode, target, world)
	}

	fmt.Println(fmt.Sprintf("render completed %dms", time.Since(start).Milliseconds()))
	return nil
}

func (r *Renderer) processBlocks(bc BlockChannel, cameraNode *scene.Node, target Target, world *scene.Scene) {
	for block := bc.NextBlock(); block.Width != 0; block = bc.NextBlock() {
		r.renderBlock(block, cameraNode, target, world)
	}
}

func (r *Renderer) renderBlock(block Block, cameraNode *scene.Node, target Target, world *scene.Scene) {
	//sampler := sampling.NewPixelSampler()
	sampler := sampling.NewRandomSampler(8)

	for y := 0; y < block.Height; y++ {
		for x := 0; x < block.Width; x++ {
			sampler.Reset(x + block.X, y + block.Y)
			target.Set(x + block.X, y + block.Y, r.renderPixel(sampler.Reset(x + block.X, y + block.Y), cameraNode, world))
		}
	}
}

func (r *Renderer) renderPixel(pixel algebra.Vector2, sampler sampling.Sampler2D, cameraNode *scene.Node, world *scene.Scene) algebra.Vector3 {
	halfWidth := float64(TestImageWidth) / 2.0
	halfHeight := float64(TestImageHeight) / 2.0

	renderInfo := shading.RenderInfo{
		Pixel: algebra.ZeroVector2,
		Width:  TestImageWidth,
		Height: TestImageHeight,
		Aspect: float64(TestImageWidth) / float64(TestImageHeight),
	}

	result := algebra.ZeroVector3

	valid := true

	// Here we should split the pixel up into multiple rays for anti-aliasing
	for renderInfo.Pixel, valid = sampler.Next(pixel); valid {
		renderInfo.Pixel.X = (renderInfo.Pixel.X - halfWidth) / halfWidth
		renderInfo.Pixel.Y = -(renderInfo.Pixel.Y - halfHeight) / halfHeight

		direction := cameraNode.CastRay(renderInfo.Pixel.X, renderInfo.Pixel.Y, renderInfo.Aspect)
		traceInfo := scene.NewTraceInfo(renderInfo, cameraNode.Location, direction)

		// TODO: Perhaps ShadeRay should return a color?
		result = result.Add(world.ShadeRay(traceInfo))
	}

	return result.DivideScalar(algebra.MnFloat(sampler.Length()))
}
