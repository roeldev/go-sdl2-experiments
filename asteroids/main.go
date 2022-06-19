// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	_ "embed"
	"image/color"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/roeldev/go-sdl2-experiments/internal"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/colors"
)

//go:embed "splashscreen.jpg"
var splashscreen []byte

func main() {
	sdlkit.FailOnErr(sdl.Init(sdl.INIT_VIDEO))
	defer sdl.Quit()

	splash := sdlkit.MustNewSplashScreen(320, 240)
	//goland:noinspection GoUnhandledErrorResult
	defer splash.Destroy()
	_ = splash.DisplayImage(splashscreen)

	if sdl.SetHint(sdl.HINT_RENDER_DRIVER, "opengl") {
		sdlkit.DefaultOptions.WindowFlags += sdl.WINDOW_OPENGL
	}

	// sdlkit.DefaultOptions.WindowFlags += sdl.WINDOW_MAXIMIZED
	sdlkit.DefaultOptions.BgColor = color.RGBA(colors.Black)

	// hide mouse
	_, _ = sdl.ShowCursor(sdl.DISABLE)

	// 853
	stage := sdlkit.MustNewStage(internal.ExampleName(), 800, 480, sdlkit.DefaultOptions)
	defer stage.Destroy()
	_ = splash.Destroy()

	stage.MustAddScene(newAsteroids(stage))
	sdlkit.FailOnErr(sdlkit.RunLoop(stage))
}
