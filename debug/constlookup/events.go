// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package constlookup

import (
	"github.com/veandco/go-sdl2/sdl"
)

//goland:noinspection GoUnusedGlobalVariable,SpellCheckingInspection
var (
	EventTypes = eventTypes{
		sdl.FIRSTEVENT:               "sdl.FIRSTEVENT",
		sdl.QUIT:                     "sdl.QUIT",
		sdl.APP_TERMINATING:          "sdl.APP_TERMINATING",
		sdl.APP_LOWMEMORY:            "sdl.APP_LOWMEMORY",
		sdl.APP_WILLENTERBACKGROUND:  "sdl.APP_WILLENTERBACKGROUND",
		sdl.APP_DIDENTERBACKGROUND:   "sdl.APP_DIDENTERBACKGROUND",
		sdl.APP_WILLENTERFOREGROUND:  "sdl.APP_WILLENTERFOREGROUND",
		sdl.APP_DIDENTERFOREGROUND:   "sdl.APP_DIDENTERFOREGROUND",
		sdl.DISPLAYEVENT:             "sdl.DISPLAYEVENT",
		sdl.WINDOWEVENT:              "sdl.WINDOWEVENT",
		sdl.SYSWMEVENT:               "sdl.SYSWMEVENT",
		sdl.KEYDOWN:                  "sdl.KEYDOWN",
		sdl.KEYUP:                    "sdl.KEYUP",
		sdl.TEXTEDITING:              "sdl.TEXTEDITING",
		sdl.TEXTINPUT:                "sdl.TEXTINPUT",
		sdl.KEYMAPCHANGED:            "sdl.KEYMAPCHANGED",
		sdl.MOUSEMOTION:              "sdl.MOUSEMOTION",
		sdl.MOUSEBUTTONDOWN:          "sdl.MOUSEBUTTONDOWN",
		sdl.MOUSEBUTTONUP:            "sdl.MOUSEBUTTONUP",
		sdl.MOUSEWHEEL:               "sdl.MOUSEWHEEL",
		sdl.JOYAXISMOTION:            "sdl.JOYAXISMOTION",
		sdl.JOYBALLMOTION:            "sdl.JOYBALLMOTION",
		sdl.JOYHATMOTION:             "sdl.JOYHATMOTION",
		sdl.JOYBUTTONDOWN:            "sdl.JOYBUTTONDOWN",
		sdl.JOYBUTTONUP:              "sdl.JOYBUTTONUP",
		sdl.JOYDEVICEADDED:           "sdl.JOYDEVICEADDED",
		sdl.JOYDEVICEREMOVED:         "sdl.JOYDEVICEREMOVED",
		sdl.CONTROLLERAXISMOTION:     "sdl.CONTROLLERAXISMOTION",
		sdl.CONTROLLERBUTTONDOWN:     "sdl.CONTROLLERBUTTONDOWN",
		sdl.CONTROLLERBUTTONUP:       "sdl.CONTROLLERBUTTONUP",
		sdl.CONTROLLERDEVICEADDED:    "sdl.CONTROLLERDEVICEADDED",
		sdl.CONTROLLERDEVICEREMOVED:  "sdl.CONTROLLERDEVICEREMOVED",
		sdl.CONTROLLERDEVICEREMAPPED: "sdl.CONTROLLERDEVICEREMAPPED",
		sdl.FINGERDOWN:               "sdl.FINGERDOWN",
		sdl.FINGERUP:                 "sdl.FINGERUP",
		sdl.FINGERMOTION:             "sdl.FINGERMOTION",
		sdl.DOLLARGESTURE:            "sdl.DOLLARGESTURE",
		sdl.DOLLARRECORD:             "sdl.DOLLARRECORD",
		sdl.MULTIGESTURE:             "sdl.MULTIGESTURE",
		sdl.CLIPBOARDUPDATE:          "sdl.CLIPBOARDUPDATE",
		sdl.DROPFILE:                 "sdl.DROPFILE",
		sdl.DROPTEXT:                 "sdl.DROPTEXT",
		sdl.DROPBEGIN:                "sdl.DROPBEGIN",
		sdl.DROPCOMPLETE:             "sdl.DROPCOMPLETE",
		sdl.AUDIODEVICEADDED:         "sdl.AUDIODEVICEADDED",
		sdl.AUDIODEVICEREMOVED:       "sdl.AUDIODEVICEREMOVED",
		sdl.SENSORUPDATE:             "sdl.SENSORUPDATE",
		sdl.RENDER_TARGETS_RESET:     "sdl.RENDER_TARGETS_RESET",
		sdl.RENDER_DEVICE_RESET:      "sdl.RENDER_DEVICE_RESET",
		sdl.USEREVENT:                "sdl.USEREVENT",
		sdl.LASTEVENT:                "sdl.LASTEVENT",
	}

	WindowEventTypes = windowEventTypes{
		sdl.WINDOWEVENT_NONE:         "sdl.WINDOWEVENT_NONE",
		sdl.WINDOWEVENT_SHOWN:        "sdl.WINDOWEVENT_SHOWN",
		sdl.WINDOWEVENT_HIDDEN:       "sdl.WINDOWEVENT_HIDDEN",
		sdl.WINDOWEVENT_EXPOSED:      "sdl.WINDOWEVENT_EXPOSED",
		sdl.WINDOWEVENT_MOVED:        "sdl.WINDOWEVENT_MOVED",
		sdl.WINDOWEVENT_RESIZED:      "sdl.WINDOWEVENT_RESIZED",
		sdl.WINDOWEVENT_SIZE_CHANGED: "sdl.WINDOWEVENT_SIZE_CHANGED",
		sdl.WINDOWEVENT_MINIMIZED:    "sdl.WINDOWEVENT_MINIMIZED",
		sdl.WINDOWEVENT_MAXIMIZED:    "sdl.WINDOWEVENT_MAXIMIZED",
		sdl.WINDOWEVENT_RESTORED:     "sdl.WINDOWEVENT_RESTORED",
		sdl.WINDOWEVENT_ENTER:        "sdl.WINDOWEVENT_ENTER",
		sdl.WINDOWEVENT_LEAVE:        "sdl.WINDOWEVENT_LEAVE",
		sdl.WINDOWEVENT_FOCUS_GAINED: "sdl.WINDOWEVENT_FOCUS_GAINED",
		sdl.WINDOWEVENT_FOCUS_LOST:   "sdl.WINDOWEVENT_FOCUS_LOST",
		sdl.WINDOWEVENT_CLOSE:        "sdl.WINDOWEVENT_CLOSE",
		sdl.WINDOWEVENT_TAKE_FOCUS:   "sdl.WINDOWEVENT_TAKE_FOCUS",
		sdl.WINDOWEVENT_HIT_TEST:     "sdl.WINDOWEVENT_HIT_TEST",
	}
)

type eventTypes map[uint32]string

func (l eventTypes) Lookup(c uint32) string { return l[c] }

type windowEventTypes map[uint8]string

func (l windowEventTypes) Lookup(c uint8) string { return l[c] }
