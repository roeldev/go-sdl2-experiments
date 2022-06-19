// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdlkit

import (
	"github.com/go-pogo/errors"
	sdlgfx "github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/colors"
)

type Drawable interface {
	Draw(canvas *Canvas)
}

type DrawableFunc func(canvas *Canvas)

func (fn DrawableFunc) Draw(canvas *Canvas) { fn(canvas) }

type Canvas struct {
	engine  *sdl.Renderer
	camera  *Camera
	target  *sdl.Texture
	texture *sdl.Texture
	errors  []error

	fillColor sdl.Color
	lineColor sdl.Color
	lineStyle [1]int32 // thickness

	antiAlias bool
	blendMode sdl.BlendMode

	fill bool
	line bool
}

func NewCanvas(engine *sdl.Renderer) *Canvas {
	return &Canvas{
		engine: engine,
	}
}

func (c *Canvas) catchErr(err ...error) {
	if len(err) == 1 && err[0] == nil {
		return
	}

	c.errors = append(c.errors, err...)
}

func (c *Canvas) Renderer() *sdl.Renderer { return c.engine }

func (c *Canvas) Render(r Renderable) { c.catchErr(r.Render(c.engine)) }

func (c *Canvas) Camera() *Camera { return c.camera }

func (c *Canvas) SetCamera(cam *Camera) { c.camera = cam }

func (c *Canvas) CreateTexture(format uint32, access int, w, h int32) (*sdl.Texture, error) {
	tx, err := c.engine.CreateTexture(format, access|sdl.TEXTUREACCESS_TARGET, w, h)
	if err != nil {
		return nil, err
	}

	c.target = c.engine.GetRenderTarget()
	if err = c.engine.SetRenderTarget(tx); err != nil {
		c.target = nil
		return nil, err
	}
	if err = tx.SetBlendMode(c.blendMode); err != nil {
		return nil, err
	}

	c.texture = tx
	return tx, nil
}

func (c *Canvas) CreateTextureClip(format uint32, access int, w, h int32) (TextureClip, error) {
	tx, err := c.CreateTexture(format, access, w, h)
	return TextureClip{
		Texture:  tx,
		Location: sdl.Rect{W: w, H: h},
	}, errors.Trace(err)
}

func (c *Canvas) SetDrawAntiAlias(aa bool) { c.antiAlias = aa }

func (c *Canvas) SetDrawBlendMode(mode sdl.BlendMode) {
	if err := c.engine.SetDrawBlendMode(mode); err != nil {
		c.catchErr(err)
	} else {
		c.blendMode = mode
	}

	if c.texture != nil {
		c.catchErr(c.texture.SetBlendMode(mode))
	}
}

func (c *Canvas) GetFill() (sdl.Color, bool) {
	return c.fillColor, c.fill
}

func (c *Canvas) BeginFill(color sdl.Color) {
	if c.blendMode == sdl.BLENDMODE_NONE {
		color.A = 0xFF
	}

	c.fillColor = color
	c.fill = true
}

func (c *Canvas) BeginFillAlpha(color sdl.Color, alpha uint8) {
	if alpha < 0xFF {
		if err := c.engine.SetDrawBlendMode(sdl.BLENDMODE_BLEND); err != nil {
			c.catchErr(err)
		} else {
			color.A = alpha
		}
	}

	c.fillColor = color
	c.fill = true
}

func (c *Canvas) BeginFillRGBA(r, g, b, a uint8) {
	if c.blendMode == sdl.BLENDMODE_NONE {
		a = 0xFF
	}

	c.fillColor.R = r
	c.fillColor.G = g
	c.fillColor.B = b
	c.fillColor.A = a
	c.fill = true
}

func (c *Canvas) EndFill() { c.fill = false }

func (c *Canvas) GetLineStyle() (int32, sdl.Color, bool) {
	return c.lineStyle[0], c.lineColor, c.line
}

func (c *Canvas) BeginLineStyle(thickness int32, color sdl.Color) {
	c.lineColor = color
	c.lineStyle[0] = thickness
	c.line = true
}

func (c *Canvas) EndLineStyle() { c.line = false }

func (c *Canvas) Draw(d Drawable) { d.Draw(c) }

