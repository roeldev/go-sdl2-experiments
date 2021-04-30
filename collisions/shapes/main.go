// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"image/color"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit-examples/internal"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/colors"
)

func main() {
	sdlkit.FailOnErr(sdl.Init(sdl.INIT_VIDEO))
	defer sdl.Quit()

	sdlkit.DefaultOptions.BgColor = color.RGBA(colors.Black)

	stage := sdlkit.MustNewStage(internal.ExampleName(), 1024, 768, sdlkit.DefaultOptions)
	defer stage.Destroy()

	stage.MustAddScene(newScene(stage))
	sdlkit.FailOnErr(sdlkit.RunLoop(stage))
}
