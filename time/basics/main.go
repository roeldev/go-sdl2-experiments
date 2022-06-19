// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"image/color"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit"

	"github.com/roeldev/go-sdl2-experiments/internal"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/colors"
)

func main() {
	sdlkit.FailOnErr(sdl.Init(sdl.INIT_VIDEO))
	defer sdl.Quit()

	sdlkit.DefaultOptions.BgColor = color.RGBA(colors.LightGrey)

	stage := sdlkit.MustNewStage(internal.ExampleName(), 400, 300, sdlkit.DefaultOptions)
	defer stage.Destroy()

	sdlkit.FailOnErr(newWorld(stage).Run())
}
