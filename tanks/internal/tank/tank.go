// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tank

import (
	"math"

	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/colors"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/display/draw"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/event"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/geom"
	math2 "github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/math"
)

const (
	accelerationForwards  = 50
	accelerationBackwards = 30
	maxSpeedForwards      = 100
	maxSpeedBackwards     = 60
	steerAngle            = 0.6
	turretRotationSpeed   = 2
)

type Control interface {
	Input() UserInput
}

type UserInput struct {
	Forwards,
	Backwards,
	SteerLeft,
	SteerRight,
	Brake float64
}

// http://engineeringdotnet.blogspot.com/2010/04/simple-2d-car-physics-in-games.html
type Tank struct {
	Turret
	*body
	Control

	halfBase  float64    // half of the "wheelbase" (front to back of tank body)
	frontAxis geom.Point // absolute position of tank's front axis
	backAxis  geom.Point // absolute position of tank's back axis

	// friction float64
	speed float64 // current total speed

	// steerAngle float64    // amount of radians the tank is steering to the left or right
	heading [3]float64 // angle of rotation
}

func NewPlayer(atlas *sdlkit.TextureAtlas, color Color, control Control) *Tank {
	tnk := newTank(atlas, color)
	tnk.Control = control
	return tnk
}

func newTank(atlas *sdlkit.TextureAtlas, color Color) *Tank {
	body := newBody(atlas, color)
	tnk := &Tank{
		body:     body,
		halfBase: (body.height / 2) * 0.7,
	}

	tnk.Turret = NewTurretSmall(atlas, tnk)
	tnk.SetHeading(0)
	return tnk
}

func (tnk *Tank) RegisterEvents(em *event.Manager) {
	em.RegisterHandler(tnk.Control)
}

func (tnk *Tank) Color() Color     { return tnk.body.color }
func (tnk *Tank) Heading() float64 { return tnk.heading[0] }

func (tnk *Tank) GetX() float64 { return tnk.body.shape.X }
func (tnk *Tank) GetY() float64 { return tnk.body.shape.Y }

func (tnk *Tank) SetX(x float64) {
	tnk.body.shape.X = x
	tnk.frontAxis.X = x + (tnk.halfBase * tnk.heading[1])
	tnk.backAxis.X = x - (tnk.halfBase * tnk.heading[1])
}

func (tnk *Tank) SetY(y float64) {
	tnk.body.shape.Y = y
	tnk.frontAxis.Y = y + (tnk.halfBase * tnk.heading[2])
	tnk.backAxis.Y = y - (tnk.halfBase * tnk.heading[2])
}

func (tnk *Tank) SetHeading(heading float64) {
	rotation := heading - tnk.heading[0]
	if tnk.Turret.destRotation == tnk.Turret.curRotation {
		tnk.Turret.destRotation += rotation
	}
	tnk.Turret.curRotation += rotation

	tnk.heading[0] = heading
	tnk.heading[1] = math.Cos(heading)
	tnk.heading[2] = math.Sin(heading)
}

func (tnk *Tank) SetTurretRotation(radians float64) {
	tnk.Turret.destRotation = radians
}

// TurretPosition is the absolute center point of the turret's dome on top of
// the Tank.
func (tnk *Tank) TurretPosition() *geom.Vector {
	return tnk.Turret.offset.Clone().
		Rotate(-tnk.heading[0]).
		Add(tnk.GetX(), tnk.GetY())
}

