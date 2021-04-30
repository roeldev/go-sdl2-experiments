// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/colors"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/display"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/ecs"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/event"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/geom"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/input"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/physics"
)

const (
	BodyComponent ecs.ComponentTag = 1 << iota
	ColliderComponent
	TransformComponent
	ColorComponent
)

type scene struct {
	stage  *sdlkit.Stage
	events event.Manager
	ecs    *ecs.Manager
	canvas *sdlkit.Canvas
	mouse  *input.MouseState
	keys   input.KeyboardState
	fps    *display.FpsDisplay

	polygonEntity ecs.Entity
}

func newScene(stage *sdlkit.Stage) (sdlkit.Scene, error) {
	sc := &scene{
		stage:  stage,
		ecs:    ecs.NewManager(),
		canvas: sdlkit.NewCanvas(stage.Renderer()),
		mouse:  input.NewMouseState(0),
		fps:    display.NewFpsDisplay(stage.Time(), 10, 10),
	}

	sc.events.MustRegisterHandler(stage, sc.mouse)

	return sc, nil
}

func (demo *scene) SceneName() string { return "" }

func (demo *scene) createShapeEntity(body physics.Body) ecs.Entity {
	entity := demo.ecs.Create(nil)
	entity.AddComponent(ColorComponent, colors.RandColor(sdlkit.RNG()))
	entity.AddComponent(BodyComponent, body)
	entity.AddComponent(ColliderComponent, body)

	return entity
}

func (demo *scene) Activate() error {
	width, height := demo.stage.FWidth(), demo.stage.FHeight()
	centerX, centerY := width/2, height/2

	demo.createShapeEntity(physics.NewStaticBody(
		&geom.Ellipse{X: width - 200, Y: 265, RadiusX: 100, RadiusY: 60},
		nil,
	))
	demo.createShapeEntity(physics.NewStaticBody(
		&geom.Circle{X: 100, Y: 100, Radius: 40},
		nil,
	))
	demo.createShapeEntity(physics.NewStaticBody(
		&geom.Rect{X: 140, Y: height - 120, W: 180, H: 180},
		nil,
	))

	// demo.ecs.Create(nil).
	// 	AddComponent(ColorComponent, colors.RandColor(sdlkit.RNG())).
	// 	AddComponent(PhysicsBodyComponent, physics.NewDynamicBody(
	// 		geom.NewRegularPolygon(800, 650, 40, 3),
	// 	))

	demo.ecs.Register(newTriangle())

	demo.polygonEntity = demo.createShapeEntity(physics.NewDynamicBody(
		// 3+uint8(sdlkit.RNG().Int31n(10))
		geom.NewRegularPolygon(centerX, centerY, 100, 5),
		nil,
	))

	return nil
}

func (demo *scene) Process() error { return demo.events.Process() }

func (demo *scene) Update(dt float64) {
	poly := demo.polygonEntity.Component(BodyComponent).(*physics.DynamicBody)
	poly.Transformer().Rotate(dt)
	poly.Update()
}

func (demo *scene) Render(_ *sdl.Renderer) error {
	for _, entity := range demo.ecs.Entities(ColliderComponent, ColorComponent) {
		body := entity.Component(BodyComponent).(physics.Body)
		demo.canvas.BeginLineStyle(2, colors.Blue)
		demo.canvas.DrawSdlRect(body.Bounds().Rect())
		demo.canvas.EndLineStyle()

		demo.canvas.BeginFill(colors.Yellow)
		demo.canvas.DrawPixelXY(body.Shape(), 3)
		demo.canvas.EndFill()

		if body.HitTest(demo.mouse.X, demo.mouse.Y) {
			demo.canvas.BeginLineStyle(1, colors.Green)
		} else {
			color := entity.Component(ColorComponent).(sdl.Color)
			demo.canvas.BeginLineStyle(1, color)
		}

		demo.canvas.DrawShape(body.Shape())
	}

	demo.canvas.BeginFill(colors.Red)
	demo.canvas.DrawPixelXY(demo.mouse, 2)
	demo.canvas.EndFill()

	demo.canvas.Draw(demo.fps)
	return demo.canvas.Done()
}
