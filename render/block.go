package render

type EvaluatePixel func(x int, y int)

type Block struct {
	X int
	Y int
	Width int
	Height int
}

func (b *Block) Enumerate(callback EvaluatePixel) {
	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			callback(x + b.X, y + b.Y)
		}
	}
}
