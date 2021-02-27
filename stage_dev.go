// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !prod

package sdlkit

import (
	"github.com/veandco/go-sdl2/sdl"
)

func init() {
	DefaultOptions.WindowTitleFps = true
	DefaultOptions.RendererFlags &^= sdl.RENDERER_PRESENTVSYNC
}
