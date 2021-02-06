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
		"Center":       {APCenter, geom.Point{X: 25, Y: 40}},
		"Top":          {APTop, geom.Point{X: 0, Y: 20}},
		"TopCenter":    {APTopCenter, geom.Point{X: 25, Y: 20}},
		"Bottom":       {APBottom, geom.Point{X: 0, Y: 60}},
		"BottomCenter": {APBottomCenter, geom.Point{X: 25, Y: 60}},
		"Left":         {APLeft, geom.Point{X: 10, Y: 0}},
		"LeftMiddle":   {APLeftMiddle, geom.Point{X: 10, Y: 40}},
		"Right":        {APRight, geom.Point{X: 40, Y: 0}},
		"RightMiddle":  {APRightMiddle, geom.Point{X: 40, Y: 40}},
		"TopLeft":      {APTopLeft, geom.Point{X: 10, Y: 20}},
		"TopRight":     {APTopRight, geom.Point{X: 40, Y: 20}},
		"BottomLeft":   {APBottomLeft, geom.Point{X: 10, Y: 60}},
		"BottomRight":  {APBottomRight, geom.Point{X: 40, Y: 60}},
	}

	area := sdl.Rect{X: 10, Y: 20, W: 30, H: 40}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var have geom.Point
			AlignRect(tc.Pos, &have, area)
			assert.Equal(t, tc.Want, have)
		})
	}
}
