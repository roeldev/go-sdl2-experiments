// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdlkit

import (
	"github.com/veandco/go-sdl2/sdl"
)

type TrackMouseBtnState uint8

//goland:noinspection GoUnusedConst
const (
	TrackMouseBtnLeft TrackMouseBtnState = 1 << iota
	TrackMouseBtnRight
	TrackMouseBtnMiddle
	TrackMouseBtnX1
	TrackMouseBtnX2

	TrackAllMouseButtons = TrackMouseBtnLeft | TrackMouseBtnRight | TrackMouseBtnMiddle | TrackMouseBtnX1 | TrackMouseBtnX2
)

type MouseState struct {
	X, Y      float64
	BtnLeft   *MouseBtnState
	BtnRight  *MouseBtnState
	BtnMiddle *MouseBtnState
	BtnX1     *MouseBtnState
	BtnX2     *MouseBtnState
}

func NewMouseState(trackBtns TrackMouseBtnState) *MouseState {
	ms := &MouseState{}
	if trackBtns&TrackMouseBtnLeft != 0 {
		ms.BtnLeft = &MouseBtnState{btn: sdl.BUTTON_LEFT}
	}
	if trackBtns&TrackMouseBtnRight != 0 {
		ms.BtnRight = &MouseBtnState{btn: sdl.BUTTON_RIGHT}
	}
	if trackBtns&TrackMouseBtnMiddle != 0 {
		ms.BtnMiddle = &MouseBtnState{btn: sdl.BUTTON_MIDDLE}
	}
	if trackBtns&TrackMouseBtnX1 != 0 {
		ms.BtnX1 = &MouseBtnState{btn: sdl.BUTTON_X1}
	}
	if trackBtns&TrackMouseBtnX2 != 0 {
		ms.BtnX2 = &MouseBtnState{btn: sdl.BUTTON_X2}
	}
	return ms
}

func (ms *MouseState) GetX() float64 { return ms.X }
func (ms *MouseState) GetY() float64 { return ms.Y }

func (ms *MouseState) HandleMouseButtonEvent(e *sdl.MouseButtonEvent) error {
	switch e.Button {
	case sdl.BUTTON_LEFT:
		if ms.BtnLeft != nil {
			ms.BtnLeft.updateMouseBtnState(e)
		}
	case sdl.BUTTON_RIGHT:
		if ms.BtnRight != nil {
			ms.BtnRight.updateMouseBtnState(e)
		}
	case sdl.BUTTON_MIDDLE:
		if ms.BtnMiddle != nil {
			ms.BtnMiddle.updateMouseBtnState(e)
		}
	case sdl.BUTTON_X1:
		if ms.BtnX1 != nil {
			ms.BtnX1.updateMouseBtnState(e)
		}
	case sdl.BUTTON_X2:
		if ms.BtnX2 != nil {
			ms.BtnX2.updateMouseBtnState(e)
		}
	}
	return nil
}

func (ms *MouseState) HandleMouseMotionEvent(e *sdl.MouseMotionEvent) error {
	ms.X = float64(e.X)
	ms.Y = float64(e.Y)
	return nil
}

type MouseBtnState struct {
	X, Y     float64
	Pressed  bool
	Released bool
	Clicks   uint8

	// one of
	// - sdl.BUTTON_LEFT
	// - sdl.BUTTON_MIDDLE
	// - sdl.BUTTON_RIGHT
	// - sdl.BUTTON_X1
	// - sdl.BUTTON_X2
	btn uint8
}

func (btn *MouseBtnState) GetX() float64 { return btn.X }
func (btn *MouseBtnState) GetY() float64 { return btn.Y }

func (btn *MouseBtnState) updateMouseBtnState(e *sdl.MouseButtonEvent) {
	btn.X = float64(e.X)
	btn.Y = float64(e.Y)
	btn.Pressed = e.State == sdl.PRESSED
	btn.Released = !btn.Pressed
	btn.Clicks = e.Clicks
}

func (btn *MouseBtnState) HandleMouseButtonEvent(e *sdl.MouseButtonEvent) error {
	if e.Button == btn.btn {
		btn.updateMouseBtnState(e)
	}
	return nil
}
