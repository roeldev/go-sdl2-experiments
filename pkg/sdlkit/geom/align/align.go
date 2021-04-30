// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package align

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/geom"
)

type Alignment uint

const (
	ToCenter Alignment = iota + 1
	ToTop
	ToTopCenter
	ToBottom
	ToBottomCenter
	ToLeft
	ToLeftMiddle
	ToRight
	ToRightMiddle
	ToTopLeft
	ToTopRight
	ToBottomLeft
	ToBottomRight
)

func Values(to Alignment, x, y *float64, rX, rY, rW, rH float64) {
	// x pos
	switch to {
	case ToCenter:
		fallthrough
	case ToTopCenter:
		fallthrough
	case ToBottomCenter:
		*x = rX + (rW / 2)

	case ToLeft:
		fallthrough
	case ToLeftMiddle:
		fallthrough
	case ToTopLeft:
		fallthrough
	case ToBottomLeft:
		*x = rX

	case ToRight:
		fallthrough
	case ToRightMiddle:
		fallthrough
	case ToTopRight:
		fallthrough
	case ToBottomRight:
		*x = rX + rW
	}

	// y pos
	switch to {
	case ToCenter:
		fallthrough
	case ToLeftMiddle:
		fallthrough
	case ToRightMiddle:
		*y = rY + (rH / 2)

	case ToTop:
		fallthrough
	case ToTopLeft:
		fallthrough
	case ToTopCenter:
		fallthrough
	case ToTopRight:
		*y = rY

	case ToBottom:
		fallthrough
	case ToBottomLeft:
		fallthrough
	case ToBottomCenter:
		fallthrough
	case ToBottomRight:
		*y = rY + rH
	}
}

func Point(to Alignment, pt *geom.Point, rX, rY, rW, rH float64) {
	Values(to, &pt.X, &pt.Y, rX, rY, rW, rH)
}

func PointInSdlRect(to Alignment, pt *geom.Point, r sdl.Rect) {
	Values(to, &pt.X, &pt.Y, float64(r.X), float64(r.Y), float64(r.W), float64(r.H))
}

func XYInSdlRect(to Alignment, xy geom.XYSetter, r sdl.Rect) {
	var x, y float64
	Values(to, &x, &y, float64(r.X), float64(r.Y), float64(r.W), float64(r.H))
	xy.SetX(x)
	xy.SetY(y)
}
