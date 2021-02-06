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

// PointFromInt creates a new Point from int32 values.
func PointFromInt(x, y int32) *Point {
	return &Point{X: float64(x), Y: float64(y)}
}

// PointFromXY creates a new Point from a XYGetter.
func PointFromXY(xy XYGetter) *Point {
	return &Point{X: xy.GetX(), Y: xy.GetY()}
}

func (p Point) GetX() float64 { return p.X }
func (p Point) GetY() float64 { return p.Y }

func (p *Point) SetX(x float64) { p.X = x }
func (p *Point) SetY(y float64) { p.Y = y }

// Vector returns a Vector with the same X and Y values as Point.
func (p Point) Vector() Vector { return Vector{X: p.X, Y: p.Y} }

// HitTest returns a boolean indicating whether the X and Y values of Point are
// within the target HitTester.
func (p Point) HitTest(target HitTester) bool {
	return target.HitTest(p.X, p.Y)
}

// InRect returns a bool indicating whether the X and Y values of Point are
// within the sdl.Rect.
func (p Point) InRect(r sdl.Rect) bool {
	return hitTestRect(p.X, p.Y,
		float64(r.X), float64(r.Y),
		float64(r.W), float64(r.H),
	)
}

// InFRect returns a bool indicating whether the X and Y values of Point are
// within the sdl.FRect.
func (p Point) InFRect(r sdl.FRect) bool {
	return hitTestRect(p.X, p.Y,
		float64(r.X), float64(r.Y),
		float64(r.W), float64(r.H),
	)
}
