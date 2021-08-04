package render

import "strange-secrets.com/mantra/algebra"

type Target interface {
	Name() string
	Width() int
	Height() int
	Set(x int, y int, c algebra.Vector3)
	SavePngImage(fileName string) error
}
