// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embed"
	"image/color"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit-examples/internal"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/colors"
)

//go:embed "assets"
var assets embed.FS

func main() {
	sdlkit.FailOnErr(sdl.Init(sdl.INIT_VIDEO))
	defer sdl.Quit()

	sdlkit.DefaultOptions.WindowFlags += sdl.WINDOW_RESIZABLE
	sdlkit.DefaultOptions.BgColor = color.RGBA(colors.Blue)

	stage := sdlkit.MustNewStage(internal.ExampleName(), 1024, 768, sdlkit.DefaultOptions)
	defer stage.Destroy()

	stage.MustAddScene(newScene(stage, assets))
	sdlkit.FailOnErr(sdlkit.RunLoop(stage))
}
