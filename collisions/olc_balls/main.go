// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// example inspired by
// https://github.com/OneLoneCoder/videos/blob/master/OneLoneCoder_Balls1.cpp
package main

import (
	"image/color"

	"github.com/roeldev/go-sdl2-experiments/internal"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/colors"
	"github.com/veandco/go-sdl2/sdl"
)

const UseOpengl = false

func main() {
	sdlkit.FailOnErr(sdl.Init(sdl.INIT_VIDEO))
	defer sdl.Quit()

	//goland:noinspection GoBoolExpressions
	if UseOpengl && sdl.SetHint(sdl.HINT_RENDER_DRIVER, "opengl") {
		sdlkit.DefaultOptions.WindowFlags += sdl.WINDOW_OPENGL
	}

	sdlkit.DefaultOptions.WindowFlags += sdl.WINDOW_RESIZABLE
	sdlkit.DefaultOptions.BgColor = color.RGBA(colors.Black)

	stage := sdlkit.MustNewStage(internal.ExampleName(), 1024, 768, sdlkit.DefaultOptions)
	defer stage.Destroy()

	stage.MustAddScene(newScene(stage))
	sdlkit.FailOnErr(sdlkit.RunLoop(stage))
}
