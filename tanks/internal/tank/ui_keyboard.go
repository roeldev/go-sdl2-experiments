// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tank

import (
	"github.com/veandco/go-sdl2/sdl"
)

type KeyboardMouse struct {
	input UserInput
}

func (kc *KeyboardMouse) Input() UserInput { return kc.input }

func (kc *KeyboardMouse) HandleKeyboardEvent(e *sdl.KeyboardEvent) error {
	value := 0.0
	if e.State == sdl.PRESSED {
		value = 1
	}

	switch e.Keysym.Sym {
	case sdl.K_UP:
		fallthrough
	case sdl.K_w:
		kc.input.Forwards = value

	case sdl.K_DOWN:
		fallthrough
	case sdl.K_s:
		kc.input.Backwards = value

	case sdl.K_LEFT:
		fallthrough
	case sdl.K_a:
		kc.input.SteerLeft = value

	case sdl.K_RIGHT:
		fallthrough
	case sdl.K_d:
		kc.input.SteerRight = value

	case sdl.K_SPACE:
		kc.input.Brake = value
	}

	return nil
}
