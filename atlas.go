// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdlkit

import (
	"github.com/go-pogo/errors"
	"github.com/veandco/go-sdl2/sdl"
)

type TextureAtlas struct {
	tx      *sdl.Texture
	sources []sdl.Rect
	names   map[string]int
	uniform bool
}

func NewTextureAtlas(tx *sdl.Texture, srcs map[string]sdl.Rect) *TextureAtlas {
	ta := &TextureAtlas{
		tx:      tx,
		sources: make([]sdl.Rect, 0, len(srcs)),
		names:   make(map[string]int, len(srcs)),
	}

	for n, s := range srcs {
		if i, exists := ta.names[n]; exists {
			ta.sources[i] = s
			continue
		}

		ta.names[n] = len(ta.sources)
		ta.sources = append(ta.sources, s)
	}

	return ta
}

func NewUniformTextureAtlas(tx *sdl.Texture, cellW, cellH int32, total uint8) (*TextureAtlas, error) {
	if total < 1 {
		return nil, errors.Newf("sdlkit: a TextureAtlas needs at least 1 cell")
	}

	_, _, txW, txH, err := tx.Query()
	if err != nil {
		return nil, err
	}

	var x, y int32
	ta := &TextureAtlas{
		tx:      tx,
		sources: make([]sdl.Rect, 0, total),
		names:   make(map[string]int),
		uniform: true,
	}

Loop:
	for y = 0; y < txH; y += cellH {
		for x = 0; x < txW; x += cellW {
			ta.sources = append(ta.sources, sdl.Rect{X: x, Y: y, W: cellW, H: cellH})

			total--
			if total <= 0 {
				break Loop
			}
		}
	}

	return ta, nil
}

func (ta *TextureAtlas) Texture() *sdl.Texture { return ta.tx }

func (ta *TextureAtlas) Index(i int) (sdl.Rect, bool) {
	if len(ta.sources) <= i {
		return sdl.Rect{}, false
	}

	return ta.sources[i], true
}

func (ta *TextureAtlas) Name(name string) (sdl.Rect, bool) {
	if i, ok := ta.names[name]; ok {
		return ta.sources[i], true
	}

	return sdl.Rect{}, false
}

func (ta *TextureAtlas) Names() []string {
	res := make([]string, 0, len(ta.names))
	for n := range ta.names {
		res = append(res, n)
	}
	return res
}
