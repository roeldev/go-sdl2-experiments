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
		Pos  AlignPosition
		Want Point
	}{
		"Center":       {APCenter, Point{X: 25, Y: 40}},
		"Top":          {APTop, Point{X: 0, Y: 20}},
		"TopCenter":    {APTopCenter, Point{X: 25, Y: 20}},
		"Bottom":       {APBottom, Point{X: 0, Y: 60}},
		"BottomCenter": {APBottomCenter, Point{X: 25, Y: 60}},
		"Left":         {APLeft, Point{X: 10, Y: 0}},
		"LeftMiddle":   {APLeftMiddle, Point{X: 10, Y: 40}},
		"Right":        {APRight, Point{X: 40, Y: 0}},
		"RightMiddle":  {APRightMiddle, Point{X: 40, Y: 40}},
		"TopLeft":      {APTopLeft, Point{X: 10, Y: 20}},
		"TopRight":     {APTopRight, Point{X: 40, Y: 20}},
		"BottomLeft":   {APBottomLeft, Point{X: 10, Y: 60}},
		"BottomRight":  {APBottomRight, Point{X: 40, Y: 60}},
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
