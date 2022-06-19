// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package display

import (
	"github.com/veandco/go-sdl2/sdl"
)

type ColorTransform struct {
}

type TextureTransform struct {
	Rotation,
	ScaleX, ScaleY,
	TranslateX, TranslateY float64
	Alpha     uint8
	BlendMode sdl.BlendMode
	Flip      sdl.RendererFlip
}

func (tt *TextureTransform) Reset() {
	tt.Alpha = 1
	tt.BlendMode = sdl.BLENDMODE_BLEND
	tt.Flip = sdl.FLIP_NONE
}

func ResetTextureDisplay(td *TextureTransform) {
	td.Alpha = 1
	td.BlendMode = sdl.BLENDMODE_BLEND
	td.Flip = sdl.FLIP_NONE
}
