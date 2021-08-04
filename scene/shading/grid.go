package shading

import (
	"math"
	"strange-secrets.com/mantra/algebra"
)

var (
	defaultGridCellColor = algebra.Vector3{1,1,1}
	defaultGridLineColor = algebra.Vector3{0,0,0}
)

const (
	GridMaterialName = "grid"

	defaultGridCellWidth = 2.5
	defaultGridLineWidth = 0.05
)

type GridMaterial struct {
	Color		algebra.Vector3
	LineColor	algebra.Vector3
	LineWidth	algebra.MnFloat
	CellWidth   algebra.MnFloat
}

func NewGridMaterial() Material {
	return &GridMaterial{
		Color: defaultGridCellColor,
		LineColor: defaultGridLineColor,
		CellWidth: defaultGridCellWidth,
		LineWidth: defaultGridLineWidth,
	}
}

func (g *GridMaterial) Evaluate(shadeInfo ShadeInfo) algebra.Vector3 {
	x := algebra.SmoothStep(0.0, g.LineWidth, math.Mod(math.Abs(shadeInfo.Point.X / g.CellWidth), 1.0))
	z := algebra.SmoothStep(0.0, g.LineWidth, math.Mod(math.Abs(shadeInfo.Point.Z / g.CellWidth), 1.0))

	return g.LineColor.Lerp(g.Color, x * z)
}
