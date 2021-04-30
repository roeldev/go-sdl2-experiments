// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	sdlgfx "github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/colors"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/geom"
)

const DefaultBallRadius = 10

type ball struct {
	geom.Circle

	Vel geom.Vector // velocity (speed and direction)
	Acc geom.Vector // acceleration

	stage *sdlkit.Stage
	color sdl.Color
	debug bool
}

func newBall(stage *sdlkit.Stage, rad float64) *ball {
	return &ball{
		Circle: geom.Circle{Radius: rad},

		stage: stage,
		color: sdl.Color(colors.White),
		debug: Debug,
	}
}

func (b *ball) Reset() {
	b.X, b.Y = b.stage.FWidth()/2, b.stage.FHeight()/2
	b.Vel = geom.Vector{X: 250, Y: 250}
	b.Acc = geom.Vector{X: 0.01, Y: 0.01}

	if rng.Float32() < 0.5 {
		b.Vel.X *= -1
		if rng.Float32() >= 0.5 {
			b.Vel.Y *= -1
		}
	}
}

func (b *ball) Update(clock *sdlkit.Clock) bool {
	b.Acc.Scale(1 + (clock.Delta64 / 5))

	// increase velocity by acceleration
	if b.Vel.X < 0 {
		b.Vel.X -= b.Acc.X * clock.Delta64
	} else {
		b.Vel.X += b.Acc.X * clock.Delta64
	}
	if b.Vel.Y < 0 {
		b.Vel.Y -= b.Acc.Y * clock.Delta64
	} else {
		b.Vel.Y += b.Acc.Y * clock.Delta64
	}

	// add velocity to position
	b.X += b.Vel.X * clock.Delta64
	b.Y += b.Vel.Y * clock.Delta64

	// left and right of screen
	if b.X < -b.Radius || b.X > b.stage.FWidth()+b.Radius {
		return false
	}

	edgeT := b.Radius
	edgeB := b.stage.FHeight() - b.Radius
	if b.Y < edgeT {
		b.Y += edgeT - b.Y
		b.Vel.Y *= -1
	} else if b.Y > edgeB {
		b.Y += edgeB - b.Y
		b.Vel.Y *= -1
	}
	return true
}

func (b *ball) Render(r *sdl.Renderer) error {
	sdlgfx.AACircleColor(r,
		int32(b.X),
		int32(b.Y),
		int32(b.Radius),
		b.color,
	)
	return nil
}
