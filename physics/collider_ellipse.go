// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package physics

import (
	"github.com/go-pogo/sdlkit/geom"
)

type EllipseCollider struct {
	shape *geom.Ellipse
}

func (ec *EllipseCollider) Shape() geom.Shape { return ec.shape }

func (ec *EllipseCollider) Bounds() AABB {
	return BoundsFromCenterSize(
		ec.shape.X, ec.shape.Y,
		ec.shape.RadiusX*2,
		ec.shape.RadiusY*2,
	)
}

// HitTestXY returns true when the position of the XYGetter is inside Ellipse.
func (ec *EllipseCollider) HitTestXY(xy geom.XYGetter) bool {
	return ec.HitTest(xy.GetX(), xy.GetY())
}

func ResolveCollidingEllipses(a, b *geom.Ellipse) bool {
	return false
}
