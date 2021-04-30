// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/display"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/event"
)

type asteroidsGame struct {
	stage  *sdlkit.Stage
	events event.Manager

	debugLayer sdlkit.Layer

	starField *starField
	player    *player
}

func newAsteroids(stage *sdlkit.Stage) (sdlkit.Scene, error) {
	game := &asteroidsGame{
		stage:     stage,
		starField: newStarField(40),
		player:    newPlayer(stage.FWidth()/2, stage.FHeight()/2),
	}

	game.events.RegisterHandler(stage)
	game.debugLayer.Append(display.NewFpsDisplay(stage.Time(), 10, 10))

	return game, nil
}

func (game *asteroidsGame) SceneName() string { return "asteroids" }

func (game *asteroidsGame) Activate() error {
	return game.starField.Render(game.stage.Renderer(), game.stage.Size())
}

func (game *asteroidsGame) Process() error {
	return game.events.Process()
}

func (game *asteroidsGame) Update(_ float64) {}

func (game *asteroidsGame) Render(r *sdl.Renderer) error {
	return sdlkit.Render(r,
		game.starField,
		game.player,
		game.debugLayer,
	)
}