func (c *Canvas) DrawPixel(x, y, size int32) {
	if c.camera != nil && !c.camera.disabled {
		x = c.camera.TranslateX(x)
		y = c.camera.TranslateY(y)
	}

	if size < 2 {
		c.catchErr(
			c.engine.SetDrawColor(c.fillColor.R, c.fillColor.G, c.fillColor.B, c.fillColor.A),
			c.engine.DrawPoint(x, y),
		)
		return
	}

	c.catchErr(
		c.engine.SetDrawColor(c.fillColor.R, c.fillColor.G, c.fillColor.B, c.fillColor.A),
		c.engine.FillRect(&sdl.Rect{
			X: x - size/2,
			Y: y - size/2,
			W: size,
			H: size,
		}),
	)
}

func (c *Canvas) DrawPixelF(x, y float64, size int32) {
	c.DrawPixel(int32(x), int32(y), size)
}

func (c *Canvas) DrawSdlPoint(pt sdl.Point, size int32) {
	c.DrawPixel(pt.X, pt.Y, size)
}

func (c *Canvas) DrawSdlFPoint(pt sdl.FPoint, size int32) {
	c.DrawPixel(int32(pt.X), int32(pt.Y), size)
}

func (c *Canvas) DrawLine(x1, y1, x2, y2 int32) {
	if !c.line {
		return
	}

	if c.camera != nil && !c.camera.disabled {
		x1 = c.camera.TranslateX(x1)
		y1 = c.camera.TranslateY(y1)
		x2 = c.camera.TranslateX(x2)
		y2 = c.camera.TranslateY(y2)
	}

	if c.lineStyle[0] > 1 {
		sdlgfx.ThickLineColor(c.engine, x1, y1, x2, y2, c.lineStyle[0], c.lineColor)
	} else if c.antiAlias {
		sdlgfx.AALineColor(c.engine, x1, y1, x2, y2, c.lineColor)
	} else {
		sdlgfx.LineColor(c.engine, x1, y1, x2, y2, c.lineColor)
	}
}

func (c *Canvas) DrawLineF(x1, y1, x2, y2 float64) {
	c.DrawLine(int32(x1), int32(y1), int32(x2), int32(y2))
}

func (c *Canvas) DrawEllipse(x, y, radX, radY int32) {
	if c.camera != nil && !c.camera.disabled {
		x = c.camera.TranslateX(x)
		y = c.camera.TranslateY(y)
	}

	if c.fill {
		sdlgfx.FilledEllipseColor(c.engine, x, y, radX, radY, c.fillColor)
	}
	if c.line && c.antiAlias {
		sdlgfx.AAEllipseColor(c.engine, x, y, radX, radY, c.lineColor)
	} else if c.line {
		sdlgfx.EllipseColor(c.engine, x, y, radX, radY, c.lineColor)
	}
}

func (c *Canvas) DrawEllipseF(x, y, rx, ry float64) {
	c.DrawEllipse(int32(x), int32(y), int32(rx), int32(ry))
}

func (c *Canvas) DrawCircle(x, y, rad int32) { c.DrawEllipse(x, y, rad, rad) }

func (c *Canvas) DrawCircleF(x, y, rad float64) {
	c.DrawEllipse(int32(x), int32(y), int32(rad), int32(rad))
}

func (c *Canvas) DrawRect(x, y, w, h int32) {
	c.DrawSdlRect(sdl.Rect{X: x, Y: y, W: w, H: h})
}

func (c *Canvas) DrawRectF(x, y, w, h float64) {
	c.DrawSdlRect(sdl.Rect{X: int32(x), Y: int32(y), W: int32(w), H: int32(h)})
}

func (c *Canvas) DrawSdlRect(rect sdl.Rect) {
	if c.camera != nil && !c.camera.disabled {
		rect.X = c.camera.TranslateX(rect.X)
		rect.Y = c.camera.TranslateY(rect.Y)
	}

	if c.fill {
		c.catchErr(
			c.engine.SetDrawColor(c.fillColor.R, c.fillColor.G, c.fillColor.B, c.fillColor.A),
			c.engine.FillRect(&rect),
		)
	}
	if c.line {
		c.catchErr(
			c.engine.SetDrawColor(c.lineColor.R, c.lineColor.G, c.lineColor.B, c.lineColor.A),
			c.engine.DrawRect(&rect),
		)
	}
}

