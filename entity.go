// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdlkit

import (
	"github.com/go-pogo/errors"
	"github.com/veandco/go-sdl2/sdl"
)

type Entity interface {
	Drawable
	Update(clock *Clock)
}

type Drawable interface {
	Draw(r *sdl.Renderer) error
}

type DrawableFunc func(r *sdl.Renderer) error

func (d DrawableFunc) Draw(r *sdl.Renderer) error { return d(r) }

type Layer []Drawable

func (l *Layer) Append(d Drawable) { *l = append(*l, d) }

func (l Layer) Draw(r *sdl.Renderer) error {
	var err error
	for _, e := range l {
		errors.Append(&err, e.Draw(r))
	}
	return err
}
