package colors

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/math"
)

func AlphaMod(color sdl.Color, alpha float64) sdl.Color {
	color.A = uint8(math.Clamp(float64(255)*alpha, 0, 255))
	return color
}
