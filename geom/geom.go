// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package geom defines a two-dimensional coordinate system.
package geom

import (
	"math"
)

type XYGetter interface {
	GetX() float64
	GetY() float64
}

type XYSetter interface {
	SetX(x float64)
	SetY(y float64)
}

type XY interface {
	XYGetter
	XYSetter
}

const (
	R2D = 180 / math.Pi // multiply with radians to get degrees
	D2R = math.Pi / 180 // multiply with degrees to get radians
)

func RadToDeg(rad float64) float64 { return rad * R2D }

func DegToRad(deg float64) float64 { return deg * D2R }

func Norm(v, a, b float64) float64 {
	if a < b {
		return (v - a) / (b - a)
	} else {
		return (v - b) / (a - b)
	}
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

func Centroid(points ...[2]float64) (x, y float64) {
	for _, p := range points {
		x += p[0]
		y += p[1]
	}
	n := float64(len(points))
	return x / n, y / n
}