func (c *Canvas) DrawSdlFRect(rect sdl.FRect) {
	if c.camera != nil && !c.camera.disabled {
		rect.X += float32(c.camera.TranslateXF(0))
		rect.Y += float32(c.camera.TranslateYF(0))
	}

	if c.fill {
		c.catchErr(
			c.engine.SetDrawColor(c.fillColor.R, c.fillColor.G, c.fillColor.B, c.fillColor.A),
			c.engine.FillRectF(&rect),
		)
	}
	if c.line {
		c.catchErr(
			c.engine.SetDrawColor(c.lineColor.R, c.lineColor.G, c.lineColor.B, c.lineColor.A),
			c.engine.DrawRectF(&rect),
		)
	}
}

func (c *Canvas) DrawRoundRect(x, y, w, h, rad int32) {
	if c.camera != nil && !c.camera.disabled {
		x = c.camera.TranslateX(x)
		y = c.camera.TranslateY(y)
	}

	x2, y2 := x+w, y+h
	if c.fill {
		sdlgfx.RoundedBoxColor(c.engine, x, y, x2, y2, rad, c.fillColor)
	}
	if c.line {
		sdlgfx.RoundedRectangleColor(c.engine, x, y, x2, y2, rad, c.lineColor)
	}
}

func (c *Canvas) DrawRoundRectF(x, y, w, h, rad float64) {
	c.DrawRoundRect(int32(x), int32(y), int32(w), int32(h), int32(rad))
}

func (c *Canvas) DrawPolygon(vx, vy []int16) {
	if c.camera != nil && !c.camera.disabled {
		tx, ty := int16(c.camera.TranslateX(0)), int16(c.camera.TranslateY(0))
		for i := 0; i < len(vx); i++ {
			vx[i] += tx
			vy[i] += ty
		}
	}

	if c.fill {
		sdlgfx.FilledPolygonColor(c.engine, vx, vy, c.fillColor)
	}
	if c.line && c.antiAlias {
		sdlgfx.AAPolygonColor(c.engine, vx, vy, c.lineColor)
	} else if c.line {
		sdlgfx.PolygonColor(c.engine, vx, vy, c.lineColor)
	}
}

func (c *Canvas) DrawTexture(tx *sdl.Texture, src *sdl.Rect, dest sdl.Rect) {
	if c.camera != nil && !c.camera.disabled {
		dest.X = c.camera.TranslateX(dest.X)
		dest.Y = c.camera.TranslateY(dest.Y)
	}

	c.catchErr(c.engine.Copy(tx, src, &dest))
	c.catchErr(
		c.engine.SetDrawColor(colors.Black.R, colors.Black.G, colors.Black.B, 10),
		c.engine.DrawRect(&dest),
	)
}

func (c *Canvas) DrawTextureEx(tx *sdl.Texture, src *sdl.Rect, dest sdl.Rect, deg float64, origin sdl.Point, flip sdl.RendererFlip) {
	if c.camera != nil && !c.camera.disabled {
		dest.X = c.camera.TranslateX(dest.X)
		dest.Y = c.camera.TranslateY(dest.Y)
	}

	c.catchErr(c.engine.CopyEx(tx, src, &dest, deg, &origin, flip))
}

func (c *Canvas) DrawTextureClip(clip TextureClip, dest sdl.Rect) {
	c.DrawTexture(clip.Texture, &clip.Location, dest)
}

func (c *Canvas) DrawTextureClipEx(clip TextureClip, dest sdl.Rect, deg float64, origin sdl.Point, flip sdl.RendererFlip) {
	c.DrawTextureEx(clip.Texture, &clip.Location, dest, deg, origin, flip)
}

func (c *Canvas) Done() (err error) {
	if c.texture != nil {
		c.catchErr(c.engine.SetRenderTarget(c.target))
		c.target = nil
		c.texture = nil
	}

	if len(c.errors) != 0 {
		err = errors.Combine(c.errors...)
		c.errors = c.errors[:0]
	}
	return
}
