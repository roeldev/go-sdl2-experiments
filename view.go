package sdlkit

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Viewport sdl.Rect

func (vp Viewport) Rect() sdl.Rect {
	return sdl.Rect{X: vp.X, Y: vp.Y, W: vp.W, H: vp.H}
}

func (vp Viewport) Border(margin int32) (t, b, l, r int32) {
	return vp.Y + margin,
		vp.H - margin,
		vp.X + margin,
		vp.W - margin
}

func (vp Viewport) FBorder(margin float64) (t, b, l, r float64) {
	return float64(vp.Y) + margin,
		float64(vp.H) - margin,
		float64(vp.X) + margin,
		float64(vp.W) - margin
}
