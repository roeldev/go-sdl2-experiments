// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package physics

import (
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/geom"
)

type HitTester interface {
	// HitTest returns true when the x and y values are within the HitTester.
	HitTest(x, y float64) bool
	HitTestXY(xy geom.XYGetter) bool
}

func HitTestCircle(x, y float64, circle geom.Circle) bool {
	dx, dy := circle.X-x, circle.Y-y
	return ((dx * dx) + (dy * dy)) < (circle.Radius * circle.Radius)
}

// HitTest returns true when position [x y] is inside the Circle. It calculates
// the squared distance between [x y] and the Circle's center X Y and compares
// this with the squared Radius.
func (cc *CircleCollider) HitTest(x, y float64) bool {
	dx, dy := cc.shape.X-x, cc.shape.Y-y
	return ((dx * dx) + (dy * dy)) < (cc.shape.Radius * cc.shape.Radius)
}

func HitTestEllipse(x, y float64, ellipse geom.Ellipse) bool {
	return ((x-ellipse.X)*(x-ellipse.X)/(ellipse.RadiusX*ellipse.RadiusX))+
		((y-ellipse.Y)*(y-ellipse.Y)/(ellipse.RadiusY*ellipse.RadiusY)) <= 1
}

// HitTest returns true when position [x y] is inside the EllipseCollider's
// geom.Ellipse shape.
// https://math.stackexchange.com/questions/76457/check-if-a-point-is-within-an-ellipse
func (ec *EllipseCollider) HitTest(x, y float64) bool {
	return ((x-ec.shape.X)*(x-ec.shape.X)/(ec.shape.RadiusX*ec.shape.RadiusX))+
		((y-ec.shape.Y)*(y-ec.shape.Y)/(ec.shape.RadiusY*ec.shape.RadiusY)) <= 1
}

func HitTestRect(x, y float64, rect geom.Rect) bool {
	w, h := rect.W/2, rect.H/2
	return (x >= rect.X-w) && (x < rect.X+w) && (y >= rect.Y-h) && (y < rect.Y+h)
}

// HitTest returns true when the x and y values are within the Rect.
func (rc *RectCollider) HitTest(x, y float64) bool {
	w, h := rc.shape.W/2, rc.shape.H/2
	return (x >= rc.shape.X-w) && (x < rc.shape.X+w) && (y >= rc.shape.Y-h) && (y < rc.shape.Y+h)
}

// // https://wrf.ecse.rpi.edu/Research/Short_Notes/pnpoly.html
func HitTestPolygon(x, y float64, vertices []geom.Point) (hit bool) {
	n := len(vertices)
	for i, j := 0, n-1; i < n; i++ {
		u, v := vertices[i], vertices[j]
		if ((u.Y > y) != (v.Y > y)) &&
			(x < ((v.X-u.X)*(y-u.Y))/(v.Y-u.Y)+u.X) {
			hit = !hit
		}
		j = i
	}
	return hit
}
