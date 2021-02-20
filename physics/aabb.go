package physics

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/go-pogo/sdlkit/geom"
)

// AABB is an axis-aligned bounding box with X and Y indicating its top left
// corner.
type AABB struct {
	TopLeft, BottomRight geom.Point
}

func BoundsFromCenterSize(x, y, w, h float64) AABB {
	return AABB{
		TopLeft:     geom.Point{X: x - w/2, Y: y - h/2},
		BottomRight: geom.Point{X: x + w/2, Y: y + h/2},
	}
}

func (b AABB) GetX() float64   { return b.TopLeft.X }
func (b AABB) GetY() float64   { return b.TopLeft.Y }
func (b AABB) Width() float64  { return b.BottomRight.X - b.TopLeft.X }
func (b AABB) Height() float64 { return b.BottomRight.Y - b.TopLeft.Y }

func (b AABB) Center() geom.Point {
	return geom.Point{
		X: b.TopLeft.X + (b.Width() / 2),
		Y: b.TopLeft.Y + (b.Height() / 2),
	}
}

func (b AABB) Rect() *sdl.Rect {
	return &sdl.Rect{
		X: int32(b.TopLeft.X), Y: int32(b.TopLeft.Y),
		W: int32(b.Width()), H: int32(b.Height())}
}

// HitTest returns true when the x and y values are within the AABB.
func (b AABB) HitTest(x, y float64) bool {
	return (x >= b.TopLeft.X) && (x < b.BottomRight.X) && (y >= b.TopLeft.Y) && (y < b.BottomRight.Y)
}

// HitTestXY returns true when the position of the XYGetter is inside AABB.
func (b AABB) HitTestXY(xy geom.XYGetter) bool {
	return b.HitTest(xy.GetX(), xy.GetY())
}

func (b AABB) Intersects(target AABB) bool {
	panic("physics: Intersects is not implemented")
	return false
}

func (b AABB) Union(target AABB) AABB {
	return AABB{
		TopLeft: geom.Point{
			X: math.Min(b.TopLeft.X, target.TopLeft.X),
			Y: math.Min(b.TopLeft.Y, target.TopLeft.Y),
		},
		BottomRight: geom.Point{
			X: math.Max(b.BottomRight.X, target.BottomRight.X),
			Y: math.Max(b.BottomRight.Y, target.BottomRight.Y),
		},
	}
}
