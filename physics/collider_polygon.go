// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package physics

import (
	"math"

	"github.com/go-pogo/sdlkit/geom"
)

type PolygonCollider struct {
	poly geom.PolygonShape
}

func (pc *PolygonCollider) Shape() geom.Shape          { return pc.poly }
func (pc *PolygonCollider) Polygon() geom.PolygonShape { return pc.poly }

func (pc *PolygonCollider) Bounds() AABB {
	tl := geom.Point{X: math.MaxFloat64, Y: math.MaxFloat64}
	br := geom.Point{X: math.SmallestNonzeroFloat64, Y: math.SmallestNonzeroFloat64}
	for _, pt := range pc.poly.Vertices() {
		tl.X = math.Min(tl.X, pt.X)
		tl.Y = math.Min(tl.Y, pt.Y)
		br.X = math.Max(br.X, pt.X)
		br.Y = math.Max(br.Y, pt.Y)
	}

	return AABB{tl, br}
}

// HitTest tests if the provided point [x y] is inside the PolygonCollider's
// geom.Polygon shape.
func (pc *PolygonCollider) HitTest(x, y float64) (hit bool) {
	return HitTestPolygon(x, y, pc.poly.Vertices())
}

// HitTestXY forwards the hit-test call to the geom.polygon HitTest method and
// returns true when the point [x y] is inside the collider.
func (pc *PolygonCollider) HitTestXY(xy geom.XYGetter) bool {
	return HitTestPolygon(xy.GetX(), xy.GetY(), pc.poly.Vertices())
}

func ResolveCollidingPolygons(a, b *geom.Polygon) bool {
	return false
}
