// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dump

import (
	"github.com/veandco/go-sdl2/sdl"

	cl "github.com/go-pogo/sdlkit/debug/constlookup"
)

func Event(event sdl.Event) string {
	d := newDumper(event)
	d.Const("Type", event.GetType(), cl.EventTypes.Lookup(event.GetType()), "unknown event type")
	d.Number("Timestamp", event.GetTimestamp())

	switch e := event.(type) {
	case *sdl.DisplayEvent:
		d.Number("Display", e.Display)
		d.Number("Event", e.Event)
		d.Number("Data1", e.Data1)

	case *sdl.WindowEvent:
		d.Const("Event", e.Event, cl.WindowEventTypes.Lookup(e.Event), "unknown window event")
		d.Number("WindowID", e.WindowID)
		if e.Data1 != 0 {
			d.Number("Data1", e.Data1)
		}
		if e.Data2 != 0 {
			d.Number("Data2", e.Data2)
		}

	case *sdl.KeyboardEvent:
		d.Number("WindowID", e.WindowID)
		d.Const("State", e.State, cl.PressedReleasedState.Lookup(e.State), "unknown key press state")
		d.Number("Repeat", e.Repeat)
		d.Struct("Keysym", keysymDumper(e.Keysym))

	case *sdl.TextEditingEvent:
		d.Number("WindowID", e.WindowID)
		d.String("Text", e.GetText())
		d.Number("Start", e.Start)
		d.Number("Length", e.Length)

	case *sdl.TextInputEvent:
		d.Number("WindowID", e.WindowID)
		d.String("Text", e.GetText())

	case *sdl.MouseMotionEvent:
		d.Number("WindowID", e.WindowID)
		d.Number("Which", e.Which)
		d.Const("State", e.State, cl.MouseButtons.Lookup(uint8(e.State)), "")
		d.Number("X", e.X)
		d.Number("Y", e.Y)
		d.Number("XRel", e.XRel)
		d.Number("YRel", e.YRel)

	case *sdl.MouseButtonEvent:
		d.Number("WindowID", e.WindowID)
		d.Number("Which", e.Which)
		d.Const("Button", e.Button, cl.MouseButtons.Lookup(e.Button), "unknown button")
		d.Const("State", e.State, cl.PressedReleasedState.Lookup(e.State), "unknown press state")
		d.Number("Clicks", e.Clicks)
		d.Number("X", e.X)
		d.Number("Y", e.Y)

	case *sdl.MouseWheelEvent:
		d.Number("WindowID", e.WindowID)
		d.Number("Which", e.Which)
		d.Number("X", e.X)
		d.Number("Y", e.Y)
		d.Const("Direction", e.Direction, cl.MouseWheelDirections.Lookup(e.Direction), "unknown mouse wheel direction")

	case *sdl.JoyAxisEvent:
		d.Number("Which", e.Which) // sdl.JoystickID
		d.Number("Axis", e.Axis)
		d.Number("Value", e.Value)

	case *sdl.JoyBallEvent:
		d.Number("Which", e.Which) // sdl.JoystickID
		d.Number("Ball", e.Ball)
		d.Number("XRel", e.XRel)
		d.Number("YRel", e.YRel)

	case *sdl.JoyHatEvent:
		d.Number("Which", e.Which) // sdl.JoystickID
		d.Number("Hat", e.Hat)
		d.Const("Value", e.Value, cl.JoyHats.Lookup(e.Value), "unknown joy hat")

	case *sdl.JoyButtonEvent:
		d.Number("Which", e.Which) // sdl.JoystickID
		d.Number("Button", e.Button)
		d.Const("State", e.State, cl.PressedReleasedState.Lookup(e.State), "unknown press state")

	case *sdl.JoyDeviceAddedEvent:
		d.Number("Which", e.Which) // sdl.JoystickID

	case *sdl.JoyDeviceRemovedEvent:
		d.Number("Which", e.Which) // sdl.JoystickID

	case *sdl.ControllerAxisEvent:
		d.Number("Which", e.Which) // sdl.JoystickID
		d.Const("Axis", e.Axis, cl.ControllerAxis.Lookup(e.Axis), "unknown controller axis")
		d.Number("Value", e.Value)

	case *sdl.ControllerButtonEvent:
		d.Number("Which", e.Which) // sdl.JoystickID
		d.Const("Button", e.Button, cl.ControllerButtons.Lookup(e.Button), "unknown controller button")
		d.Const("State", e.State, cl.PressedReleasedState.Lookup(e.State), "unknown press state")

	case *sdl.ControllerDeviceEvent:
		d.Number("Which", e.Which) // sdl.JoystickID

	case *sdl.AudioDeviceEvent:
		d.Number("Which", e.Which)
		d.Number("IsCapture", e.IsCapture)

	case *sdl.TouchFingerEvent:
		d.Number("TouchID", e.TouchID)   // sdl.TouchID
		d.Number("FingerID", e.FingerID) // sdl.FingerID
		d.Number("X", e.X)
		d.Number("Y", e.Y)
		d.Number("DX", e.DX)
		d.Number("DY", e.DY)
		d.Number("Pressure", e.Pressure)

	case *sdl.MultiGestureEvent:
		d.Number("TouchID", e.TouchID) // sdl.TouchID
		d.Number("DTheta", e.DTheta)
		d.Number("DDist", e.DDist)
		d.Number("X", e.X)
		d.Number("Y", e.Y)
		d.Number("NumFingers", e.NumFingers)

	case *sdl.DollarGestureEvent:
		d.Number("TouchID", e.TouchID)     // sdl.TouchID
		d.Number("GestureID", e.GestureID) // sdl.FingerID
		d.Number("NumFingers", e.NumFingers)
		d.Number("Error", e.Error)
		d.Number("X", e.X)
		d.Number("Y", e.Y)

	case *sdl.DropEvent:
		d.String("File", e.File)
		d.Number("WindowID", e.WindowID)

	case *sdl.SensorEvent:
		d.Number("Which", e.Which)
		// Data      [6]float32

	// these do not contain any additional fields:
	// case *sdl.RenderEvent:
	// case *sdl.QuitEvent:
	// case *sdl.OSEvent:
	// case *sdl.ClipboardEvent:

	case *sdl.UserEvent:
		d.Number("WindowID", e.WindowID)
		d.Number("Code", e.Code)
		d.Number("Data1", e.Data1)
		d.Number("Data2", e.Data2)

	case *sdl.SysWMEvent:
		// todo: dump *sdl.SysWMmsg
	}

	return dump(d)
}
