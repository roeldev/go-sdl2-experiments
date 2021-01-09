// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdlkit

import (
	"encoding/xml"
	"io/fs"
	"path"

	"github.com/go-pogo/errors"
	sdlimg "github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	sdlttf "github.com/veandco/go-sdl2/ttf"
)

type AssetsLoader struct {
	fs  fs.ReadFileFS
	ren *sdl.Renderer

	// Surfaces
	Textures *TexturesMap
	Fonts    *FontsMap
}

func NewAssetsLoader(fs fs.ReadFileFS, r *sdl.Renderer) *AssetsLoader {
	return &AssetsLoader{
		fs:  fs,
		ren: r,
	}
}

func (l *AssetsLoader) Read(file string) ([]byte, error) {
	b, err := l.fs.ReadFile(file)
	return b, errors.Trace(err)
}

func (l *AssetsLoader) ReadRW(file string) (*sdl.RWops, error) {
	b, err := l.fs.ReadFile(file)
	if err != nil {
		return nil, errors.Trace(err)
	}

	rw, err := sdl.RWFromMem(b)
	return rw, errors.Trace(err)
}

func (l *AssetsLoader) Surface(file string) (*sdl.Surface, error) {
	src, err := l.ReadRW(file)
	if err != nil {
		return nil, errors.Trace(err)
	}

	sf, err := sdlimg.LoadRW(src, true)
	return sf, errors.Trace(err)
}

func LoadTextureFromMem(ren *sdl.Renderer, data []byte) (*sdl.Texture, error) {
	src, err := sdl.RWFromMem(data)
	if err != nil {
		return nil, errors.Trace(err)
	}

	tx, err := sdlimg.LoadTextureRW(ren, src, true)
	return tx, errors.Trace(err)
}

func (l *AssetsLoader) Texture(file string) (*sdl.Texture, error) {
	src, err := l.fs.ReadFile(file)
	if err != nil {
		return nil, errors.Trace(err)
	}

	tx, err := LoadTextureFromMem(l.ren, src)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if l.Textures != nil {
		l.Textures.Add(file, tx)
	}
	return tx, nil
}

func (l *AssetsLoader) TextureAtlas(file string, subs map[string]sdl.Rect) (*TextureAtlas, error) {
	tx, err := l.Texture(file)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return NewTextureAtlas(tx, subs), nil
}

func (l *AssetsLoader) TextureAtlasXml(file string) (*TextureAtlas, error) {
	data, err := l.fs.ReadFile(file)
	if err != nil {
		return nil, errors.Trace(err)
	}

	var x struct {
		File string `xml:"imagePath,attr"`
		Subs []struct {
			Name string `xml:"name,attr"`
			X    int32  `xml:"x,attr"`
			Y    int32  `xml:"y,attr"`
			W    int32  `xml:"width,attr"`
			H    int32  `xml:"height,attr"`
		} `xml:"SubTexture"`
	}

	if err = xml.Unmarshal(data, &x); err != nil {
		return nil, errors.Trace(err)
	}

	subs := make(map[string]sdl.Rect, len(x.Subs))
	for _, sub := range x.Subs {
		subs[sub.Name] = sdl.Rect{X: sub.X, Y: sub.Y, W: sub.W, H: sub.H}
	}

	a, err := l.TextureAtlas(path.Join(path.Dir(file), x.File), subs)
	return a, errors.Trace(err)
}

func (l *AssetsLoader) UniformTextureAtlas(file string, w, h int32, total uint8) (*TextureAtlas, error) {
	tx, err := l.Texture(file)
	if err != nil {
		return nil, errors.Trace(err)
	}

	a, err := NewUniformTextureAtlas(tx, w, h, total)
	return a, errors.Trace(err)
}

func (l *AssetsLoader) Font(file string, size int, index uint) (font *sdlttf.Font, err error) {
	src, err := l.ReadRW(file)
	if err != nil {
		return nil, errors.Trace(err)
	}

	if index < 0 {
		font, err = sdlttf.OpenFontRW(src, 1, size)
	} else {
		font, err = sdlttf.OpenFontIndexRW(src, 1, size, int(index))
	}
	if err != nil {
		return nil, errors.Trace(err)
	}
	if l.Fonts != nil {
		(*l.Fonts)[file] = font
	}
	return font, nil
}
