// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package geom

import (
	"github.com/veandco/go-sdl2/sdl"
)

type AlignPosition uint

const (
	APCenter AlignPosition = iota + 1
	APTop
	APTopCenter
	APBottom
	APBottomCenter
	APLeft
	APLeftMiddle
	APRight
	APRightMiddle
	APTopLeft
	APTopRight
	APBottomLeft
	APBottomRight
)

func Align(pos AlignPosition, pt *Point, x, y, w, h float64) {
	// x pos
	switch pos {
	case APCenter:
		fallthrough
	case APTopCenter:
		fallthrough
	case APBottomCenter:
		pt.X = x + (w / 2)

	case APLeft:
		fallthrough
	case APLeftMiddle:
		fallthrough
	case APTopLeft:
		fallthrough
	case APBottomLeft:
		pt.X = x

	case APRight:
		fallthrough
	case APRightMiddle:
		fallthrough
	case APTopRight:
		fallthrough
	case APBottomRight:
		pt.X = x + w
	}

	// y pos
	switch pos {
	case APCenter:
		fallthrough
	case APLeftMiddle:
		fallthrough
	case APRightMiddle:
		pt.Y = y + (h / 2)

	case APTop:
		fallthrough
	case APTopLeft:
		fallthrough
	case APTopCenter:
		fallthrough
	case APTopRight:
		pt.Y = y

	case APBottom:
		fallthrough
	case APBottomLeft:
		fallthrough
	case APBottomCenter:
		fallthrough
	case APBottomRight:
		pt.Y = y + h
	}
}

func AlignRect(pos AlignPosition, pt *Point, r sdl.Rect) {
	Align(pos, pt, float64(r.X), float64(r.Y), float64(r.W), float64(r.H))
}

func AlignFRect(pos AlignPosition, pt *Point, r sdl.FRect) {
	Align(pos, pt, float64(r.X), float64(r.Y), float64(r.W), float64(r.H))
}
