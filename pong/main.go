// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit-examples/internal"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/colors"
)

const Debug = false

var rng = sdlkit.RNG()

func main() {
	sdlkit.FailOnErr(sdl.Init(sdl.INIT_VIDEO))
	defer sdl.Quit()

	sdlkit.DefaultOptions.WindowFlags += sdl.WINDOW_RESIZABLE
	sdlkit.DefaultOptions.BgColor = colors.DarkSlateGray

	stage := sdlkit.MustNewStage(internal.ExampleName(), 1024, 576, sdlkit.DefaultOptions)
	defer stage.Destroy()

	sdlkit.FailOnErr(stage.AddScene(newGame(stage)))
	sdlkit.FailOnErr(sdlkit.RunLoop(stage))
}
