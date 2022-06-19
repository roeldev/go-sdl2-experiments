// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package display

import (
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/geom"
)

type Display interface {
	sdlkit.Renderable
	Position() *geom.Point
}
