// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package geom

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Point sdl.FPoint

// Vector returns a Vector with the same X and Y values as the Point.
func (p Point) Vector() Vector { return Vector{X: p.X, Y: p.Y} }

// InCircle returns true when x/y is inside the circle, defined by the provided
// cx, cy and rad values. It calculates the squared distance between the x/y
// and cx/cy values and compares this with the squared rad value.
func InCircle(x, y, cx, cy, rad float32) bool {
	dx, dy := cx-x, cy-y
	return ((dx * dx) + (dy * dy)) < (rad * rad)
}

// InCircle returns true when the Point is inside the circle, defined by the
// provided cx, cy and rad values. It calculates the squared distance between
// the Point and cx/cy and compares this with the squared rad value.
func (p Point) InCircle(cx, cy, rad float32) bool {
	return InCircle(p.X, p.Y, cx, cy, rad)
}

func InRect(x, y, rx, ry, rw, rh float32) bool {
	return (x >= rx) && (x < rx+rw) &&
		(y >= ry) && (y < ry+rh)
}

func (p Point) InRect(r sdl.Rect) bool {
	return InRect(p.X, p.Y, float32(r.X), float32(r.Y), float32(r.W), float32(r.H))
}

func (p Point) InFRect(r sdl.FRect) bool {
	return InRect(p.X, p.Y, r.X, r.Y, r.W, r.H)
}
