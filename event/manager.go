// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package event

import (
	"log"

	"github.com/go-pogo/errors"
	"github.com/veandco/go-sdl2/sdl"
)

type Manager struct {
	displayEvent          []DisplayEventHandler
	windowEvent           []WindowEventHandler
	keyboardEvent         []KeyboardEventHandler
	textEditingEvent      []TextEditingEventHandler
	textInputEvent        []TextInputEventHandler
	mouseMotionEvent      []MouseMotionEventHandler
	mouseButtonEvent      []MouseButtonEventHandler
	mouseWheelEvent       []MouseWheelEventHandler
	joyAxisEvent          []JoyAxisEventHandler
	joyBallEvent          []JoyBallEventHandler
	joyHatEvent           []JoyHatEventHandler
	joyButtonEvent        []JoyButtonEventHandler
	joyDeviceAddedEvent   []JoyDeviceAddedEventHandler
	joyDeviceRemovedEvent []JoyDeviceRemovedEventHandler
	controllerAxisEvent   []ControllerAxisEventHandler
	controllerButtonEvent []ControllerButtonEventHandler
	controllerDeviceEvent []ControllerDeviceEventHandler
	audioDeviceEvent      []AudioDeviceEventHandler
	touchFingerEvent      []TouchFingerEventHandler
	multiGestureEvent     []MultiGestureEventHandler
	dollarGestureEvent    []DollarGestureEventHandler
	dropEvent             []DropEventHandler
	sensorEvent           []SensorEventHandler
	renderEvent           []RenderEventHandler
	// quitEvent             []QuitEventHandler
	oSEvent        []OSEventHandler
	clipboardEvent []ClipboardEventHandler
	userEvent      []UserEventHandler
	sysWMEvent     []SysWMEventHandler
}

func (m *Manager) Register(handler ...interface{}) {
	for _, h := range handler {
		if v, ok := h.(DisplayEventHandler); ok {
			m.displayEvent = append(m.displayEvent, v)
		}
		if v, ok := h.(WindowEventHandler); ok {
			m.windowEvent = append(m.windowEvent, v)
		}
		if v, ok := h.(KeyboardEventHandler); ok {
			m.keyboardEvent = append(m.keyboardEvent, v)
		}
		if v, ok := h.(TextEditingEventHandler); ok {
			m.textEditingEvent = append(m.textEditingEvent, v)
		}
		if v, ok := h.(TextInputEventHandler); ok {
			m.textInputEvent = append(m.textInputEvent, v)
		}
		if v, ok := h.(MouseMotionEventHandler); ok {
			m.mouseMotionEvent = append(m.mouseMotionEvent, v)
		}
		if v, ok := h.(MouseButtonEventHandler); ok {
			m.mouseButtonEvent = append(m.mouseButtonEvent, v)
		}
		if v, ok := h.(MouseWheelEventHandler); ok {
			m.mouseWheelEvent = append(m.mouseWheelEvent, v)
		}
		if v, ok := h.(JoyAxisEventHandler); ok {
			m.joyAxisEvent = append(m.joyAxisEvent, v)
		}
		if v, ok := h.(JoyBallEventHandler); ok {
			m.joyBallEvent = append(m.joyBallEvent, v)
		}
		if v, ok := h.(JoyHatEventHandler); ok {
			m.joyHatEvent = append(m.joyHatEvent, v)
		}
		if v, ok := h.(JoyButtonEventHandler); ok {
			m.joyButtonEvent = append(m.joyButtonEvent, v)
		}
		if v, ok := h.(JoyDeviceAddedEventHandler); ok {
			m.joyDeviceAddedEvent = append(m.joyDeviceAddedEvent, v)
		}
		if v, ok := h.(JoyDeviceRemovedEventHandler); ok {
			m.joyDeviceRemovedEvent = append(m.joyDeviceRemovedEvent, v)
		}
		if v, ok := h.(ControllerAxisEventHandler); ok {
			m.controllerAxisEvent = append(m.controllerAxisEvent, v)
		}
		if v, ok := h.(ControllerButtonEventHandler); ok {
			m.controllerButtonEvent = append(m.controllerButtonEvent, v)
		}
		if v, ok := h.(ControllerDeviceEventHandler); ok {
			m.controllerDeviceEvent = append(m.controllerDeviceEvent, v)
		}
		if v, ok := h.(AudioDeviceEventHandler); ok {
			m.audioDeviceEvent = append(m.audioDeviceEvent, v)
		}
		if v, ok := h.(TouchFingerEventHandler); ok {
			m.touchFingerEvent = append(m.touchFingerEvent, v)
		}
		if v, ok := h.(MultiGestureEventHandler); ok {
			m.multiGestureEvent = append(m.multiGestureEvent, v)
		}
		if v, ok := h.(DollarGestureEventHandler); ok {
			m.dollarGestureEvent = append(m.dollarGestureEvent, v)
		}
		if v, ok := h.(DropEventHandler); ok {
			m.dropEvent = append(m.dropEvent, v)
		}
		if v, ok := h.(SensorEventHandler); ok {
			m.sensorEvent = append(m.sensorEvent, v)
		}
		if v, ok := h.(RenderEventHandler); ok {
			m.renderEvent = append(m.renderEvent, v)
		}
		// if v, ok := v.(QuitEventHandler); ok {
		// 	m.quitEvent = append(m.quitEvent, v)
		// }
		if v, ok := h.(OSEventHandler); ok {
			m.oSEvent = append(m.oSEvent, v)
		}
		if v, ok := h.(ClipboardEventHandler); ok {
			m.clipboardEvent = append(m.clipboardEvent, v)
		}
		if v, ok := h.(UserEventHandler); ok {
			m.userEvent = append(m.userEvent, v)
		}
		if v, ok := h.(SysWMEventHandler); ok {
			m.sysWMEvent = append(m.sysWMEvent, v)
		}
	}
}

