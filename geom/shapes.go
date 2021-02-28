// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package geom

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type Shape interface {
	XY
	Area() float64
}

// Circle cannot be transformed.
type Circle struct {
	X, Y, Radius float64
}

func (c Circle) GetX() float64   { return c.X }
func (c Circle) GetY() float64   { return c.Y }
func (c *Circle) SetX(x float64) { c.X = x }
func (c *Circle) SetY(y float64) { c.Y = y }

func (c Circle) Area() float64 { return math.Pi * c.Radius * c.Radius }

// Ellipse cannot be transformed.
// https://g6auc.me.uk/ellipses/index.html
type Ellipse struct {
	X, Y             float64
	RadiusX, RadiusY float64
}

func (e Ellipse) GetX() float64   { return e.X }
func (e Ellipse) GetY() float64   { return e.Y }
func (e *Ellipse) SetX(x float64) { e.X = x }
func (e *Ellipse) SetY(y float64) { e.Y = y }

func (e Ellipse) Area() float64 { return math.Pi * e.RadiusX * e.RadiusY }

// A Rect is an axis-aligned rectangle where X Y is its center point. It cannot
// be transformed using a Transform or Matrix. It can however change position
// and change its width and height.
// Use NewQuad instead to create a four-sided polygon if you need a rectangular
// shape that needs to be rotated, sheared or transformed in any other way.
type Rect struct {
	X, Y, // center point
	W, H float64
}

func (r Rect) GetX() float64   { return r.X }
func (r Rect) GetY() float64   { return r.Y }
func (r *Rect) SetX(x float64) { r.X = x }
func (r *Rect) SetY(y float64) { r.Y = y }

func (r Rect) Area() float64 { return r.W * r.H }

func (r Rect) Rect() *sdl.Rect {
	return &sdl.Rect{
		X: int32(r.X - r.W/2),
		Y: int32(r.Y - r.H/2),
		W: int32(r.W),
		H: int32(r.H),
	}
}
