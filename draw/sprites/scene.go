// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"io/fs"
	"log"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/display"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/event"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/geom/align"
)

type scene struct {
	event.Manager

	stage  *sdlkit.Stage
	assets fs.ReadFileFS
	atlas  *sdlkit.TextureAtlas

	bgLayer sdlkit.Layer
	ground  *display.Sprite
	house   *display.Sprite

	needsUpdate bool
}

func newScene(stage *sdlkit.Stage, assets fs.ReadFileFS) (sdlkit.Scene, error) {
	load := sdlkit.NewAssetsLoader(assets, stage.Renderer())
	atlas, err := load.TextureAtlasXml("assets/spritesheet_default.xml")
	if err != nil {
		return nil, err
	}

	demo := &scene{
		stage:       stage,
		assets:      assets,
		atlas:       atlas,
		needsUpdate: true,
		ground:      display.MustNewSprite(load.TextureClip("assets/groundLayer2.png")),
	}

	demo.RegisterHandler(stage, demo)
	demo.bgLayer.Append(demo.ground)
	demo.ground.StretchMode = display.StretchTile

	return demo, nil
}

func (demo *scene) SceneName() string { return "" }

func (demo *scene) Activate() error {
	tc, err := demo.atlas.GetFromName("houseAlt1.png")
	if err != nil {
		return err
	}

	demo.house = display.NewSprite(tc)
	// demo.bgLayer.Append(demo.house)
	return nil
}

func (demo *scene) HandleWindowSizeChangedEvent(_ *sdl.WindowEvent) error {
	demo.needsUpdate = true
	return nil
}

func (demo *scene) Update(_ float64) {
	if !demo.needsUpdate {
		return
	}

	demo.needsUpdate = false
	screenSize := demo.stage.Size()
	align.XYInSdlRect(align.ToCenter, demo.house, screenSize)

	demo.ground.W = demo.stage.FWidth()
	demo.ground.X = demo.ground.W / 2
	demo.ground.Y = demo.ground.H / 2
	log.Println(demo.stage.Renderer().GetScale())
}

func (demo *scene) Render(r *sdl.Renderer) error {
	return sdlkit.Render(r, demo.bgLayer)
}

func (demo *scene) Destroy() error {
	return demo.atlas.Destroy()
}
