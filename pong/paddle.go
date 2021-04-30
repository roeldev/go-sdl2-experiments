package main

import (
	"github.com/go-pogo/errors"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/colors"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/geom"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/math"
)

const (
	paddleLeft paddlePos = 1 << iota
	paddleRight

	PaddleSpeed = 300
)

type paddlePos uint8

const (
	paddleWidth  = 10
	paddleHeight = 60
)

type paddle struct {
	geom.Rect
	Vel geom.Vector

	pos        paddlePos
	cmin, cmax float64
	color      sdl.Color

	keyUp, keyDown sdl.Scancode

	MoveUp, MoveDown bool
}

func newPaddle(pos paddlePos, u, d sdl.Scancode) *paddle {
	return &paddle{
		Rect: geom.Rect{W: paddleWidth, H: paddleHeight},

		pos:     pos,
		cmin:    paddleHeight / 2,
		keyUp:   u,
		keyDown: d,
		color:   colors.RandWhite(rng),
	}
}

func (p *paddle) IsComputer() bool { return p.keyUp == 0 || p.keyDown == 0 }

func (p *paddle) HandleKeyboardEvent(event *sdl.KeyboardEvent) error {
	switch event.Keysym.Scancode {
	case p.keyUp:
		p.MoveUp = event.State == sdl.PRESSED

	case p.keyDown:
		p.MoveDown = event.State == sdl.PRESSED
	}
	return nil
}

func (p *paddle) Update(clock *sdlkit.Clock) {
	if p.MoveUp && !p.MoveDown {
		p.Vel.Y = -PaddleSpeed
	} else if p.MoveDown {
		p.Vel.Y = PaddleSpeed
	} else {
		p.Vel.Y = 0
	}

	p.Y += p.Vel.Y * clock.Delta64
	p.Y = math.Clamp(p.Y, p.cmin, p.cmax)
}

func (p *paddle) UpdateClampArea(w, h float64) {
	p.cmax = h - paddleHeight/2

	switch p.pos {
	case paddleLeft:
		p.X = 10 + (paddleWidth / 2)
	case paddleRight:
		p.X = w - 10 - (paddleWidth / 2)
	}
}

func (p *paddle) Render(r *sdl.Renderer) (err error) {
	errors.Append(&err,
		r.SetDrawColor(p.color.R, p.color.G, p.color.B, p.color.A),
		r.DrawRect(p.Bounds().Rect()),
	)
	return err
}
