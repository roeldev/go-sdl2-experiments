// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"math"

	"github.com/go-pogo/errors"
	sdlgfx "github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit"
)

func drawCircle(rad, x, y int32, color sdl.Color) sdlkit.Renderable {
	col := color
	return sdlkit.RenderableFunc(func(r *sdl.Renderer) error {
		sdlgfx.FilledCircleColor(r, x, y, rad, col)
		return nil
	})
}

func drawRect(w, h, x, y int32, color sdl.Color) sdlkit.Renderable {
	rect := sdl.Rect{X: x, Y: y, W: w, H: h}
	return sdlkit.RenderableFunc(func(r *sdl.Renderer) error {
		var err error
		errors.Append(&err,
			r.SetDrawColor(color.R, color.G, color.B, 0xFF),
			r.FillRect(&rect),
		)
		return err
	})
}

func drawSquare(size, x, y int32, color sdl.Color) sdlkit.Renderable {
	return drawRect(size, size, x, y, color)
}

// this is basically a sprite
func drawImg(tx *sdl.Texture, scale float32, x, y int32, flip sdl.RendererFlip) sdlkit.Renderable {
	_, _, w, h, _ := tx.Query()
	if scale != 1 {
		w = int32(math.RoundToEven(float64(float32(w) * scale)))
		h = int32(math.RoundToEven(float64(float32(h) * scale)))
	}

	rect := &sdl.Rect{X: x, Y: y, W: w, H: h}
	return sdlkit.RenderableFunc(func(r *sdl.Renderer) error {
		if flip == sdl.FLIP_NONE {
			return r.Copy(tx, nil, rect)
		} else {
			return r.CopyEx(tx, nil, rect, 0, nil, flip)
		}
	})
}
