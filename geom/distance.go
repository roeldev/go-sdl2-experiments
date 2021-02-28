// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package geom

import (
	"math"
)

// Distance calculates the distance between the points [x1 y1] and [x2 y2].
func Distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(DistanceSq(x1, y1, x2, y2))
}

// Dist calculates the distance between the Vector and the target Vector.
func (v Vector) Dist(target Vector) float64 {
	return math.Sqrt(v.DistSq(target))
}

// DistanceSq calculates the squared distance between the points [x1 y1] and
// [x2 y2].
func DistanceSq(x1, y1, x2, y2 float64) float64 {
	dx, dy := x2-x1, y2-y1
	return (dx * dx) + (dy * dy)
}

// DistSq calculates the squared distance between the Vector and the target
// Vector.
func (v Vector) DistSq(target Vector) float64 {
	dx, dy := target.X-v.X, target.Y-v.Y
	return (dx * dx) + (dy * dy)
}
