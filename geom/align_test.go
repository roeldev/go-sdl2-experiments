// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package geom

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/veandco/go-sdl2/sdl"
)

func TestPos(t *testing.T) {
	tests := map[string]struct {
		Pos  Position
		Want Point
	}{
		"Center":       {PosCenter, Point{X: 25, Y: 40}},
		"Top":          {PosTop, Point{X: 0, Y: 20}},
		"TopCenter":    {PosTopCenter, Point{X: 25, Y: 20}},
		"Bottom":       {PosBottom, Point{X: 0, Y: 60}},
		"BottomCenter": {PosBottomCenter, Point{X: 25, Y: 60}},
		"Left":         {PosLeft, Point{X: 10, Y: 0}},
		"LeftMiddle":   {PosLeftMiddle, Point{X: 10, Y: 40}},
		"Right":        {PosRight, Point{X: 40, Y: 0}},
		"RightMiddle":  {PosRightMiddle, Point{X: 40, Y: 40}},
		"TopLeft":      {PosTopLeft, Point{X: 10, Y: 20}},
		"TopRight":     {PosTopRight, Point{X: 40, Y: 20}},
		"BottomLeft":   {PosBottomLeft, Point{X: 10, Y: 60}},
		"BottomRight":  {PosBottomRight, Point{X: 40, Y: 60}},
	}

	area := sdl.FRect{X: 10, Y: 20, W: 30, H: 40}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var have Point
			AlignFRect(tc.Pos, &have, area)
			assert.Equal(t, tc.Want, have)
		})
	}
}
