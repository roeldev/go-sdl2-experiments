// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package physics

import (
	"math"

	"github.com/go-pogo/sdlkit/geom"
)

type CircleCollider struct {
	shape *geom.Circle
}

func (cc *CircleCollider) Shape() geom.Shape { return cc.shape }

func (cc *CircleCollider) Bounds() AABB {
	return BoundsFromCenterSize(
		cc.shape.X, cc.shape.Y,
		cc.shape.Radius*2,
		cc.shape.Radius*2,
	)
}

// HitTestXY forwards the hit-test call to the geom.Circle HitTest method and
// returns true when the point [x y] is inside the collider.
func (cc *CircleCollider) HitTestXY(xy geom.XYGetter) bool {
	return cc.HitTest(xy.GetX(), xy.GetY())
}

func ResolveCollidingCircles(a, b *geom.Circle) bool {
	x := a.X - b.X
	y := a.Y - b.Y
	r := a.Radius + b.Radius

	distance := (x * x) + (y * y)
	if distance > (r * r) {
		return false
	}

	distance = math.Sqrt(distance)
	overlap := (distance - r) / 2

	a.X -= overlap * x / distance
	a.Y -= overlap * y / distance
	b.X += overlap * x / distance
	b.Y += overlap * y / distance
	return true
}
