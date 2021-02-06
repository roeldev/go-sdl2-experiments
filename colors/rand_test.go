// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build test_colors

package colors

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/veandco/go-sdl2/sdl"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func TestRand(t *testing.T) {
	tests := map[string]struct {
		fn func(r *rand.Rand) sdl.Color
		in []sdl.Color
	}{
		"all":    {fn: RandColor, in: AllColors},
		"blue":   {fn: RandBlue, in: BlueColors},
		"brown":  {fn: RandBrown, in: BrownColors},
		"green":  {fn: RandGreen, in: GreenColors},
		"grey":   {fn: RandGrey, in: GreyColors},
		"orange": {fn: RandOrange, in: OrangeColors},
		"pink":   {fn: RandPink, in: PinkColors},
		"purple": {fn: RandPurple, in: PurpleColors},
		"red":    {fn: RandRed, in: RedColors},
		"white":  {fn: RandWhite, in: WhiteColors},
		"yellow": {fn: RandYellow, in: YellowColors},
	}

	a := 0
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			n := len(tc.in)
			a += n
			c := make(map[sdl.Color]int, n)
			for i := 0; i < n*100; i++ {
				t.Run("contains", func(t *testing.T) {
					v := tc.fn(rng)
					assert.Contains(t, tc.in, v)
					c[v]++
				})
			}

			assert.Len(t, c, n)
		})
	}
}

func TestDoubles(t *testing.T) {
	tests := map[string][]sdl.Color{
		"all":    AllColors,
		"blue":   BlueColors,
		"brown":  BrownColors,
		"green":  GreenColors,
		"grey":   GreyColors,
		"orange": OrangeColors,
		"pink":   PinkColors,
		"purple": PurpleColors,
		"red":    RedColors,
		"white":  WhiteColors,
		"yellow": YellowColors,
	}

	for name, colors := range tests {
		t.Run(name, func(t *testing.T) {
			n := len(colors)
			u := make(map[sdl.Color]int, n)

			for _, c := range colors {
				u[c]++
				if u[c] > 1 {
					fmt.Printf("%s already contains %#v\n", name, c)
				}
			}

			assert.Len(t, u, n)
		})
	}
}
