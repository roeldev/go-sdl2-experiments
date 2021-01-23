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
	// Origin point is relative to position.
	Origin Point

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
		rad := DegToRad((180 - angle) * float64(i))
		points[i].X = math.Cos(rad) * cr
		points[i].Y = math.Sin(rad) * cr
	}

	return NewPolygon(x, y, points)
}

// NewTrigon creates a new Polygon with 3 sides.
func NewTrigon(x, y, w, h float64) *Polygon {
	t := NewPolygon(x, y, []Point{{X: w}, {Y: h / 2}, {Y: -h / 2}})

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

	// t.area = (w * h) / 2
	return t
}

// NewQuad creates a new Polygon where all angles are 90 degrees. It differs
// from a Rectangle in the fact it can be rotated, scaled and skewed using a
// Matrix.
func NewQuad(x, y, w, h float64) *Polygon {
	q := NewPolygon(x, y, []Point{{0, 0}, {w, 0}, {w, h}, {0, h}})
	q.Origin.X = -w / 2
	q.Origin.Y = -h / 2
	// q.area = w * h
	return q
}

func (p *Polygon) Position() Point { return Point{X: p.X, Y: p.Y} }

func (p *Polygon) Model() []Point { return p.model }

func (p *Polygon) Edges() []Point {
	res := make([]Point, p.len)
	for i, pt := range p.actual {
		res[i].X = pt.X + p.X
		res[i].Y = pt.Y + p.Y
	}
	return res
}

func (p *Polygon) LinePoints() []sdl.Point { return p.lines }

func (p *Polygon) Transform(matrix Matrix) {
	for i, pt := range p.model {
		ax := ((pt.X + p.Origin.X) * matrix[ME_A]) + ((pt.Y + p.Origin.Y) * matrix[ME_C]) + matrix[ME_TX]
		ay := ((pt.X + p.Origin.X) * matrix[ME_B]) + ((pt.Y + p.Origin.Y) * matrix[ME_D]) + matrix[ME_TY]

		p.actual[i].X = ax
		p.actual[i].Y = ay
		p.lines[i].X = int32(p.X + ax)
		p.lines[i].Y = int32(p.Y + ay)
	}

	// close draw loop with first point
	p.lines[p.len] = p.lines[0]
}
