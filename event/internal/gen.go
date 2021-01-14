// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"log"
	"path"
	"reflect"
	"runtime"
	"text/template"

	"github.com/veandco/go-sdl2/sdl"
)

var handlers = []eventHandler{
	{
		Event: new(sdl.DisplayEvent),
		Main:  "Display",
	},
	{
		Event: new(sdl.WindowEvent),
		Main:  "Window",
		// https://wiki.libsdl.org/SDL_WindowEventID
		Subs: map[string]map[string]string{
			"Event": {
				"sdl.WINDOWEVENT_SHOWN":        "WindowShown",
				"sdl.WINDOWEVENT_HIDDEN":       "WindowHidden",
				"sdl.WINDOWEVENT_EXPOSED":      "WindowExposed",
				"sdl.WINDOWEVENT_MOVED":        "WindowMoved",
				"sdl.WINDOWEVENT_RESIZED":      "WindowResized",
				"sdl.WINDOWEVENT_SIZE_CHANGED": "WindowSizeChanged",
				"sdl.WINDOWEVENT_MINIMIZED":    "WindowMinimized",
				"sdl.WINDOWEVENT_MAXIMIZED":    "WindowMaximized",
				"sdl.WINDOWEVENT_RESTORED":     "WindowRestored",
				"sdl.WINDOWEVENT_ENTER":        "WindowEnter",
				"sdl.WINDOWEVENT_LEAVE":        "WindowLeave",
				"sdl.WINDOWEVENT_FOCUS_GAINED": "WindowFocusGained",
				"sdl.WINDOWEVENT_FOCUS_LOST":   "WindowFocusLost",
				"sdl.WINDOWEVENT_CLOSE":        "WindowClose",
				"sdl.WINDOWEVENT_TAKE_FOCUS":   "WindowTakeFocus",
				"sdl.WINDOWEVENT_HIT_TEST":     "WindowHitTest",
			},
		},
	},
	{
		Event: new(sdl.KeyboardEvent),
		Main:  "Keyboard",
		// https://wiki.libsdl.org/SDL_KeyboardEvent
		Subs: map[string]map[string]string{
			"Type": {
				"sdl.KEYDOWN": "KeyDown",
				"sdl.KEYUP":   "KeyUp",
			},
		},
	},
	{
		Event: new(sdl.TextEditingEvent),
		Main:  "TextEditing",
	},
	{
		Event: new(sdl.TextInputEvent),
		Main:  "TextInput",
	},
	{
		Event: new(sdl.MouseMotionEvent),
		Main:  "MouseMotion",
	},
	{
		Event: new(sdl.MouseButtonEvent),
		Main:  "MouseButton",
		Subs: map[string]map[string]string{
			"Type": {
				"sdl.MOUSEBUTTONDOWN": "MouseButtonDown",
				"sdl.MOUSEBUTTONUP":   "MouseButtonUp",
			},
		},
	},
	{
		Event: new(sdl.MouseWheelEvent),
		Main:  "MouseWheel",
	},
	{
		Event: new(sdl.JoyAxisEvent),
		Main:  "JoyAxis",
	},
	{
		Event: new(sdl.JoyBallEvent),
		Main:  "JoyBall",
	},
	{
		Event: new(sdl.JoyHatEvent),
		Main:  "JoyHat",
	},
	{
		Event: new(sdl.JoyButtonEvent),
		Main:  "JoyButton",
		Subs: map[string]map[string]string{
			"Type": {
				"sdl.JOYBUTTONDOWN": "JoyButtonDown",
				"sdl.JOYBUTTONUP":   "JoyButtonUp",
			},
		},
	},
	{
		Event: new(sdl.JoyDeviceAddedEvent),
		Main:  "JoyDeviceAdded",
	},
	{
		Event: new(sdl.JoyDeviceRemovedEvent),
		Main:  "JoyDeviceRemoved",
	},
	{
		Event: new(sdl.ControllerAxisEvent),
		Main:  "ControllerAxis",
	},
	{
		Event: new(sdl.ControllerButtonEvent),
		Main:  "ControllerButton",
		Subs: map[string]map[string]string{
			"Type": {
				"sdl.CONTROLLERBUTTONDOWN": "ControllerButtonDown",
				"sdl.CONTROLLERBUTTONUP":   "ControllerButtonUp",
			},
		},
	},
	{
		Event: new(sdl.ControllerDeviceEvent),
		Main:  "ControllerDevice",
		Subs: map[string]map[string]string{
			"Type": {
				"sdl.CONTROLLERDEVICEADDED":    "ControllerDeviceAdded",
				"sdl.CONTROLLERDEVICEREMOVED":  "ControllerDeviceRemoved",
				"sdl.CONTROLLERDEVICEREMAPPED": "ControllerDeviceMapped",
			},
		},
	},
	{
		Event: new(sdl.AudioDeviceEvent),
		Main:  "AudioDevice",
		Subs: map[string]map[string]string{
			"Type": {
				"sdl.AUDIODEVICEADDED":   "AudioDeviceAdded",
				"sdl.AUDIODEVICEREMOVED": "AudioDeviceRemoved",
			},
		},
	},
	{
		Event: new(sdl.TouchFingerEvent),
		Main:  "TouchFinger",
		Subs: map[string]map[string]string{
			"Type": {
				"sdl.FINGERMOTION": "TouchFingerMotion",
				"sdl.FINGERDOWN":   "TouchFingerDown",
				"sdl.FINGERUP":     "TouchFingerUp",
			},
		},
	},
	{
		Event: new(sdl.MultiGestureEvent),
		Main:  "MultiGesture",
	},
	{
		Event: new(sdl.DollarGestureEvent),
		Main:  "DollarGesture",
	},
	{
		Event: new(sdl.DropEvent),
		Main:  "Drop",
	},
	{
		Event: new(sdl.SensorEvent),
		Main:  "Sensor",
	},
	{
		Event: new(sdl.RenderEvent),
		Main:  "Render",
	},
	// {
	// 	Event: new(sdl.QuitEvent),
	// 	Main:  "Quit",
	// },
	{
		Event: new(sdl.OSEvent),
		Main:  "OS",
	},
	{
		Event: new(sdl.ClipboardEvent),
		Main:  "Clipboard",
	},
	{
		Event: new(sdl.UserEvent),
		Main:  "User",
	},
	{
		Event: new(sdl.SysWMEvent),
		Main:  "SysWM",
	},
}

