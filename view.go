// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdlkit

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Viewport sdl.Rect

func (vp Viewport) Rect() sdl.Rect {
	return sdl.Rect{X: vp.X, Y: vp.Y, W: vp.W, H: vp.H}
}

func (vp Viewport) Center() sdl.Point {
	return sdl.Point{X: vp.X + vp.W/2, Y: vp.Y + vp.H/2}
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
