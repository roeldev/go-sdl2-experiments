// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/go-pogo/errors"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/geom"
)

type box struct {
	posX, posY float64
	velocity   geom.Vector
	rect       sdl.Rect
	color      sdl.Color
}

func newBox(w, h int32, x, y, vel float64, color sdl.Color) *box {
	return &box{
		posX:     x,
		posY:     y,
		velocity: geom.Vector{X: vel, Y: vel},
		rect:     sdl.Rect{W: w, H: h},
		color:    color,
	}
}

func (box *box) update(dt float64, w, h float64) {
	box.posX += box.velocity.X * dt
	box.posY += box.velocity.Y * dt

	if box.posX <= 0 {
		box.posX *= -1
		box.velocity.X *= -1
	} else {
		mx := w - float64(box.rect.W)
		if box.posX >= mx {
			box.posX -= box.posX - mx
			box.velocity.X *= -1
		}
	}
	if box.posY <= 0 {
		box.posY *= -1
		box.velocity.Y *= -1
	} else {
		my := h - float64(box.rect.H)
		if box.posY >= my {
			box.posY -= box.posY - my
			box.velocity.Y *= -1
		}
	}
}

func (box *box) Render(ren *sdl.Renderer) error {
	box.rect.X = int32(box.posX)
	box.rect.Y = int32(box.posY)

	var err error
	errors.Append(&err,
		ren.SetDrawColor(box.color.R, box.color.G, box.color.B, 0xFF),
		ren.FillRect(&box.rect),
	)
	return err
}
