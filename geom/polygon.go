// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package geom

import (
	"math"
)

type PolygonShape interface {
	Shape
	Model() []Point
	Vertices() []Point
}

// https://en.wikipedia.org/wiki/Simple_polygon
// https://en.wikipedia.org/wiki/Polygon
type Polygon struct {
	// X and Y indicate the absolute position of the Polygon.
	X, Y float64

	origin Point   // origin point relative to position.
	len    int     // number of edges
	model  []Point // original points relative to pos
	actual []Point // actual points relative to pos
}

func NewPolygon(x, y float64, pts []Point) *Polygon {
	n := len(pts)
	p := &Polygon{
		X:      x,
		Y:      y,
		len:    n,
		model:  pts,
		actual: make([]Point, n),
	}
	p.Transform(IdentityMatrix())
	return p
}

// NewRegularPolygon creates a new regular polygon where all angles are equal
// in measure and all sides have the same length. Value cr is the circumradius
// between the center of the Polygon and any vertex (edge point). Value n
// indicates the amount of vertices (edges) the polygon has, 3 for a triangle,
// 4 for a quad etc. This value cannot be lower than 3.
// See https://en.wikipedia.org/wiki/Regular_polygon for additional information
// about regular polygons.
func NewRegularPolygon(x, y float64, cr float64, n uint8) *Polygon {
	if n < 3 {
		panic("geom: cannot create regular Polygon with less than 3 sides")
	}

	angle := float64(n-2) * 180 / float64(n)
	points := make([]Point, n)

	var i uint8
	for ; i < n; i++ {
		rad := DegToRad(360 - (180-angle)*float64(i))
		points[i].X = math.Cos(rad) * cr
		points[i].Y = math.Sin(rad) * cr
	}

	return NewPolygon(x, y, points)
}

// NewTrigon creates a new Polygon with 3 sides.
func NewTrigon(x, y, w, h float64) *Polygon {
	t := NewPolygon(x, y, []Point{
		{X: w},      // right middle
		{Y: -h / 2}, // top left
		{Y: h / 2},  // bottom left
	})

	// calculate centroid of triangle points
	var dx, dy float64
	for _, p := range t.model {
		dx += p.X
		dy += p.Y
	}

	dx /= float64(t.len)
	dy /= float64(t.len)

	// update points according to centroid
	for i := range t.model {
		t.model[i].X -= dx
		t.model[i].Y -= dy
	}

	t.Transform(IdentityMatrix())
	return t
}

// NewQuad creates a new polygon where all angles are 90 degrees. It differs
// from a Rect in the fact it can be rotated, scaled and skewed using a
// Matrix.
func NewQuad(x, y, w, h float64) *Polygon {
	w, h = w/2, h/2
	return NewPolygon(x, y, []Point{
		{X: w, Y: -h},  // top right
		{X: -w, Y: -h}, // top left
		{X: -w, Y: h},  // bottom left
		{X: w, Y: h},   // bottom right
	})
}

func (p *Polygon) GetX() float64  { return p.X }
func (p *Polygon) GetY() float64  { return p.Y }
func (p *Polygon) SetX(x float64) { p.X = x }
func (p *Polygon) SetY(y float64) { p.Y = y }

func (p *Polygon) Origin() *Point { return &p.origin }

func (p *Polygon) AbsoluteOrigin() Point {
	return Point{X: p.GetX() + p.origin.X, Y: p.GetY() + p.origin.Y}
}

func (p *Polygon) Model() []Point { return p.model }

// https://www.wikihow.com/Calculate-the-Area-of-a-Polygon
func (p *Polygon) Area() float64 {
	var res float64
	var p1, p2 Point

	n := p.len - 1
	for i := 0; i <= n; i++ {
		p1 = p.actual[i]
		if i == n {
			p2 = p.actual[0]
		} else {
			p2 = p.actual[i+1]
		}

		res += (p1.X * p2.Y) - (p1.Y * p2.X)
	}

	return math.Abs(res) / 2
}

// Vertices returns the edges (vertices or vertexes) of the Polygon.
func (p *Polygon) Vertices() []Point {
	res := make([]Point, p.len)
	for i, pt := range p.actual {
		res[i].X = pt.X + p.X
		res[i].Y = pt.Y + p.Y
	}
	return res
}

func (p *Polygon) Transform(matrix Matrix) {
	for i, pt := range p.model {
		ax := ((pt.X + p.origin.X) * matrix[ME_A]) + ((pt.Y + p.origin.Y) * matrix[ME_C]) + matrix[ME_TX]
		ay := ((pt.X + p.origin.X) * matrix[ME_B]) + ((pt.Y + p.origin.Y) * matrix[ME_D]) + matrix[ME_TY]

		p.actual[i].X = ax
		p.actual[i].Y = ay
	}
}
