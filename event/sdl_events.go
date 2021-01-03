// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package event

import (
	"github.com/veandco/go-sdl2/sdl"
)

type DisplayEventHandler interface {
	HandleDisplayEvent(e *sdl.DisplayEvent) error
}

type WindowEventHandler interface {
	HandleWindowEvent(e *sdl.WindowEvent) error
}

type KeyboardEventHandler interface {
	HandleKeyboardEvent(e *sdl.KeyboardEvent) error
}

type TextEditingEventHandler interface {
	HandleTextEditingEvent(e *sdl.TextEditingEvent) error
}

type TextInputEventHandler interface {
	HandleTextInputEvent(e *sdl.TextInputEvent) error
}

type MouseMotionEventHandler interface {
	HandleMouseMotionEvent(e *sdl.MouseMotionEvent) error
}

type MouseButtonEventHandler interface {
	HandleMouseButtonEvent(e *sdl.MouseButtonEvent) error
}

type MouseWheelEventHandler interface {
	HandleMouseWheelEvent(e *sdl.MouseWheelEvent) error
}

type JoyAxisEventHandler interface {
	HandleJoyAxisEvent(e *sdl.JoyAxisEvent) error
}

type JoyBallEventHandler interface {
	HandleJoyBallEvent(e *sdl.JoyBallEvent) error
}

type JoyHatEventHandler interface {
	HandleJoyHatEvent(e *sdl.JoyHatEvent) error
}

type JoyButtonEventHandler interface {
	HandleJoyButtonEvent(e *sdl.JoyButtonEvent) error
}

type JoyDeviceAddedEventHandler interface {
	HandleJoyDeviceAddedEvent(e *sdl.JoyDeviceAddedEvent) error
}

type JoyDeviceRemovedEventHandler interface {
	HandleJoyDeviceRemovedEvent(e *sdl.JoyDeviceRemovedEvent) error
}

type ControllerAxisEventHandler interface {
	HandleControllerAxisEvent(e *sdl.ControllerAxisEvent) error
}

type ControllerButtonEventHandler interface {
	HandleControllerButtonEvent(e *sdl.ControllerButtonEvent) error
}

type ControllerDeviceEventHandler interface {
	HandleControllerDeviceEvent(e *sdl.ControllerDeviceEvent) error
}

type AudioDeviceEventHandler interface {
	HandleAudioDeviceEvent(e *sdl.AudioDeviceEvent) error
}

type TouchFingerEventHandler interface {
	HandleTouchFingerEvent(e *sdl.TouchFingerEvent) error
}

type MultiGestureEventHandler interface {
	HandleMultiGestureEvent(e *sdl.MultiGestureEvent) error
}

type DollarGestureEventHandler interface {
	HandleDollarGestureEvent(e *sdl.DollarGestureEvent) error
}

type DropEventHandler interface {
	HandleDropEvent(e *sdl.DropEvent) error
}

type SensorEventHandler interface {
	HandleSensorEvent(e *sdl.SensorEvent) error
}

type RenderEventHandler interface {
	HandleRenderEvent(e *sdl.RenderEvent) error
}

// type QuitEventHandler interface {
// 	HandleQuitEvent(e *sdl.QuitEvent) error
// }

type OSEventHandler interface {
	HandleOSEvent(e *sdl.OSEvent) error
}

type ClipboardEventHandler interface {
	HandleClipboardEvent(e *sdl.ClipboardEvent) error
}

type UserEventHandler interface {
	HandleUserEvent(e *sdl.UserEvent) error
}

type SysWMEventHandler interface {
	HandleSysWMEvent(e *sdl.SysWMEvent) error
}
