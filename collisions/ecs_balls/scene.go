// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"math"

	"github.com/go-pogo/errors"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/colors"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/display"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/ecs"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/event"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/input"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/physics"
)

const (
	PrerenderBalls = true

	BodyComponent ecs.ComponentTag = 1 << iota
	ColorComponent
)

var rng = sdlkit.RNG()

type scene struct {
	stage  *sdlkit.Stage
	ecs    *ecs.Manager
	canvas *sdlkit.Canvas
	events event.Manager
	mouse  *input.MouseState

	balls        []*ball
	collisions   [][2]*ball
	selectedBall *ball

	ballsLayer sdlkit.Layer
	debugLayer sdlkit.Layer
}

func newScene(stage *sdlkit.Stage) (sdlkit.Scene, error) {
	sc := &scene{
		stage: stage,
		ecs:   ecs.NewManager(),
		mouse: input.NewMouseState(input.TrackMouseBtnLeft | input.TrackMouseBtnRight),
	}

	sc.debugLayer.Append(display.NewFpsDisplay(stage.Time(), 10, 10))

	amount := sc.ballsAmount()
	sc.collisions = make([][2]*ball, 0, amount)
	sc.events.RegisterHandler(stage, sc.mouse, sc)
	sc.addBalls(amount)

	return sc, nil
}

func (sc *scene) SceneName() string { return "ecs_balls" }

func (sc *scene) ballsAmount() int {
	return int(sc.stage.Width() * sc.stage.Height() / 8000)
}

func (sc *scene) addBalls(amount int) {
	if amount < 1 {
		return
	}

	stageSize := sc.stage.Size()
	for i := 0; i < amount; i++ {
		b := sc.addBall()
		b.RandPosition(stageSize.W, stageSize.H)
		b.RandVelocity(1)
	}
}

func (sc *scene) addBall() *ball {
	rad := MinBallRadius + rng.Int31n(MaxBallRadius-MinBallRadius)
	col := colors.RandColor(rng)
	for col == sdl.Color(sc.stage.BgColor) {
		col = colors.RandColor(rng)
	}

	ball := newBall(float64(rad), col)
	sc.balls = append(sc.balls, ball)
	sc.ballsLayer.Append(ball)

	//goland:noinspection GoBoolExpressions
	if !PrerenderBalls {
		return ball
	}
	if sc.canvas == nil {
		sc.canvas = sdlkit.NewCanvas(sc.stage.Renderer())
		sc.canvas.AntiAlias(true)
		sc.canvas.BlendMode(sdl.BLENDMODE_BLEND)
	}

	_ = ball.Prerender(sc.canvas)
	return ball
}

func (sc *scene) HandleWindowSizeChangedEvent(_ *sdl.WindowEvent) error {
	stageSize := sc.stage.Size()
	for _, b := range sc.balls {
		b.RandPosition(stageSize.W, stageSize.H)
		b.RandVelocity(1)
	}

	sc.addBalls(sc.ballsAmount())
	return nil
}

func (sc *scene) HandleRenderEvent(_ *sdl.RenderEvent) error {
	var err error
	for _, b := range sc.balls {
		errors.Append(&err,
			b.tx.Destroy(),
			b.Prerender(sc.canvas),
		)
	}
	return err
}

func (sc *scene) Process() error {
	err := sc.events.Process()
	if err != nil {
		return err
	}

	if sc.selectedBall == nil {
		if sc.mouse.BtnLeft.Pressed && sc.mouse.BtnLeft.Clicks == 2 {
			ball := sc.addBall()
			ball.X = sc.mouse.BtnLeft.X
			ball.Y = sc.mouse.BtnLeft.Y
			ball.RandVelocity(10)
		} else if sc.mouse.BtnLeft.Pressed || sc.mouse.BtnRight.Pressed {
			for _, ball := range sc.balls {
				if !ball.HitTestXY(sc.mouse) {
					continue
				}

				sc.selectedBall = ball
				ball.Vel.Zero()
				ball.Acc.Zero()
				break
			}
		}
	} else {
		if sc.mouse.BtnLeft.Pressed {
			// ball follows mouse position
			sc.selectedBall.X = sc.mouse.X
			sc.selectedBall.Y = sc.mouse.Y
		} else if !sc.mouse.BtnRight.Pressed {
			if sc.mouse.BtnRight.Released {
				sc.selectedBall.Acc.Zero()
				sc.selectedBall.Vel.X = 10 * (sc.selectedBall.X - sc.mouse.X)
				sc.selectedBall.Vel.Y = 10 * (sc.selectedBall.Y - sc.mouse.Y)
			}
			sc.selectedBall = nil
		}
	}

	return nil
}

func (sc *scene) Update(dt float64) {
	var checkCollisions bool
	stageW, stageH := sc.stage.FWidth(), sc.stage.FHeight()

	// first update positions of balls
	for _, b := range sc.balls {
		if b.Vel.X != 0 || b.Vel.Y != 0 {
			b.Update(dt)
			b.ClampPosition(stageW, stageH)
			checkCollisions = true
		} else if sc.selectedBall != nil {
			b.ClampPosition(stageW, stageH)
			checkCollisions = true
		}
	}

	if !checkCollisions {
		return
	}

	// sc.ecs.Entities(BodyComponent)
	physics.DefaultSystem.ResolveStaticCollisions()
	physics.DefaultSystem.ResolveDynamicCollisions(nil)

	for _, col := range sc.collisions {
		x := col[0].X - col[1].X
		y := col[0].Y - col[1].Y

		distance := math.Sqrt((x * x) + (y * y))

		nx := (col[1].X - col[0].X) / distance
		ny := (col[1].Y - col[0].Y) / distance
		kx := col[0].Vel.X - col[1].Vel.X
		ky := col[0].Vel.Y - col[1].Vel.Y

		p := 2.0 * ((nx * kx) + (ny * ky)) / (col[0].Mass + col[1].Mass)

		col[0].Vel.X -= p * col[1].Mass * nx
		col[0].Vel.Y -= p * col[1].Mass * ny
		col[1].Vel.X += p * col[0].Mass * nx
		col[1].Vel.Y += p * col[0].Mass * ny
	}
}

func (sc *scene) Render(ren *sdl.Renderer) error {
	_ = sdlkit.Render(ren,
		sc.ballsLayer,
		sc.debugLayer,
	)

	_ = ren.SetDrawColor(colors.Red.R, colors.Red.G, colors.Red.B, 0xFF)
	for _, col := range sc.collisions {
		_ = ren.DrawLine(
			int32(col[0].X),
			int32(col[0].Y),
			int32(col[1].X),
			int32(col[1].Y),
		)
	}
	if sc.selectedBall != nil && sc.mouse.BtnRight.Pressed {
		_ = ren.SetDrawColor(
			colors.BlueViolet.R,
			colors.BlueViolet.G,
			colors.BlueViolet.B,
			0xFF,
		)
		_ = ren.DrawLine(
			int32(sc.selectedBall.X),
			int32(sc.selectedBall.Y),
			int32(sc.mouse.X),
			int32(sc.mouse.Y),
		)
	}

	return nil
}

func (sc *scene) Destroy() error {
	var err error
	for _, b := range sc.balls {
		errors.Append(&err, b.Destroy())
	}
	return err
}
