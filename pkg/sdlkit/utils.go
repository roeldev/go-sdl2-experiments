// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdlkit

import (
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

var rng *rand.Rand

// RNG returns a new rand.Rand with the current unix time as source.
func RNG() *rand.Rand {
	if rng == nil {
		rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	return rng
}

func ShrinkRect(rect sdl.Rect, amount int32) sdl.Rect {
	return sdl.Rect{
		X: rect.X + amount,
		Y: rect.Y + amount,
		W: rect.W - amount - amount,
		H: rect.H - amount - amount,
	}
}