func (m *Manager) Handle(event sdl.Event) (err error) {
	switch e := event.(type) {
	case *sdl.QuitEvent:
		log.Println("QUIT", e)
		return Quit

	case *sdl.DisplayEvent:
		for _, h := range m.displayEvent {
			errors.Append(&err, h.HandleDisplayEvent(e))
		}

	case *sdl.WindowEvent:
		for _, h := range m.windowEvent {
			errors.Append(&err, h.HandleWindowEvent(e))
		}

	case *sdl.KeyboardEvent:
		for _, h := range m.keyboardEvent {
			errors.Append(&err, h.HandleKeyboardEvent(e))
		}

	case *sdl.TextEditingEvent:
		for _, h := range m.textEditingEvent {
			errors.Append(&err, h.HandleTextEditingEvent(e))
		}

	case *sdl.TextInputEvent:
		for _, h := range m.textInputEvent {
			errors.Append(&err, h.HandleTextInputEvent(e))
		}

	case *sdl.MouseMotionEvent:
		for _, h := range m.mouseMotionEvent {
			errors.Append(&err, h.HandleMouseMotionEvent(e))
		}

	case *sdl.MouseButtonEvent:
		for _, h := range m.mouseButtonEvent {
			errors.Append(&err, h.HandleMouseButtonEvent(e))
		}

	case *sdl.MouseWheelEvent:
		for _, h := range m.mouseWheelEvent {
			errors.Append(&err, h.HandleMouseWheelEvent(e))
		}

	case *sdl.JoyAxisEvent:
		for _, h := range m.joyAxisEvent {
			errors.Append(&err, h.HandleJoyAxisEvent(e))
		}

	case *sdl.JoyBallEvent:
		for _, h := range m.joyBallEvent {
			errors.Append(&err, h.HandleJoyBallEvent(e))
		}

	case *sdl.JoyHatEvent:
		for _, h := range m.joyHatEvent {
			errors.Append(&err, h.HandleJoyHatEvent(e))
		}

	case *sdl.JoyButtonEvent:
		for _, h := range m.joyButtonEvent {
			errors.Append(&err, h.HandleJoyButtonEvent(e))
		}

	case *sdl.JoyDeviceAddedEvent:
		for _, h := range m.joyDeviceAddedEvent {
			errors.Append(&err, h.HandleJoyDeviceAddedEvent(e))
		}

	case *sdl.JoyDeviceRemovedEvent:
		for _, h := range m.joyDeviceRemovedEvent {
			errors.Append(&err, h.HandleJoyDeviceRemovedEvent(e))
		}

	case *sdl.ControllerAxisEvent:
		for _, h := range m.controllerAxisEvent {
			errors.Append(&err, h.HandleControllerAxisEvent(e))
		}

	case *sdl.ControllerButtonEvent:
		for _, h := range m.controllerButtonEvent {
			errors.Append(&err, h.HandleControllerButtonEvent(e))
		}

	case *sdl.ControllerDeviceEvent:
		for _, h := range m.controllerDeviceEvent {
			errors.Append(&err, h.HandleControllerDeviceEvent(e))
		}

	case *sdl.AudioDeviceEvent:
		for _, h := range m.audioDeviceEvent {
			errors.Append(&err, h.HandleAudioDeviceEvent(e))
		}

	case *sdl.TouchFingerEvent:
		for _, h := range m.touchFingerEvent {
			errors.Append(&err, h.HandleTouchFingerEvent(e))
		}

	case *sdl.MultiGestureEvent:
		for _, h := range m.multiGestureEvent {
			errors.Append(&err, h.HandleMultiGestureEvent(e))
		}

	case *sdl.DollarGestureEvent:
		for _, h := range m.dollarGestureEvent {
			errors.Append(&err, h.HandleDollarGestureEvent(e))
		}

	case *sdl.DropEvent:
		for _, h := range m.dropEvent {
			errors.Append(&err, h.HandleDropEvent(e))
		}

	case *sdl.SensorEvent:
		for _, h := range m.sensorEvent {
			errors.Append(&err, h.HandleSensorEvent(e))
		}

	case *sdl.RenderEvent:
		for _, h := range m.renderEvent {
			errors.Append(&err, h.HandleRenderEvent(e))
		}

	// case *sdl.QuitEvent:
	// 	for _, h := range m.quitEvent {
	// 		errors.Append(&err, h.HandleQuitEvent(e))
	// 	}

	case *sdl.OSEvent:
		for _, h := range m.oSEvent {
			errors.Append(&err, h.HandleOSEvent(e))
		}

	case *sdl.ClipboardEvent:
		for _, h := range m.clipboardEvent {
			errors.Append(&err, h.HandleClipboardEvent(e))
		}

	case *sdl.UserEvent:
		for _, h := range m.userEvent {
			errors.Append(&err, h.HandleUserEvent(e))
		}

	case *sdl.SysWMEvent:
		for _, h := range m.sysWMEvent {
			errors.Append(&err, h.HandleSysWMEvent(e))
		}
	}

	return
}

func (m *Manager) Loop() error {
	for {
		event := sdl.PollEvent()
		if event == nil {
			return nil
		}
		if err := m.Handle(event); err != nil {
			return err
		}
	}
}
