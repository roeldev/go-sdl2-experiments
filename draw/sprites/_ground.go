// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/colors"
)

type ground struct {
	tx *sdl.Texture
}

func newGround(tx *sdl.Texture) *ground {
	green := colors.RandGreen(sdlkit.RNG())
	_ = tx.SetColorMod(green.R, green.G, green.B)

	return &ground{tx: tx}
}

func (g *ground) Draw(r *sdl.Renderer) error {
	return r.Copy(g.tx, nil, &sdl.Rect{W: 1024, H: 200})
}
