// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package geom

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Transform struct {
	Rotation, ScaleX, ScaleY float64
}

func ResetTransform(t *Transform) {
	t.Rotation = 0
	t.ScaleX = 1
	t.ScaleY = 1
}

type Bounds struct {
	sdl.Rect
	i32Orig [2]int32
	f64Orig [2]float64
}

func RectBounds(w, h int32) Bounds {
	return Bounds{
		Rect:    sdl.Rect{W: w, H: h},
		i32Orig: [2]int32{w, h},
		f64Orig: [2]float64{float64(w), float64(h)},
	}
}

func (b *Bounds) Update(pos Point) {
	b.X = int32(pos.X - (b.f64Orig[0] / 2))
	b.Y = int32(pos.Y - (b.f64Orig[1] / 2))
	b.W = b.i32Orig[0]
	b.H = b.i32Orig[1]
}

func (b *Bounds) UpdateAndTransform(pos Point, tr Transform) {
	w, h := b.f64Orig[0], b.f64Orig[1]
	if tr.ScaleX != 0 && tr.ScaleX != 1 {
		w *= tr.ScaleX
		b.W = int32(w)
	} else {
		b.W = b.i32Orig[0]
	}
	if tr.ScaleY != 0 && tr.ScaleY != 1 {
		h *= tr.ScaleY
		b.H = int32(h)
	} else {
		b.H = b.i32Orig[1]
	}

	b.X = int32(pos.X - (w / 2))
	b.Y = int32(pos.Y - (h / 2))
}
