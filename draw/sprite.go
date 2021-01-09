// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package draw

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/go-pogo/sdlkit/geom"
)

type Sprite struct {
	geom.Point // center of sprite
	geom.Transform
	Flip   sdl.RendererFlip
	Bounds geom.Bounds

	tx  *sdl.Texture
	src sdl.Rect // location within texture
}

func NewSprite(tx *sdl.Texture, w, h int32) (*Sprite, error) {
	if w == 0 || h == 0 {
		var err error
		_, _, w, h, err = tx.Query()
		if err != nil {
			return nil, err
		}
	}

	return &Sprite{
		Bounds: geom.RectBounds(w, h),
		tx:     tx,
		src:    sdl.Rect{W: w, H: h},
	}, nil
}

func (s *Sprite) Update(_ float32) {
	s.Bounds.Transform(s.Point, s.Transform)
}

func (s *Sprite) Draw(r *sdl.Renderer) error {
	return r.CopyEx(s.tx, &s.src, &s.Bounds.Rect, float64(s.Rotation), nil, s.Flip)
}
