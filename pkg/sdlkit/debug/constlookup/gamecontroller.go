// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package constlookup

import "github.com/veandco/go-sdl2/sdl"

//goland:noinspection GoUnusedGlobalVariable,SpellCheckingInspection
var (
	ControllerAxis = controllerAxis{
		sdl.CONTROLLER_AXIS_INVALID:      "sdl.CONTROLLER_AXIS_INVALID",
		sdl.CONTROLLER_AXIS_LEFTX:        "sdl.CONTROLLER_AXIS_LEFTX",
		sdl.CONTROLLER_AXIS_LEFTY:        "sdl.CONTROLLER_AXIS_LEFTY",
		sdl.CONTROLLER_AXIS_RIGHTX:       "sdl.CONTROLLER_AXIS_RIGHTX",
		sdl.CONTROLLER_AXIS_RIGHTY:       "sdl.CONTROLLER_AXIS_RIGHTY",
		sdl.CONTROLLER_AXIS_TRIGGERLEFT:  "sdl.CONTROLLER_AXIS_TRIGGERLEFT",
		sdl.CONTROLLER_AXIS_TRIGGERRIGHT: "sdl.CONTROLLER_AXIS_TRIGGERRIGHT",
		sdl.CONTROLLER_AXIS_MAX:          "sdl.CONTROLLER_AXIS_MAX",
	}

	ControllerButtons = controllerButtons{
		sdl.CONTROLLER_BUTTON_INVALID:       "sdl.CONTROLLER_BUTTON_INVALID",
		sdl.CONTROLLER_BUTTON_A:             "sdl.CONTROLLER_BUTTON_A",
		sdl.CONTROLLER_BUTTON_B:             "sdl.CONTROLLER_BUTTON_B",
		sdl.CONTROLLER_BUTTON_X:             "sdl.CONTROLLER_BUTTON_X",
		sdl.CONTROLLER_BUTTON_Y:             "sdl.CONTROLLER_BUTTON_Y",
		sdl.CONTROLLER_BUTTON_BACK:          "sdl.CONTROLLER_BUTTON_BACK",
		sdl.CONTROLLER_BUTTON_GUIDE:         "sdl.CONTROLLER_BUTTON_GUIDE",
		sdl.CONTROLLER_BUTTON_START:         "sdl.CONTROLLER_BUTTON_START",
		sdl.CONTROLLER_BUTTON_LEFTSTICK:     "sdl.CONTROLLER_BUTTON_LEFTSTICK",
		sdl.CONTROLLER_BUTTON_RIGHTSTICK:    "sdl.CONTROLLER_BUTTON_RIGHTSTICK",
		sdl.CONTROLLER_BUTTON_LEFTSHOULDER:  "sdl.CONTROLLER_BUTTON_LEFTSHOULDER",
		sdl.CONTROLLER_BUTTON_RIGHTSHOULDER: "sdl.CONTROLLER_BUTTON_RIGHTSHOULDER",
		sdl.CONTROLLER_BUTTON_DPAD_UP:       "sdl.CONTROLLER_BUTTON_DPAD_UP",
		sdl.CONTROLLER_BUTTON_DPAD_DOWN:     "sdl.CONTROLLER_BUTTON_DPAD_DOWN",
		sdl.CONTROLLER_BUTTON_DPAD_LEFT:     "sdl.CONTROLLER_BUTTON_DPAD_LEFT",
		sdl.CONTROLLER_BUTTON_DPAD_RIGHT:    "sdl.CONTROLLER_BUTTON_DPAD_RIGHT",
		sdl.CONTROLLER_BUTTON_MAX:           "sdl.CONTROLLER_BUTTON_MAX",
	}
)

type controllerAxis map[int8]string

func (l controllerAxis) Lookup(c int8) string { return l[c] }

type controllerButtons map[int8]string

func (l controllerButtons) Lookup(c int8) string { return l[c] }
