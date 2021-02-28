// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package math provides additional game math related functions.
package math

import (
	stdmath "math"
)

const (
	PiHalf   = stdmath.Pi / 2
	PiDouble = stdmath.Pi * 2
	R2D      = 180 / stdmath.Pi // multiply with radians to get degrees
	D2R      = stdmath.Pi / 180 // multiply with degrees to get radians
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
