package render

type LinearBlockChannel struct {
	channel chan Block
}

func minInt(a int, b int) int {
	if b < a {
		return b
	}

	return a
}

func NewLinearBlockChannel(target Target, blockWidth int, blockHeight int) *LinearBlockChannel {
	lbc := &LinearBlockChannel{
		channel: make(chan Block),
	}

	imageWidth := target.Width()
	imageHeight := target.Height()

	for y := 0; y < imageHeight; y += blockHeight {
		for x := 0; x < imageWidth; x += blockWidth {
			lbc.channel <- Block{
				X: x,
				Y: y,
				Width: minInt(blockWidth, imageWidth - x),
				Height: minInt(blockHeight, imageHeight - y),
			}
		}
	}

	close(lbc.channel)

	return lbc
}

func (l *LinearBlockChannel) NextBlock() Block {
	return <-l.channel
}
