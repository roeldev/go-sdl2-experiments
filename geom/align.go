// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package geom

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Position uint

const (
	PosCenter Position = iota + 1
	PosTop
	PosTopCenter
	PosBottom
	PosBottomCenter
	PosLeft
	PosLeftMiddle
	PosRight
	PosRightMiddle
	PosTopLeft
	PosTopRight
	PosBottomLeft
	PosBottomRight
)

func Align(pos Position, pt *Point, x, y, w, h float32) {
	// x pos
	switch pos {
	case PosCenter:
		fallthrough
	case PosTopCenter:
		fallthrough
	case PosBottomCenter:
		pt.X = x + (w / 2)

	case PosLeft:
		fallthrough
	case PosLeftMiddle:
		fallthrough
	case PosTopLeft:
		fallthrough
	case PosBottomLeft:
		pt.X = x

	case PosRight:
		fallthrough
	case PosRightMiddle:
		fallthrough
	case PosTopRight:
		fallthrough
	case PosBottomRight:
		pt.X = x + w
	}

	// y pos
	switch pos {
	case PosCenter:
		fallthrough
	case PosLeftMiddle:
		fallthrough
	case PosRightMiddle:
		pt.Y = y + (h / 2)

	case PosTop:
		fallthrough
	case PosTopLeft:
		fallthrough
	case PosTopCenter:
		fallthrough
	case PosTopRight:
		pt.Y = y

	case PosBottom:
		fallthrough
	case PosBottomLeft:
		fallthrough
	case PosBottomCenter:
		fallthrough
	case PosBottomRight:
		pt.Y = y + h
	}
}

func AlignRect(pos Position, pt *Point, r sdl.Rect) {
	Align(pos, pt, float32(r.X), float32(r.Y), float32(r.W), float32(r.H))
}

func AlignFRect(pos Position, pt *Point, r sdl.FRect) {
	Align(pos, pt, r.X, r.Y, r.W, r.H)
}