func main() {
	_, dir, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("unable to resolve current dir location")
	}

	vars := tmplVars{
		Cases:    handlers,
		Handlers: make(map[string]sdl.Event, len(handlers)*2),
	}
	for _, h := range handlers {
		vars.Handlers[h.Main] = h.Event
		for _, subs := range h.Subs {
			for _, sh := range subs {
				vars.Handlers[sh] = h.Event
			}
		}
	}

	tmpl := template.New("").Funcs(template.FuncMap{
		"eventName":       eventName,
		"interfaceName":   interfaceName,
		"handlerFuncName": handlerFuncName,
	})

	var buf bytes.Buffer
	if err := template.Must(tmpl.Parse(handlerFileSrc)).Execute(&buf, vars); err != nil {
		log.Fatal("error while executing template:", err)
	}

	data, err := format.Source(buf.Bytes())
	if err != nil {
		log.Println("error while formatting code:", err)
		data = buf.Bytes()
	}

	filename := path.Join(path.Dir(path.Dir(dir)), "handlers.go")
	if err := ioutil.WriteFile(filename, data, 0644); err != nil {
		log.Fatal("error writing", filename+":", err)
	}
}

type eventHandler struct {
	Event sdl.Event
	Main  string
	Subs  map[string]map[string]string
}

type tmplVars struct {
	Cases    []eventHandler
	Handlers map[string]sdl.Event
}

func eventName(event sdl.Event) string { return reflect.TypeOf(event).String() }

func interfaceName(str string) string { return str + "EventHandler" }

func handlerFuncName(str string) string { return "Handle" + str + "Event" }

const handlerFileSrc = `// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// generated by go generate; DO NOT EDIT.

package event

import "github.com/veandco/go-sdl2/sdl"

{{ range $name, $event := .Handlers }}
type {{ interfaceName $name }} interface {
	{{ handlerFuncName $name }}({{ eventName $event }}) error
}
{{ end }}

type handlers struct {
	{{ range $name, $_ := .Handlers }}
	{{ $name }} []{{ interfaceName $name }}{{ end }}
}

func (h *handlers) register(handlers ...interface{}) {
	for _, handler := range handlers {
		{{ range $name, $_ := .Handlers }}
		if v, ok := handler.({{ interfaceName $name }}); ok {
			h.{{ $name }} = append(h.{{ $name }}, v)
		} {{ end }}
	}
}

func (h *handlers) handle(event sdl.Event) error {
	var err error
	switch e := event.(type) {
		{{ range $_, $case := .Cases }}
		case {{ eventName $case.Event }}:
		for _, x := range h.{{ $case.Main }} {
			errors.Append(&err, x.{{ handlerFuncName $case.Main }}(e))
		}

		{{ range $p, $subs := $case.Subs }}
		switch e.{{ $p }} { {{ range $cn, $sh := $subs }}
			case {{ $cn }}:
			for _, x := range h.{{ $sh }} {
				errors.Append(&err, x.{{ handlerFuncName $sh }}(e))
			} {{ end }}
		} {{ end }}
		{{ end }}
	}

	return err
}
`
