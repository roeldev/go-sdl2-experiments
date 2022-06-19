package colors

import (
	"github.com/veandco/go-sdl2/sdl"
)

type RgbaColor sdl.Color

func (color RgbaColor) RGBA() (r, g, b, a uint32) {
	return uint32(color.R), uint32(color.G), uint32(color.B), uint32(color.A)
}
