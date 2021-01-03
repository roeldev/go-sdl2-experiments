// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"

	"github.com/roeldev/go-x11colors"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/go-pogo/sdlkit"
)

func printErr(err error) {
	_, _ = fmt.Fprint(os.Stderr, err)
	os.Exit(1)
}

func main() {
	sdlkit.FailOnErr(sdl.Init(sdl.INIT_VIDEO))
	defer sdl.Quit()

	sdlkit.DefaultStageOpts.BgColor = x11colors.LightGrey

	stage, err := sdlkit.NewStage("example 2", 400, 300, sdlkit.DefaultStageOpts)
	if err != nil {
		sdlkit.FailOnErr(err)
	}

	world := newWorld(stage)
	defer world.Destroy()
	sdlkit.FailOnErr(world.Run())
}
