// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package constlookup

import "github.com/veandco/go-sdl2/sdl"

//goland:noinspection GoUnusedGlobalVariable,SpellCheckingInspection
var (
	// Window states.
	// (https://wiki.libsdl.org/SDL_WindowFlags)
	WindowFlags = windowFlags{
		sdl.WINDOW_FULLSCREEN:         "sdl.WINDOW_FULLSCREEN",
		sdl.WINDOW_OPENGL:             "sdl.WINDOW_OPENGL",
		sdl.WINDOW_SHOWN:              "sdl.WINDOW_SHOWN",
		sdl.WINDOW_HIDDEN:             "sdl.WINDOW_HIDDEN",
		sdl.WINDOW_BORDERLESS:         "sdl.WINDOW_BORDERLESS",
		sdl.WINDOW_RESIZABLE:          "sdl.WINDOW_RESIZABLE",
		sdl.WINDOW_MINIMIZED:          "sdl.WINDOW_MINIMIZED",
		sdl.WINDOW_MAXIMIZED:          "sdl.WINDOW_MAXIMIZED",
		sdl.WINDOW_INPUT_GRABBED:      "sdl.WINDOW_INPUT_GRABBED",
		sdl.WINDOW_INPUT_FOCUS:        "sdl.WINDOW_INPUT_FOCUS",
		sdl.WINDOW_MOUSE_FOCUS:        "sdl.WINDOW_MOUSE_FOCUS",
		sdl.WINDOW_FULLSCREEN_DESKTOP: "sdl.WINDOW_FULLSCREEN_DESKTOP",
		sdl.WINDOW_FOREIGN:            "sdl.WINDOW_FOREIGN",
		sdl.WINDOW_ALLOW_HIGHDPI:      "sdl.WINDOW_ALLOW_HIGHDPI",
		sdl.WINDOW_MOUSE_CAPTURE:      "sdl.WINDOW_MOUSE_CAPTURE",
		sdl.WINDOW_ALWAYS_ON_TOP:      "sdl.WINDOW_ALWAYS_ON_TOP",
		sdl.WINDOW_SKIP_TASKBAR:       "sdl.WINDOW_SKIP_TASKBAR",
		sdl.WINDOW_UTILITY:            "sdl.WINDOW_UTILITY",
		sdl.WINDOW_TOOLTIP:            "sdl.WINDOW_TOOLTIP",
		sdl.WINDOW_POPUP_MENU:         "sdl.WINDOW_POPUP_MENU",
		sdl.WINDOW_VULKAN:             "sdl.WINDOW_VULKAN",
	}
)

type windowFlags map[uint32]string

func (l windowFlags) Lookup(c uint32) string { return l[c] }
