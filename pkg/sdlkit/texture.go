// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdlkit

import (
	stderrors "errors"

	"github.com/go-pogo/errors"
	"github.com/veandco/go-sdl2/sdl"
)

//goland:noinspection GoErrorStringFormat
var ErrInvalidTexture = stderrors.New("Invalid texture")

type TextureClip struct {
	Texture  *sdl.Texture
	Location sdl.Rect
}

func (tc TextureClip) Size() (float64, float64) {
	return float64(tc.Location.W), float64(tc.Location.H)
}

type TextureAtlas struct {
	texture   *sdl.Texture
	locations []sdl.Rect
	names     map[string]int
	uniform   bool
}

func NewTextureAtlas(tx *sdl.Texture, locations map[string]sdl.Rect) *TextureAtlas {
	ta := &TextureAtlas{
		texture:   tx,
		locations: make([]sdl.Rect, 0, len(locations)),
		names:     make(map[string]int, len(locations)),
	}

	for name, src := range locations {
		ta.names[name] = len(ta.locations)
		ta.locations = append(ta.locations, src)
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
		texture:   tx,
		locations: make([]sdl.Rect, 0, total),
		names:     make(map[string]int),
		uniform:   true,
	}

Loop:
	for y = 0; y < txH; y += cellH {
		for x = 0; x < txW; x += cellW {
			ta.locations = append(ta.locations, sdl.Rect{X: x, Y: y, W: cellW, H: cellH})

			total--
			if total <= 0 {
				break Loop
			}
		}
	}

	return ta, nil
}

func (ta *TextureAtlas) Texture() *sdl.Texture { return ta.texture }

func (ta *TextureAtlas) Len() int { return len(ta.locations) }

func (ta *TextureAtlas) IsUniform() bool { return ta.uniform }

func (ta *TextureAtlas) Names() []string {
	res := make([]string, 0, len(ta.names))
	for n := range ta.names {
		res = append(res, n)
	}
	return res
}

func (ta *TextureAtlas) HasIndex(i int) bool {
	return i >= 0 && i < len(ta.locations)
}

func (ta *TextureAtlas) GetFomIndex(i int) (TextureClip, error) {
	if !ta.HasIndex(i) {
		return TextureClip{}, errors.Newf("sdlkit: unknown index `%d` in TextureAtlas", i)
	}

	return TextureClip{
		Texture:  ta.texture,
		Location: ta.locations[i],
	}, nil
}

func (ta *TextureAtlas) HasName(name string) bool {
	_, ok := ta.names[name]
	return ok
}

func (ta *TextureAtlas) GetFromName(name string) (TextureClip, error) {
	i, ok := ta.names[name]
	if !ok {
		return TextureClip{}, errors.Newf("sdlkit: unknown name `%s` in TextureAtlas", name)
	}

	return ta.GetFomIndex(i)
}

func (ta *TextureAtlas) Destroy() error {
	err := ta.texture.Destroy()
	if errors.Is(err, ErrInvalidTexture) {
		err = nil // texture is probably already destroyed
	}
	return err
}

type TexturesMap map[string]*sdl.Texture

func (t TexturesMap) Destroy(name string) error {
	tx := t[name]
	if tx == nil {
		return nil
	}

	t[name] = nil
	return errors.Trace(tx.Destroy())
}

// Destroy destroys all sdl.Textures within the TexturesMap.
func (t TexturesMap) DestroyAll() error {
	var err error
	for n, tx := range t {
		if tx == nil {
			continue
		}

		errors.Append(&err, tx.Destroy())
		t[n] = nil
	}
	return err
}
