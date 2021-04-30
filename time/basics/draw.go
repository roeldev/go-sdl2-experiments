// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/go-pogo/errors"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/colors"
)

func drawRect(w, h, x, y int32) sdlkit.Renderable {
	rect := sdl.Rect{X: x, Y: y, W: w, H: h}
	color := colors.FireBrick

	return sdlkit.RenderableFunc(func(r *sdl.Renderer) error {
		var err error
		errors.Append(&err,
			r.SetDrawColor(color.R, color.G, color.B, color.A),
			r.DrawRect(&rect),
		)
		return err
	})
}
