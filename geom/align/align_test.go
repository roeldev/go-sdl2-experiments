// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package align

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/go-pogo/sdlkit/geom"
)

func TestPos(t *testing.T) {
	tests := map[string]struct {
		Pos  Alignment
		Want geom.Point
	}{
		"Center":       {ToCenter, geom.Point{X: 25, Y: 40}},
		"Top":          {ToTop, geom.Point{X: 0, Y: 20}},
		"TopCenter":    {ToTopCenter, geom.Point{X: 25, Y: 20}},
		"Bottom":       {ToBottom, geom.Point{X: 0, Y: 60}},
		"BottomCenter": {ToBottomCenter, geom.Point{X: 25, Y: 60}},
		"Left":         {ToLeft, geom.Point{X: 10, Y: 0}},
		"LeftMiddle":   {ToLeftMiddle, geom.Point{X: 10, Y: 40}},
		"Right":        {ToRight, geom.Point{X: 40, Y: 0}},
		"RightMiddle":  {ToRightMiddle, geom.Point{X: 40, Y: 40}},
		"TopLeft":      {ToTopLeft, geom.Point{X: 10, Y: 20}},
		"TopRight":     {ToTopRight, geom.Point{X: 40, Y: 20}},
		"BottomLeft":   {ToBottomLeft, geom.Point{X: 10, Y: 60}},
		"BottomRight":  {ToBottomRight, geom.Point{X: 40, Y: 60}},
	}

	area := sdl.Rect{X: 10, Y: 20, W: 30, H: 40}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var have geom.Point
			PointInSdlRect(tc.Pos, &have, area)
			assert.Equal(t, tc.Want, have)
		})
	}
}
