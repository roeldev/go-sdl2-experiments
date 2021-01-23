// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdlkit

import (
	"github.com/go-pogo/errors"
	sdlgfx "github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

type Canvas struct {
	renderer *sdl.Renderer
	target   *sdl.Texture
	texture  *sdl.Texture
	errors   []error

	fillColor sdl.Color
	lineColor sdl.Color
	lineStyle [1]int32 // thickness

	antiAlias bool
	blendMode sdl.BlendMode

	fill bool
	line bool
}

func NewCanvas(renderer *sdl.Renderer) *Canvas {
	return &Canvas{renderer: renderer}
}

func (c *Canvas) catchErr(err ...error) {
	if len(err) == 1 && err[0] == nil {
		return
	}

	c.errors = append(c.errors, err...)
}

func (c *Canvas) CreateTexture(format uint32, access int, w, h int32) (*sdl.Texture, error) {
	tx, err := c.renderer.CreateTexture(format, access|sdl.TEXTUREACCESS_TARGET, w, h)
	if err != nil {
		return nil, err
	}

	c.target = c.renderer.GetRenderTarget()
	if err = c.renderer.SetRenderTarget(tx); err != nil {
		c.target = nil
		return nil, err
	}
	if err = tx.SetBlendMode(sdl.BLENDMODE_BLEND); err != nil {
		return nil, err
	}

	c.texture = tx
	return tx, nil
}

func (c *Canvas) AntiAlias(aa bool) {
	c.antiAlias = aa
}

func (c *Canvas) BlendMode(mode sdl.BlendMode) {
	c.blendMode = mode

	if c.texture != nil {
		c.catchErr(c.texture.SetBlendMode(mode))
	}
}

func (c *Canvas) BeginFill(color sdl.Color) {
	if c.blendMode == sdl.BLENDMODE_NONE {
		color.A = 0xFF
	}

	c.fillColor = color
	c.fill = true
}

func (c *Canvas) EndFill() { c.fill = false }

func (c *Canvas) LineStyle(thickness int32, color sdl.Color) {
	c.lineColor = color
	c.lineStyle[0] = thickness
	c.line = true
}

func (c *Canvas) EndLineStyle() { c.line = false }

func (c *Canvas) DrawPixel(x, y, size int32) {
	if size < 2 {
		c.catchErr(
			c.renderer.SetDrawColor(c.fillColor.R, c.fillColor.G, c.fillColor.B, c.fillColor.A),
			c.renderer.DrawPoint(x, y),
		)
		return
	}

	c.catchErr(
		c.renderer.SetDrawColor(c.fillColor.R, c.fillColor.G, c.fillColor.B, c.fillColor.A),
		c.renderer.FillRect(&sdl.Rect{
			X: x - size/2,
			Y: y - size/2,
			W: size,
			H: size,
		}),
	)
}

func (c *Canvas) DrawLine(x1, y1, x2, y2 int32) {
	if !c.line {
		return
	}
	if c.lineStyle[0] > 1 {
		sdlgfx.ThickLineColor(c.renderer, x1, y1, x2, y2, c.lineStyle[0], c.lineColor)
	} else if c.antiAlias {
		sdlgfx.AALineColor(c.renderer, x1, y1, x2, y2, c.lineColor)
	} else {
		sdlgfx.LineColor(c.renderer, x1, y1, x2, y2, c.lineColor)
	}
}

func (c *Canvas) DrawEllipse(x, y, rx, ry int32) {
	if c.fill {
		sdlgfx.FilledEllipseColor(c.renderer, x, y, rx, ry, c.fillColor)
	}
	if c.line && c.antiAlias {
		sdlgfx.AAEllipseColor(c.renderer, x, y, rx, ry, c.lineColor)
	} else if c.line {
		sdlgfx.EllipseColor(c.renderer, x, y, rx, ry, c.lineColor)
	}
}

func (c *Canvas) DrawCircle(x, y, rad int32) { c.DrawEllipse(x, y, rad, rad) }

func (c *Canvas) DrawRect(x, y, w, h int32) {
	rect := &sdl.Rect{X: x, Y: y, W: w, H: h}
	if c.fill {
		c.catchErr(
			c.renderer.SetDrawColor(c.fillColor.R, c.fillColor.G, c.fillColor.B, c.fillColor.A),
			c.renderer.FillRect(rect),
		)
	}
	if c.line {
		c.catchErr(
			c.renderer.SetDrawColor(c.lineColor.R, c.lineColor.G, c.lineColor.B, c.lineColor.A),
			c.renderer.DrawRect(rect),
		)
	}
}

func (c *Canvas) DrawSquare(x, y, size int32) { c.DrawRect(x, y, size, size) }

func (c *Canvas) DrawRoundRect(x, y, w, h, rad int32) {
	x2, y2 := x+w, y+h
	if c.fill {
		sdlgfx.RoundedBoxColor(c.renderer, x, y, x2, y2, rad, c.fillColor)
	}
	if c.line {
		sdlgfx.RoundedRectangleColor(c.renderer, x, y, x2, y2, rad, c.lineColor)
	}
}

func ConvPolygonPoints(points []sdl.Point, dx, dy int16) ([]int16, []int16) {
	vx := make([]int16, len(points))
	vy := make([]int16, len(points))
	for i, pt := range points {
		vx[i] = int16(pt.X) + dx
		vy[i] = int16(pt.Y) + dy
	}
	return vx, vy
}

func (c *Canvas) DrawPolygon(vx, vy []int16) {
	if c.fill {
		sdlgfx.FilledPolygonColor(c.renderer, vx, vy, c.fillColor)
	}
	if c.line && c.antiAlias {
		sdlgfx.AAPolygonColor(c.renderer, vx, vy, c.lineColor)
	} else if c.line {
		sdlgfx.PolygonColor(c.renderer, vx, vy, c.lineColor)
	}
}

func (c *Canvas) Done() (err error) {
	if c.texture != nil {
		_ = c.renderer.SetRenderTarget(c.target)
		c.target = nil
		c.texture = nil
	}

	if len(c.errors) != 0 {
		err = errors.Combine(c.errors...)
		c.errors = c.errors[:0]
	}
	return
}
