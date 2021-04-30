// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	_ "embed"
	"image/color"
	"os"
	"path"
	"runtime"

	"github.com/go-pogo/errors"
	sdlimg "github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/colors"
)

//go:embed "img.ico"
var icon []byte

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

	stage, err := sdlkit.NewStage("sdlkit example", 400, 300, sdlkit.DefaultOptions)
	if err != nil {
		sdlkit.FailOnErr(err)
	}

	stage.BgColor = color.RGBA(colors.White)
	_ = stage.SetWindowIcon(icon)
	defer stage.Destroy()

	tx, err := sdlimg.LoadTexture(stage.Renderer(), dir+"/img.png")
	if err != nil {
		sdlkit.FailOnErr(err)
	}

	// render to window
	sdlkit.FailOnErr(stage.ClearScreen())
	sdlkit.FailOnErr(sdlkit.Render(stage.Renderer(),
		drawSquare(100, 10, 190, colors.RandRed(sdlkit.RNG())),
		drawCircle(30, 130, 80, colors.RandBlue(sdlkit.RNG())),
		drawImg(tx, 5, 100, 50, sdl.FLIP_NONE),
		drawImg(tx, 1, 10, 10, sdl.FLIP_HORIZONTAL|sdl.FLIP_VERTICAL),
	))

	stage.PresentScreen()

	for {
		switch sdl.PollEvent().(type) {
		case *sdl.QuitEvent:
			os.Exit(0)
		}
	}
}
