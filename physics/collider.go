// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package physics

import (
	"github.com/go-pogo/sdlkit/geom"
)

type Collider interface {
	HitTester
	Shape() geom.Shape
	Bounds() AABB
}

func NewCollider(shape geom.Shape) Collider {
	switch s := shape.(type) {
	case *geom.Ellipse:
		return &EllipseCollider{
			shape: s,
		}

	case *geom.Circle:
		return &CircleCollider{
			shape: s,
		}

	case *geom.Rect:
		return &RectCollider{
			shape: s,
		}

	case geom.PolygonShape:
		return &PolygonCollider{
			poly: s,
		}
	}

	return nil
}
