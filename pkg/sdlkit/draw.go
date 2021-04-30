// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdlkit

import (
	"github.com/go-pogo/errors"
	"github.com/veandco/go-sdl2/sdl"
)

func Draw(renderer *sdl.Renderer, draw ...Drawable) error {
	var err error
	for _, d := range draw {
		errors.Append(&err, d.Draw(renderer))
	}
	return err
}

type Drawable interface {
	Draw(r *sdl.Renderer) error
}

type DrawableFunc func(r *sdl.Renderer) error

func (d DrawableFunc) Draw(r *sdl.Renderer) error { return d(r) }

type Layer []Drawable

func (l *Layer) Append(d ...Drawable) { *l = append(*l, d...) }

func (l *Layer) Prepend(d ...Drawable) {
	if len(*l) == 0 {
		l.Append(d...)
		return
	}

	slice := []Drawable(*l)
	slice = append(slice, d...)
	copy(slice[len(d):], slice)
	for i, d := range d {
		slice[i] = d
	}
	*l = slice
}

func (l Layer) Draw(r *sdl.Renderer) error { return Draw(r, l...) }
