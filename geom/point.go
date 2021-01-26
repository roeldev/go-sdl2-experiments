// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package geom

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Point struct {
	X, Y float64
}

// Vector returns a Vector with the same X and Y values as Point.
func (p Point) Vector() Vector { return Vector{X: p.X, Y: p.Y} }

// InCircle returns true when the Point is inside the circle, defined by the
// provided cx, cy and rad values. It calculates the squared distance between
// the Point and cx/cy and compares this with the squared rad value.
func (p Point) InCircle(circle Circle) bool {
	return InCircle(p.X, p.Y, circle.X, circle.Y, circle.Radius)
}

func (p Point) InRectangle(r Rectangle) bool {
	return InRect(p.X, p.Y, r.X, r.Y, r.W, r.H)
}

func (p Point) InRect(r sdl.Rect) bool {
	return InRect(p.X, p.Y,
		float64(r.X), float64(r.Y),
		float64(r.W), float64(r.H),
	)
}

func (p Point) InFRect(r sdl.FRect) bool {
	return InRect(p.X, p.Y,
		float64(r.X), float64(r.Y),
		float64(r.W), float64(r.H),
	)
}

func (p Point) InPolygon(poly Polygon) bool {
	return InPolygon(p.X, p.Y, poly.Edges())
}