func (tnk *Tank) Update(dt float64) {
	input := tnk.Input()
	if input.Brake != 0 ||
		(input.Backwards != 0 && tnk.speed > 0) ||
		(input.Forwards != 0 && tnk.speed < 0) {
		tnk.speed = math2.LerpRound(tnk.speed, 0, dt*4, 0.1)
	} else if input.Backwards == 0 && input.Forwards == 0 {
		tnk.speed = math2.LerpRound(tnk.speed, 0, dt*2, 0.1)
	} else {
		var accelerate float64
		if input.Forwards != 0 {
			accelerate = accelerationForwards * input.Forwards * dt
		} else if input.Backwards != 0 {
			accelerate = -accelerationBackwards * input.Backwards * dt
		}

		tnk.speed = math2.Clamp(tnk.speed+accelerate, -maxSpeedBackwards, maxSpeedForwards)
	}

	steer := (input.SteerRight - input.SteerLeft) * steerAngle

	tnk.frontAxis.X = (tnk.body.shape.X + (tnk.halfBase * tnk.heading[1])) + (tnk.speed * dt * math.Cos(tnk.heading[0]+steer))
	tnk.frontAxis.Y = (tnk.body.shape.Y + (tnk.halfBase * tnk.heading[2])) + (tnk.speed * dt * math.Sin(tnk.heading[0]+steer))
	tnk.backAxis.X = (tnk.body.shape.X - (tnk.halfBase * tnk.heading[1])) + (tnk.speed * dt * tnk.heading[1])
	tnk.backAxis.Y = (tnk.body.shape.Y - (tnk.halfBase * tnk.heading[2])) + (tnk.speed * dt * tnk.heading[2])

	// update tank center point which is in the middle of the front and back axis
	tnk.body.shape.X = (tnk.frontAxis.X + tnk.backAxis.X) / 2
	tnk.body.shape.Y = (tnk.frontAxis.Y + tnk.backAxis.Y) / 2

	if steer != 0 {
		tnk.SetHeading(math.Atan2(tnk.frontAxis.Y-tnk.backAxis.Y, tnk.frontAxis.X-tnk.backAxis.X))
		tnk.body.transform.SetRotation(tnk.heading[0])
		tnk.body.shape.Transform(tnk.body.transform.Matrix())
	}

	if tnk.curRotation != tnk.destRotation {
		// makes sure rotation is always using the least amount of radians
		if tnk.curRotation-tnk.destRotation > math.Pi {
			tnk.curRotation -= math2.PiDouble
		} else if tnk.destRotation-tnk.curRotation > math.Pi {
			tnk.curRotation += math2.PiDouble
		}

		rotate := tnk.destRotation - tnk.curRotation
		if tnk.Turret.curRotation > tnk.Turret.destRotation {
			if rotate > -0.01 {
				tnk.curRotation = tnk.destRotation
			} else {
				tnk.Turret.curRotation -= turretRotationSpeed * dt
			}
		} else if tnk.curRotation < tnk.Turret.destRotation {
			if rotate < 0.01 {
				tnk.curRotation = tnk.destRotation
			} else {
				tnk.Turret.curRotation += turretRotationSpeed * dt
			}
		}
	}
}

func (tnk *Tank) Draw(canvas *sdlkit.Canvas) {
	domeCenter := tnk.TurretPosition()
	tnk.Turret.update(domeCenter.X, domeCenter.Y)

	tnk.body.sprite.X = tnk.body.shape.X
	tnk.body.sprite.Y = tnk.body.shape.Y
	tnk.body.sprite.Rotation = tnk.heading[0] + math2.PiHalf

	// add 90 degrees because to "0" rotation point of the turret's barrel
	// points down instead of right
	tnk.Turret.barrelSprite.Rotation = tnk.Turret.curRotation - math2.PiHalf

	tnk.body.sprite.Draw(canvas)
	tnk.Turret.barrelSprite.Draw(canvas)

	// if tnk.Turret.domeSprite != nil {
	// 	tnk.Turret.domeSprite.Draw(canvas)
	// } else {
	canvas.BeginFill(turretColors[tnk.body.color][0])
	canvas.DrawCircle(int32(domeCenter.X-0.5), int32(domeCenter.Y-0.5), tnk.Turret.domeRadius[0])
	canvas.BeginFill(turretColors[tnk.body.color][1])
	canvas.DrawCircle(int32(domeCenter.X-0.5), int32(domeCenter.Y-0.5), tnk.Turret.domeRadius[1])
	canvas.EndFill()
	// }

	canvas.EndFill()
	canvas.BeginLineStyle(1, colors.Green)
	canvas.Draw(draw.Polygon(tnk.shape))
	canvas.BeginLineStyle(1, colors.Teal)
	canvas.DrawSdlRect(*tnk.Collider.Bounds().Rect())
	canvas.EndLineStyle()

	canvas.BeginFill(colors.Yellow)
	canvas.DrawPixelF(tnk.body.shape.X, tnk.body.shape.Y, 4)
	canvas.BeginFill(colors.Red)
	canvas.DrawPixelF(domeCenter.X, domeCenter.Y, 4)
	canvas.BeginFill(colors.Blue)
	canvas.DrawSdlPoint(tnk.Turret.barrelSprite.AbsoluteOrigin().SdlPoint(), 2)

	canvas.BeginFill(colors.Red)
	canvas.DrawPixelF(tnk.frontAxis.GetX(), tnk.frontAxis.GetY(), 6)
	canvas.BeginFill(colors.Red)
	canvas.DrawPixelF(tnk.backAxis.GetX(), tnk.backAxis.GetY(), 2)
}
