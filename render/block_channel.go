package render

type BlockChannel interface {
	NextBlock() Block
}
