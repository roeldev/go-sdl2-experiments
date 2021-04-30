// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/event"
)

type scene struct {
	stage  *sdlkit.Stage
	events event.Manager

	layer sdlkit.Layer

	// entities
	player *player
}

func newScene(stage *sdlkit.Stage) (sdlkit.Scene, error) {
	s := &scene{
		stage:  stage,
		player: newPlayer(),
	}

	s.player.pos.X, s.player.pos.Y = stage.FWidth()/2, stage.FHeight()/2
	s.events.RegisterHandler(stage, s, s.player.mover)
	s.layer.Append(s.player)
	return s, nil
}

func (sc *scene) SceneName() string { return "" }

func (s *scene) HandleKeyUpEvent(e *sdl.KeyboardEvent) error {
	if e.Keysym.Scancode == sdl.SCANCODE_ESCAPE {
		s.player.pos.X, s.player.pos.Y = s.stage.FWidth()/2, s.stage.FHeight()/2
	}

	return nil
}

func (s *scene) Process() error {
	return s.events.Process()
}

func (s *scene) Update(_ float64) {
	clock := s.stage.Clock()
	s.player.Update(clock)
}

func (s *scene) Render(r *sdl.Renderer) error {
	return sdlkit.Render(r, s.player)
}
