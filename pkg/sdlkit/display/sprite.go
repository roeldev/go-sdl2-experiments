// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package display

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/geom"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/math"
)

type Sprite struct {
	ColorTransform
	TextureTransform

	X, Y,
	W, H float64

	origin geom.Point // origin point relative to X and Y.
	clip   sdlkit.TextureClip
}

func NewSprite(clip sdlkit.TextureClip) *Sprite {
	// sprite heeft translate nodig
	s := &Sprite{
		W:    float64(clip.Location.W),
		H:    float64(clip.Location.H),
		clip: clip,
	}
	s.Reset()
	return s
}

func MustNewSprite(clip sdlkit.TextureClip, possibleErr error) *Sprite {
	if possibleErr != nil {
		sdlkit.FailOnErr(possibleErr)
	}

	return NewSprite(clip)
}

func (s *Sprite) GetX() float64  { return s.X }
func (s *Sprite) GetY() float64  { return s.Y }
func (s *Sprite) SetX(x float64) { s.X = x }
func (s *Sprite) SetY(y float64) { s.Y = y }

func (s *Sprite) Clip() sdlkit.TextureClip { return s.clip }
func (s *Sprite) Origin() *geom.Point      { return &s.origin }

func (s *Sprite) AbsoluteOrigin() geom.Point {
	return geom.Point{X: s.X + s.origin.X, Y: s.Y + s.origin.Y}
}

func (s *Sprite) Draw(canvas *sdlkit.Canvas) {
	canvas.DrawTextureClipEx(s.clip,
		sdl.Rect{
			X: int32(s.X - (s.W / 2)),
			Y: int32(s.Y - (s.H / 2)),
			W: int32(s.W),
			H: int32(s.H),
		},
		// from radians to degrees
		s.Rotation*math.R2D,
		// center point is relative to top left (0,0) of texture
		sdl.Point{X: int32(s.origin.X + s.W/2), Y: int32(s.origin.Y + s.H/2)},
		s.Flip,
	)
}

// canvas.DrawTextureClip(s.clip, sdl.Rect{
// 	X: int32(s.X - (s.W / 2)),
// 	Y: int32(s.Y - (s.H / 2)),
// 	W: int32(s.W),
// 	H: int32(s.H),
// })
