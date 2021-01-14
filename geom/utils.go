// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package geom

import (
	"math"
)

const (
	r2dPi = 180 / math.Pi
	d2rPi = math.Pi / 180
)

func RadToDeg(rad float64) float64 { return rad * r2dPi }

func DegToRad(deg float64) float64 { return deg * d2rPi }

func Norm(val, min, max float64) float64 {
	return (val - min) / (max - min)
}

func Lerp(min, max, t float64) float64 {
	return (max-min)*t + min
}

func Clamp(val, min, max float64) float64 {
	if val < min {
		val = min
	}
	if val > max {
		val = max
	}
	return val
}

func Dist(x0, y0, x1, y1 float64) float64 {
	return math.Sqrt(DistSq(x0, y0, x1, y1))
}

func DistSq(x0, y0, x1, y1 float64) float64 {
	dx, dy := x1-x0, y1-y0
	return (dx * dx) + (dy * dy)
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
