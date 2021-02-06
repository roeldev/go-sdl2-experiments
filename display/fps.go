// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package display

import (
	"fmt"

	"github.com/go-pogo/errors"
	sdlgfx "github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/go-pogo/sdlkit"
	"github.com/go-pogo/sdlkit/colors"
)

type FpsDisplay struct {
	time *sdlkit.Time

	X, Y, Scale int32
	TextColor   sdl.Color
	ShadowColor sdl.Color
}

func NewFpsDisplay(t *sdlkit.Time, x, y int32) *FpsDisplay {
	return &FpsDisplay{
		time:        t,
		X:           x,
		Y:           y,
		Scale:       2,
		TextColor:   colors.White,
		ShadowColor: sdl.Color{A: 100},
	}
}

func (d *FpsDisplay) Draw(r *sdl.Renderer) (err error) {
	x, y := d.X, d.Y

	var sx, sy float32
	if d.Scale > 1 {
		sx, sy = r.GetScale()
		errors.Append(&err, r.SetScale(float32(d.Scale), float32(d.Scale)))
		x /= d.Scale
		y /= d.Scale
	}

	fps := fmt.Sprintf("%.2f", d.time.Fps())
	sdlgfx.StringColor(r, x+1, x+1, fps, d.ShadowColor) // shadow
	sdlgfx.StringColor(r, x, y, fps, d.TextColor)

	if d.Scale > 1 {
		errors.Append(&err, r.SetScale(sx, sy))
	}
	return err
}
