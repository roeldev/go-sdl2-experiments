// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package geom

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type HitTester interface {
	// HitTest returns true when the x and y values are within the HitTester.
	HitTest(x, y float64) bool
	HitTestXY(xy XYGetter) bool
}

type Shape interface {
	XYGetter
	HitTester
	Area() float64
	Bounds() AABB
}

// Ellipse cannot be transformed.
// https://g6auc.me.uk/ellipses/index.html
type Ellipse struct {
	X, Y             float64
	RadiusX, RadiusY float64
}

func (e Ellipse) GetX() float64 { return e.X }
func (e Ellipse) GetY() float64 { return e.Y }
func (e Ellipse) Area() float64 { return math.Pi * e.RadiusX * e.RadiusY }

func (e Ellipse) Bounds() AABB {
	return AABB{
		X: e.X - e.RadiusX,
		Y: e.Y - e.RadiusY,
		W: e.RadiusX * 2,
		H: e.RadiusY * 2,
	}
}

// HitTest returns true when position [x y] is inside the Ellipse.
// https://math.stackexchange.com/questions/76457/check-if-a-point-is-within-an-ellipse
// https://stackoverflow.com/questions/7946187/point-and-ellipse-rotated-position-test-algorithm
func (e Ellipse) HitTest(x, y float64) bool {
	return ((x-e.X)*(x-e.X)/(e.RadiusX*e.RadiusX))+
		((y-e.Y)*(y-e.Y)/(e.RadiusY*e.RadiusY)) <= 1
}

// HitTestXY returns true when the position of the XYGetter is inside Ellipse.
func (e Ellipse) HitTestXY(xy XYGetter) bool { return e.HitTest(xy.GetX(), xy.GetY()) }

// Circle cannot be transformed.
type Circle struct {
	X, Y, Radius float64
}

func (c Circle) GetX() float64 { return c.X }
func (c Circle) GetY() float64 { return c.Y }
func (c Circle) Area() float64 { return math.Pi * c.Radius * c.Radius }

func (c Circle) Bounds() AABB {
	return AABB{
		X: c.X - c.Radius,
		Y: c.Y - c.Radius,
		W: c.Radius * 2,
		H: c.Radius * 2,
	}
}

// HitTest returns true when position [x y] is inside the Circle. It calculates
// the squared distance between [x y] and the Circle's center X Y and compares
// this with the squared Radius.
func (c Circle) HitTest(x, y float64) bool {
	dx, dy := c.X-x, c.Y-y
	return ((dx * dx) + (dy * dy)) < (c.Radius * c.Radius)
}

// HitTestXY returns true when the position of the XYGetter is inside Circle.
func (c Circle) HitTestXY(xy XYGetter) bool { return c.HitTest(xy.GetX(), xy.GetY()) }

// A Rect is an axis-aligned rectangle where X Y is its center point. It cannot
// be transformed using a Transform or Matrix. It can however change position
// and change its width and height.
// Use NewQuad instead to create a four-sided Polygon if you need a rectangular
// shape that needs to be rotated, sheared or transformed in any other way.
type Rect struct {
	X, Y, // center point
	W, H float64
}

func (r Rect) GetX() float64 { return r.X }
func (r Rect) GetY() float64 { return r.Y }
func (r Rect) Area() float64 { return r.W * r.H }

func (r Rect) Bounds() AABB {
	return AABB{
		X: r.X - r.W/2,
		Y: r.Y - r.H/2,
		W: r.W,
		H: r.H,
	}
}

func (r Rect) Rect() *sdl.Rect {
	return &sdl.Rect{
		X: int32(r.X - r.W/2),
		Y: int32(r.Y - r.H/2),
		W: int32(r.W),
		H: int32(r.H),
	}
}

// HitTest returns true when the x and y values are within the Rect.
func (r Rect) HitTest(x, y float64) bool {
	w, h := r.W/2, r.H/2
	return (x >= r.X-w) && (x < r.X+w) && (y >= r.Y-h) && (y < r.Y+h)
}

// HitTestXY returns true when the position of the XYGetter is inside Rect.
func (r Rect) HitTestXY(xy XYGetter) bool { return r.HitTest(xy.GetX(), xy.GetY()) }

// AABB is an axis-aligned bounding box with X and Y indicating its top left
// corner.
type AABB struct {
	X, Y, // top left
	W, H float64
	// https://help.adobe.com/en_US/FlashPlatform/reference/actionscript/3/flash/geom/Rectangle.html#methodSummary
}

func (b AABB) GetX() float64 { return b.X }
func (b AABB) GetY() float64 { return b.Y }

func (b AABB) Rect() *sdl.Rect {
	return &sdl.Rect{X: int32(b.X), Y: int32(b.Y), W: int32(b.W), H: int32(b.H)}
}

// HitTest returns true when the x and y values are within the AABB.
func (b AABB) HitTest(x, y float64) bool {
	return hitTestRect(x, y, b.X, b.Y, b.W, b.H)
}

// HitTestXY returns true when the position of the XYGetter is inside AABB.
func (b AABB) HitTestXY(xy XYGetter) bool {
	return hitTestRect(xy.GetX(), xy.GetY(), b.X, b.Y, b.W, b.H)
}

func hitTestRect(x, y, rX, rY, rW, rH float64) bool {
	return (x >= rX) && (x < rX+rW) && (y >= rY) && (y < rY+rH)
}
