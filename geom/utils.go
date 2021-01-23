// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package geom

import (
	"math"
)

func Norm(v, a, b float64) float64 {
	if a < b {
		return (v - a) / (b - a)
	} else {
		return (v - b) / (a - b)
	}
}

func Lerp(cur, dest, t float64) float64 {
	return cur + (dest-cur)*t
}

func Lerpx(cur, dest, t float64) float64 {
	diff := (dest - cur) * t
	cur += diff
	if diff > 0 && cur+0.05 > dest {
		return dest
	}
	if diff < 0 && cur-0.05 < dest {
		return dest
	}
	return cur + diff
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

// https://www.khanacademy.org/computing/computer-programming/programming-natural-simulations/programming-vectors/a/more-vector-math
func Distance(x0, y0, x1, y1 float64) float64 {
	return math.Sqrt(DistanceSq(x0, y0, x1, y1))
}

func DistanceSq(x0, y0, x1, y1 float64) float64 {
	dx, dy := x1-x0, y1-y0
	return (dx * dx) + (dy * dy)
}

func Centroid(points ...[2]float64) (x, y float64) {
	for _, p := range points {
		x += p[0]
		y += p[1]
	}
	n := float64(len(points))
	return x / n, y / n
}
