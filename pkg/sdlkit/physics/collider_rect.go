// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package physics

import (
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/geom"
)

type RectCollider struct {
	shape *geom.Rect
}

func (rc *RectCollider) Shape() geom.Shape { return rc.shape }

func (rc *RectCollider) Bounds() AABB {
	return BoundsFromCenterSize(rc.shape.X, rc.shape.Y, rc.shape.W, rc.shape.H)
}

// HitTestXY forwards the hit-test call to the geom.Rect HitTest method and
// returns true when the point [x y] is inside the collider.
func (rc *RectCollider) HitTestXY(xy geom.XYGetter) bool {
	return rc.HitTest(xy.GetX(), xy.GetY())
}

func ResolveCollidingRects(a, b *geom.Rect) bool {
	if b.X >= a.X+a.W || a.X >= b.X+b.W || b.Y >= a.Y+a.H || a.Y >= b.Y+b.H {
		return false
	}

	return true
}
