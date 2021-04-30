// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	sdlgfx "github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/geom"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/physics"
)

const (
	MinBallRadius int32 = 10
	MaxBallRadius int32 = 30
)

type ball struct {
	*physics.DynamicBody
	geom.Circle

	Vel   geom.Vector // velocity (speed and direction)
	Acc   geom.Vector // acceleration
	Mass  float64
	Color sdl.Color

	tx  *sdl.Texture
	dst sdl.Rect
}

func newBall(rad float64, col sdl.Color) *ball {
	circle := geom.Circle{Radius: rad}
	return &ball{
		DynamicBody: physics.NewDynamicBody(circle, nil),
		Circle:      circle,
		Mass:        rad * 10,
		Color:       col,
	}
}

func (b *ball) RandPosition(w, h int32) {
	b.X = float64(MaxBallRadius + rng.Int31n(w-MaxBallRadius-MaxBallRadius))
	b.Y = float64(MaxBallRadius + rng.Int31n(h-MaxBallRadius-MaxBallRadius))
}

func (b *ball) RandVelocity(f float64) {
	b.Vel.X = float64(rng.Intn(300)-150) * f
	b.Vel.Y = float64(rng.Intn(300)-150) * f
}

func (b *ball) Prerender(canvas *sdlkit.Canvas) error {
	rad := int32(b.Radius * 4)

	tx, err := canvas.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_TARGET, rad*2, rad*2)
	if err != nil {
		return err
	}

	canvas.AntiAlias(true)
	canvas.BeginFill(b.Color)
	canvas.DrawCircle(rad, rad, rad)
	if err = canvas.Done(); err != nil {
		return err
	}

	b.tx = tx
	b.dst = sdl.Rect{W: rad/2 + 1, H: rad/2 + 1}
	return nil
}

func (b *ball) Update(dt float64) {
	// add friction to acceleration
	b.Acc.X = b.Vel.X * -0.5
	b.Acc.Y = b.Vel.Y * -0.5
	// add acceleration to velocity
	b.Vel.X += b.Acc.X * dt
	b.Vel.Y += b.Acc.Y * dt
	// add velocity to position
	b.X += b.Vel.X * dt
	b.Y += b.Vel.Y * dt

	if ((b.Vel.X * b.Vel.X) + (b.Vel.Y * b.Vel.Y)) < 1 {
		b.Acc.Zero()
		b.Vel.Zero()
	}
}

func (b *ball) ClampPosition(right, bottom float64) {
	top, left := b.Radius, b.Radius

	// subtract 1 from right and bottom border to compensate for the ball's
	// outline thickness
	right -= b.Radius - 1
	bottom -= b.Radius - 1

	// bounce of off screen edges if needed
	if b.X <= left {
		b.X += left - b.X
		b.Vel.X *= -1
	} else if b.X > right {
		b.X -= b.X - right
		b.Vel.X *= -1
	}
	if b.Y <= top {
		b.Y += top - b.Y
		b.Vel.Y *= -1
	} else if b.Y > bottom {
		b.Y -= b.Y - bottom
		b.Vel.Y *= -1
	}
}

func (b *ball) Render(ren *sdl.Renderer) error {
	if b.tx == nil {
		sdlgfx.CircleColor(ren, int32(b.X), int32(b.Y), int32(b.Radius), b.Color)
		return nil
	}

	b.dst.X = int32(b.X - b.Radius)
	b.dst.Y = int32(b.Y - b.Radius)
	return ren.Copy(b.tx, nil, &b.dst)
}

func (b *ball) Destroy() error {
	if b.tx != nil {
		return b.tx.Destroy()
	}
	return nil
}
