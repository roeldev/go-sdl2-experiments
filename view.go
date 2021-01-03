package sdlkit

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Viewport sdl.Rect

func (vp Viewport) Rect() sdl.Rect {
	return sdl.Rect{X: vp.X, Y: vp.Y, W: vp.W, H: vp.H}
}

// func (vp Viewport) FRect() sdl.FRect {
// 	return sdl.FRect{
// 		X: float32(vp.X),
// 		Y: float32(vp.Y),
// 		W: float32(vp.W),
// 		H: float32(vp.H),
// 	}
// }

func (vp Viewport) Border(margin int32) (t, b, l, r int32) {
	return vp.Y + margin,
		vp.H - margin,
		vp.X + margin,
		vp.W - margin
}

func (vp Viewport) FBorder(margin float32) (t, b, l, r float32) {
	return float32(vp.Y) + margin,
		float32(vp.H) - margin,
		float32(vp.X) + margin,
		float32(vp.W) - margin
}
