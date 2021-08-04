package algebra

import "math"

type BoundingBox struct {
	Minimum Vector3
	Maximum Vector3
}

var InvalidBoundingBox = BoundingBox{
	Minimum: Vector3{math.Inf(1), math.Inf(1), math.Inf(1)},
	Maximum: Vector3{math.Inf(-1), math.Inf(-1), math.Inf(-1)},
}

var InfiniteBoundingBox = BoundingBox{
	Minimum: Vector3{math.Inf(-1), math.Inf(-1), math.Inf(-1)},
	Maximum: Vector3{math.Inf(1), math.Inf(1), math.Inf(1)},
}

// Creates a new BoundingBox that contains a single point in 3d space
func NewBoundingBox(origin Vector3) BoundingBox {
	return BoundingBox{origin, origin}
}

// Expands the bounding box to contain the specified point in 3D space
func (b *BoundingBox) Include(position Vector3) {
	b.Minimum = b.Minimum.Min(position)
	b.Maximum = b.Maximum.Max(position)
}

// Determines whether or not the specified position is contained within the bounding box
func (b BoundingBox) Contains(position Vector3) bool {
	if position.X > b.Minimum.X && position.X < b.Maximum.X {
		if position.Y > b.Minimum.Y && position.Y < b.Maximum.Y {
			if position.Z > b.Minimum.Z && position.Z < b.Maximum.Z {
				return true
			}
		}
	}

	return false
}
