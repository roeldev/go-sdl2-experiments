// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package constlookup

import (
	"github.com/veandco/go-sdl2/sdl"
)

//goland:noinspection GoUnusedGlobalVariable,SpellCheckingInspection
var (
	JoyHats = joyHats{
		sdl.HAT_CENTERED:  "sdl.HAT_CENTERED",
		sdl.HAT_UP:        "sdl.HAT_UP",
		sdl.HAT_RIGHT:     "sdl.HAT_RIGHT",
		sdl.HAT_DOWN:      "sdl.HAT_DOWN",
		sdl.HAT_LEFT:      "sdl.HAT_LEFT",
		sdl.HAT_RIGHTUP:   "sdl.HAT_RIGHTUP",
		sdl.HAT_RIGHTDOWN: "sdl.HAT_RIGHTDOWN",
		sdl.HAT_LEFTUP:    "sdl.HAT_LEFTUP",
		sdl.HAT_LEFTDOWN:  "sdl.HAT_LEFTDOWN",
	}

	JoystickTypes = joystickTypes{
		sdl.JOYSTICK_TYPE_UNKNOWN:        "sdl.JOYSTICK_TYPE_UNKNOWN",
		sdl.JOYSTICK_TYPE_GAMECONTROLLER: "sdl.JOYSTICK_TYPE_GAMECONTROLLER",
		sdl.JOYSTICK_TYPE_WHEEL:          "sdl.JOYSTICK_TYPE_WHEEL",
		sdl.JOYSTICK_TYPE_ARCADE_STICK:   "sdl.JOYSTICK_TYPE_ARCADE_STICK",
		sdl.JOYSTICK_TYPE_FLIGHT_STICK:   "sdl.JOYSTICK_TYPE_FLIGHT_STICK",
		sdl.JOYSTICK_TYPE_DANCE_PAD:      "sdl.JOYSTICK_TYPE_DANCE_PAD",
		sdl.JOYSTICK_TYPE_GUITAR:         "sdl.JOYSTICK_TYPE_GUITAR",
		sdl.JOYSTICK_TYPE_DRUM_KIT:       "sdl.JOYSTICK_TYPE_DRUM_KIT",
		sdl.JOYSTICK_TYPE_ARCADE_PAD:     "sdl.JOYSTICK_TYPE_ARCADE_PAD",
		sdl.JOYSTICK_TYPE_THROTTLE:       "sdl.JOYSTICK_TYPE_THROTTLE",
	}

	JoystickPowerLevels = joystickPowerLevels{
		sdl.JOYSTICK_POWER_UNKNOWN: "sdl.JOYSTICK_POWER_UNKNOWN",
		sdl.JOYSTICK_POWER_EMPTY:   "sdl.JOYSTICK_POWER_EMPTY",
		sdl.JOYSTICK_POWER_LOW:     "sdl.JOYSTICK_POWER_LOW",
		sdl.JOYSTICK_POWER_MEDIUM:  "sdl.JOYSTICK_POWER_MEDIUM",
		sdl.JOYSTICK_POWER_FULL:    "sdl.JOYSTICK_POWER_FULL",
		sdl.JOYSTICK_POWER_WIRED:   "sdl.JOYSTICK_POWER_WIRED",
		sdl.JOYSTICK_POWER_MAX:     "sdl.JOYSTICK_POWER_MAX",
	}
)

type joyHats map[uint8]string

func (l joyHats) Lookup(c uint8) string { return l[c] }

type joystickTypes map[uint8]string

func (l joystickTypes) Lookup(c uint8) string { return l[c] }

type joystickPowerLevels map[uint8]string

func (l joystickPowerLevels) Lookup(c uint8) string { return l[c] }
