// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdlkit

import (
	"github.com/go-pogo/errors"
	"github.com/veandco/go-sdl2/sdl"
)

type Renderable interface {
	Render(ren *sdl.Renderer) error
}

type RenderableFunc func(ren *sdl.Renderer) error

func (fn RenderableFunc) Render(ren *sdl.Renderer) error { return fn(ren) }

func Render(ren *sdl.Renderer, renderables ...Renderable) error {
	var err error
	for _, r := range renderables {
		errors.Append(&err, r.Render(ren))
	}
	return err
}

type Layer []Renderable

func (l *Layer) Clear() { *l = []Renderable(*l)[:0] }

func (l *Layer) Append(d ...Renderable) { *l = append(*l, d...) }

func (l *Layer) Prepend(d ...Renderable) {
	if len(*l) == 0 {
		l.Append(d...)
		return
	}

	slice := []Renderable(*l)
	slice = append(slice, d...)
	copy(slice[len(d):], slice)
	for i, drawable := range d {
		slice[i] = drawable
	}
	*l = slice
}

func (l Layer) Render(ren *sdl.Renderer) error { return Render(ren, l...) }
