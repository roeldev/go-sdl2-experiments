// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tank

import (
	"github.com/veandco/go-sdl2/sdl"
)

type GameController struct {
	Device *sdl.GameController
	input  UserInput
}

func (gc *GameController) Input() UserInput { return gc.input }

func (gc *GameController) x() {

	// https://wiki.libsdl.org/SDL_GameControllerGetAxis
	// gc.Device.Axis()
}

type ControllerAxisEventHandler interface {
	HandleControllerAxisEvent(*sdl.ControllerAxisEvent) error
}

type ControllerButtonEventHandler interface {
	HandleControllerButtonEvent(*sdl.ControllerButtonEvent) error
}

type ControllerButtonDownEventHandler interface {
	HandleControllerButtonDownEvent(*sdl.ControllerButtonEvent) error
}

type ControllerButtonUpEventHandler interface {
	HandleControllerButtonUpEvent(*sdl.ControllerButtonEvent) error
}

type ControllerDeviceEventHandler interface {
	HandleControllerDeviceEvent(*sdl.ControllerDeviceEvent) error
}

type ControllerDeviceAddedEventHandler interface {
	HandleControllerDeviceAddedEvent(*sdl.ControllerDeviceEvent) error
}

type ControllerDeviceMappedEventHandler interface {
	HandleControllerDeviceMappedEvent(*sdl.ControllerDeviceEvent) error
}

type ControllerDeviceRemovedEventHandler interface {
	HandleControllerDeviceRemovedEvent(*sdl.ControllerDeviceEvent) error
}
