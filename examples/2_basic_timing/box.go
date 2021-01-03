// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"image/color"

	"github.com/go-pogo/errors"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/go-pogo/sdlkit"
	"github.com/go-pogo/sdlkit/geom"
)

type box struct {
	pos   sdl.FPoint
	vec   geom.Vector
	rect  sdl.Rect
	color color.RGBA
}

func newBox(w, h int32, x, y, v float32, color color.RGBA) *box {
	return &box{
		pos:   sdl.FPoint{X: x, Y: y},
		vec:   geom.Vector{X: v, Y: v},
		rect:  sdl.Rect{W: w, H: h},
		color: color,
	}
}

func (box *box) update(dt float32, vp sdlkit.Viewport) {
	box.pos.X += box.vec.X * dt
	box.pos.Y += box.vec.Y * dt

	if box.pos.X <= 0 {
		box.pos.X *= -1
		box.vec.X *= -1
	} else {
		mx := float32(vp.W - box.rect.W)
		if box.pos.X >= mx {
			box.pos.X -= box.pos.X - mx
			box.vec.X *= -1
		}
	}
	if box.pos.Y <= 0 {
		box.pos.Y *= -1
		box.vec.Y *= -1
	} else {
		my := float32(vp.H - box.rect.H)
		if box.pos.Y >= my {
			box.pos.Y -= box.pos.Y - my
			box.vec.Y *= -1
		}
	}
}

func (box *box) Draw(r *sdl.Renderer) error {
	box.rect.X = int32(box.pos.X)
	box.rect.Y = int32(box.pos.Y)

	var err error
	errors.Append(&err,
		r.SetDrawColor(box.color.R, box.color.G, box.color.B, 0xFF),
		r.FillRect(&box.rect),
	)
	return err
}
