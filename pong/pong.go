// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/display"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/event"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/geom"
)

type pongGame struct {
	stage  *sdlkit.Stage
	events event.Manager

	bgLayer      sdlkit.Layer
	paddlesLayer sdlkit.Layer
	guiLayer     sdlkit.Layer
	debugLayer   sdlkit.Layer

	// entities
	ball *ball
	// paddles
	paddleLeft  *paddle
	paddleRight *paddle
}

func newGame(stage *sdlkit.Stage) *pongGame {
	game := &pongGame{
		stage:      stage,
		ball:       newBall(stage, DefaultBallRadius),
		paddleLeft: newPaddle(paddleLeft, 0, 0), // computer player
		// paddleLeft:  newPaddle(paddleLeft, sdl.SCANCODE_Q, sdl.SCANCODE_A),
		paddleRight: newPaddle(paddleRight, sdl.SCANCODE_UP, sdl.SCANCODE_DOWN),
	}

	stageW, stageH := stage.FWidth(), stage.FHeight()

	// add entity event handlers
	game.events.RegisterHandler(stage, game, game.paddleLeft, game.paddleRight)
	game.paddlesLayer.Append(game.paddleLeft, game.paddleRight)
	game.paddleLeft.UpdateClampArea(stageW, stageH)
	game.paddleRight.UpdateClampArea(stageW, stageH)

	// additional (non default) entities for rendering on screen
	game.debugLayer.Append(display.NewFpsDisplay(stage.Time(), 10, 10))

	return game
}

func (game *pongGame) SceneName() string { return "pong" }

func (game *pongGame) HandleWindowSizeChangedEvent(_ *sdl.WindowEvent) error {
	w, h := game.stage.FWidth(), game.stage.FHeight()
	game.paddleLeft.UpdateClampArea(w, h)
	game.paddleRight.UpdateClampArea(w, h)
	return nil
}

func (game *pongGame) HandleKeyboardEvent(event *sdl.KeyboardEvent) error {
	switch event.Keysym.Scancode {
	case sdl.SCANCODE_ESCAPE:
		game.Reset()
		return nil
	}

	return nil
}

func (game *pongGame) Activate() error {
	game.Reset()
	return nil
}

func (game *pongGame) Reset() {
	game.stage.Clock().TimeScale = sdlkit.DefaultTimeScale
	game.ball.Reset()
}

func (game *pongGame) Process() error {
	return game.events.Process()
}

func (game *pongGame) Update(dt float64) {
	c := game.stage.Clock()
	c.TimeScale += dt / 10

	if !game.ball.Update(c) {
		game.Reset()
	}

	game.paddleLeft.Update(c)
	game.paddleRight.Update(c)

	if game.paddleLeft.IsComputer() {
		game.paddleLeft.Y = game.ball.Y
	}
	if game.paddleRight.IsComputer() {
		game.paddleRight.Y = game.ball.Y
	}

	bl := game.ball.X - game.ball.Radius
	br := game.ball.X + game.ball.Radius

	bounds := game.paddleLeft.Bounds()
	if geom.InRect(bl, game.ball.Y, bounds.X, bounds.Y, bounds.W, bounds.H) {
		game.ball.X += bounds.X + bounds.W - bl
		game.ball.Vel.X *= -1
	} else {
		bounds = game.paddleRight.Bounds()
		if geom.InRect(br, game.ball.Y, bounds.X, bounds.Y, bounds.W, bounds.H) {
			game.ball.X += bounds.X - br
			game.ball.Vel.X *= -1
		}
	}
}

func (game *pongGame) Render(r *sdl.Renderer) error {
	return sdlkit.Render(r,
		game.bgLayer,
		game.paddlesLayer,
		game.ball,
		game.guiLayer,
		game.debugLayer,
	)
}
