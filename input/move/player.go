package main

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/colors"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/display"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/ecs"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/geom"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/input"
)

type player struct {
	geom.Circle
	Rotation float64
	pos      geom.Point
	mover    gameplay.Mover
	shape    display.DevCircle
}

func newPlayer() *player {
	p := &player{
		shape: display.DevCircle{
			Circle: display.Circle{
				X:         0,
				Y:         0,
				Radius:    10,
				LineColor: sdl.Color{},
				FillColor: sdl.Color(colors.Black),
			},
			Rotation: 0,
		},
	}
	p.mover = gameplay.NewKeyboardMover(p)
	return p
}

func (p *player) Position() *geom.Point { return &p.pos }

func (p *player) AddComponent(v ecs.Component) {}

func (p *player) Component(v interface{}) bool { return false }

func (p *player) Update(clock *sdlkit.Clock) {
	p.mover.Update(clock)
	p.shape.X = int32(p.pos.X)
	p.shape.Y = int32(p.pos.Y)
	p.shape.Rotation = p.mover.Velocity().Angle() // p.mover.Direction() ?
}

func (p *player) Render(r *sdl.Renderer) error {
	return p.shape.Draw(r)
}
