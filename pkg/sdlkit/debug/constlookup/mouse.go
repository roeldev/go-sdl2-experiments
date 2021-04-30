// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package constlookup

import "github.com/veandco/go-sdl2/sdl"

//goland:noinspection GoUnusedGlobalVariable,SpellCheckingInspection
var (
	PressedReleasedState = pressedReleasedState{
		sdl.RELEASED: "sdl.RELEASED",
		sdl.PRESSED:  "sdl.PRESSED",
	}
	MouseButtons = mouseButtons{
		sdl.BUTTON_LEFT:   "sdl.BUTTON_LEFT",
		sdl.BUTTON_MIDDLE: "sdl.BUTTON_MIDDLE",
		sdl.BUTTON_RIGHT:  "sdl.BUTTON_RIGHT",
		sdl.BUTTON_X1:     "sdl.BUTTON_X1",
		sdl.BUTTON_X2:     "sdl.BUTTON_X2",
	}
	MouseWheelDirections = mouseWheelDirections{
		sdl.MOUSEWHEEL_NORMAL:  "sdl.MOUSEWHEEL_NORMAL",
		sdl.MOUSEWHEEL_FLIPPED: "sdl.MOUSEWHEEL_FLIPPED",
	}
)

type pressedReleasedState map[uint8]string

func (l pressedReleasedState) Lookup(c uint8) string { return l[c] }

type mouseButtons map[uint8]string

func (l mouseButtons) Lookup(c uint8) string { return l[c] }

type mouseWheelDirections map[uint32]string

func (l mouseWheelDirections) Lookup(c uint32) string { return l[c] }
