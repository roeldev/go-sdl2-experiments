// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package geom

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

// https://en.wikipedia.org/wiki/Simple_polygon
// https://en.wikipedia.org/wiki/Polygon
type Polygon struct {
	// X and Y indicate the absolute position of the Polygon.
	X, Y float64

	origin Point       // origin point relative to position.
	len    int         // number of edges
	model  []Point     // original points relative to pos
	actual []Point     // actual points relative to pos
	lines  []sdl.Point // points of lines to draw
}

func NewPolygon(x, y float64, pts []Point) *Polygon {
	n := len(pts)
	p := &Polygon{
		X:      x,
		Y:      y,
		len:    n,
		model:  pts,
		actual: make([]Point, n),
		lines:  make([]sdl.Point, n+1),
	}
	p.Transform(IdentityMatrix())
	return p
}

// NewRegularPolygon creates a new regular Polygon where all angles are equal
// in measure and all sides have the same length. Value cr is the circumradius
// between the center of the Polygon and any edge point. Value n indicates the
// amount of edges the Polygon has, 3 for a triangle, 4 for a quad etc., and
// cannot be lower than 3.
// See https://en.wikipedia.org/wiki/Regular_polygon for additional information
// about regular polygons.
func NewRegularPolygon(x, y, cr float64, n uint8) *Polygon {
	if n < 3 {
		panic("geom: cannot create regular polygon with less than 3 sides")
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

// NewQuad creates a new Polygon where all angles are 90 degrees. It differs
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
func (p *Polygon) Origin() *Point { return &p.origin }
func (p *Polygon) Model() []Point { return p.model }

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

func (p *Polygon) Bounds() AABB {
	bounds := AABB{
		X: math.MaxFloat64,
		Y: math.MaxFloat64,
		W: math.SmallestNonzeroFloat64,
		H: math.SmallestNonzeroFloat64,
	}
	for _, pt := range p.actual {
		bounds.X = math.Min(bounds.X, pt.X)
		bounds.Y = math.Min(bounds.Y, pt.Y)
		bounds.W = math.Max(bounds.W, pt.X)
		bounds.H = math.Max(bounds.H, pt.Y)
	}

	// convert max x/y to width and height
	bounds.W -= bounds.X
	bounds.H -= bounds.Y

	// min x/y are still relative to x/y of the polygon, fix this
	bounds.X += p.X
	bounds.Y += p.Y
	return bounds
}

func (p *Polygon) Edges() []Point {
	res := make([]Point, p.len)
	for i, pt := range p.actual {
		res[i].X = pt.X + p.X
		res[i].Y = pt.Y + p.Y
	}
	return res
}

func (p *Polygon) LinePoints() []sdl.Point { return p.lines }

// https://wrf.ecse.rpi.edu/Research/Short_Notes/pnpoly.html
func (p *Polygon) HitTest(x, y float64) bool {
	if p.len > 5 && !p.Bounds().HitTest(x, y) {
		return false
	}

	var c bool
	for i, j := 0, p.len-1; i < p.len; i++ {
		p1, p2 := p.actual[i], p.actual[j]
		if ((p1.Y+p.Y > y) != (p2.Y+p.Y > y)) &&
			(x < ((p2.X-p1.X)*(y-p1.Y-p.Y))/(p2.Y-p1.Y)+p1.X+p.X) {
			c = !c
		}
		j = i
	}

	return c
}

func (p *Polygon) HitTestXY(xy XYGetter) bool { return p.HitTest(xy.GetX(), xy.GetY()) }

func (p *Polygon) Transform(matrix Matrix) {
	for i, pt := range p.model {
		ax := ((pt.X + p.origin.X) * matrix[ME_A]) + ((pt.Y + p.origin.Y) * matrix[ME_C]) + matrix[ME_TX]
		ay := ((pt.X + p.origin.X) * matrix[ME_B]) + ((pt.Y + p.origin.Y) * matrix[ME_D]) + matrix[ME_TY]

		p.actual[i].X = ax
		p.actual[i].Y = ay
		p.lines[i].X = int32(ax + p.X)
		p.lines[i].Y = int32(ay + p.Y)
	}

	// close draw loop with first point
	p.lines[p.len] = p.lines[0]
}
