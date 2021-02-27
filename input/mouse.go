// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package input

import (
	"sync"

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
	X, Y float64
	BtnLeft,
	BtnRight,
	BtnMiddle,
	BtnX1,
	BtnX2 *MouseBtnState

	mutex sync.RWMutex
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

func (ms *MouseState) Locker() *sync.RWMutex { return &ms.mutex }

func (ms *MouseState) GetX() float64 {
	ms.mutex.RLock()
	res := ms.X
	ms.mutex.RUnlock()
	return res
}

func (ms *MouseState) GetY() float64 {
	ms.mutex.RLock()
	res := ms.Y
	ms.mutex.RUnlock()
	return res
}

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
	ms.mutex.Lock()
	ms.X = float64(e.X)
	ms.Y = float64(e.Y)
	ms.mutex.Unlock()
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
	btn   uint8
	mutex sync.RWMutex
}

func (btn *MouseBtnState) Locker() *sync.RWMutex { return &btn.mutex }

func (btn *MouseBtnState) GetX() float64 {
	btn.mutex.RLock()
	res := btn.X
	btn.mutex.RUnlock()
	return res
}
func (btn *MouseBtnState) GetY() float64 {
	btn.mutex.RLock()
	res := btn.Y
	btn.mutex.RUnlock()
	return res
}

func (btn *MouseBtnState) updateMouseBtnState(e *sdl.MouseButtonEvent) {
	btn.mutex.Lock()
	btn.X = float64(e.X)
	btn.Y = float64(e.Y)
	btn.Pressed = e.State == sdl.PRESSED
	btn.Released = !btn.Pressed
	btn.Clicks = e.Clicks
	btn.mutex.Unlock()
}

func (btn *MouseBtnState) HandleMouseButtonEvent(e *sdl.MouseButtonEvent) error {
	btn.mutex.RLock()
	ok := e.Button == btn.btn
	btn.mutex.RUnlock()

	if ok {
		btn.updateMouseBtnState(e)
	}
	return nil
}
