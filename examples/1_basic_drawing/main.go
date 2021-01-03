// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "C"
import (
	"os"
	"path"
	"runtime"

	"github.com/go-pogo/errors"
	"github.com/roeldev/go-x11colors"
	sdlimg "github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/go-pogo/sdlkit"
)

func resolveRootDir(skipCallers int) (string, error) {
	_, dir, _, ok := runtime.Caller(skipCallers)
	if !ok {
		return "", errors.New("unable to resolve project dir location")
	}

	return path.Dir(dir), nil
}

func main() {
	sdlkit.FailOnErr(sdl.Init(sdl.INIT_VIDEO))
	defer sdl.Quit()

	sdlkit.FailOnErr(sdlimg.Init(sdlimg.INIT_PNG))
	defer sdlimg.Quit()

	dir, err := resolveRootDir(0)
	if err != nil {
		sdlkit.FailOnErr(err)
	}

	sdlkit.DefaultStageOpts.BgColor = x11colors.White
	sdlkit.DefaultStageOpts.Icon, _ = sdlimg.Load(dir + "/img.ico")

	stage, err := sdlkit.NewStage("example 1", 400, 300, sdlkit.DefaultStageOpts)
	if err != nil {
		sdlkit.FailOnErr(err)
	}

	defer stage.Destroy()

	tx, err := sdlimg.LoadTexture(stage.Renderer(), dir+"/img.png")
	if err != nil {
		sdlkit.FailOnErr(err)
	}

	// render to window
	rt := stage.RenderTarget()
	rt.ClearAndDraw(
		drawSquare(100, 10, 190, x11colors.RandRed(sdlkit.RNG())),
		drawCircle(30, 130, 80, x11colors.RandBlue(sdlkit.RNG())),
		drawImg(tx, 5, 100, 50, sdl.FLIP_NONE),
		drawImg(tx, 1, 10, 10, sdl.FLIP_HORIZONTAL|sdl.FLIP_VERTICAL),
	)

	sdlkit.FailOnErr(rt.Err())
	stage.PresentScreen()

	for {
		switch sdl.PollEvent().(type) {
		case *sdl.QuitEvent:
			os.Exit(0)
		}
	}
}
