// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tank

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/display"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/geom"
)

var (
	turretColors = map[Color][2]sdl.Color{
		Green: {sdl.Color{R: 0x17, G: 0x70, B: 0x3e, A: 0xff}, sdl.Color{R: 0x23, G: 0xa1, B: 0x58, A: 0xff}},
		Blue:  {sdl.Color{R: 0x26, G: 0x61, B: 0x87, A: 0xff}, sdl.Color{R: 0x31, G: 0x7a, B: 0xaa, A: 0xff}},
		Red:   {sdl.Color{R: 0x8c, G: 0x2c, B: 0x23, A: 0xff}, sdl.Color{R: 0xb4, G: 0x3a, B: 0x2d, A: 0xff}},
		Sand:  {sdl.Color{R: 0xa1, G: 0x95, B: 0x75, A: 0xff}, sdl.Color{R: 0xb8, G: 0xaa, B: 0x87, A: 0xff}},
	}
)

type Turret struct {
	offset       geom.Vector
	color        Color
	domeRadius   [2]int32
	domeSprite   *display.Sprite
	barrelSprite *display.Sprite

	curRotation,
	destRotation float64
}

func NewTurretSmall(atlas *sdlkit.TextureAtlas, tank *Tank) Turret {
	barrelSprite := display.MustNewSprite(tank.body.color.BarrelTextureClip(atlas))
	// barrelSprite.TranslateX = 13
	// barrelSprite.Origin().Y = 2

	return Turret{
		offset:       geom.Vector{X: 4},
		color:        tank.body.color,
		domeRadius:   [2]int32{9, 7},
		barrelSprite: barrelSprite,
	}
}

func (tur *Turret) Prerender(canvas *sdlkit.Canvas) error {
	if tur.domeSprite != nil {
		_ = tur.domeSprite.Clip().Texture.Destroy()
	}

	clip, err := canvas.CreateTextureClip(
		sdl.PIXELFORMAT_RGBA8888,
		sdl.TEXTUREACCESS_TARGET,
		tur.domeRadius[0]*2,
		tur.domeRadius[0]*2,
	)
	if err != nil {
		return err
	}

	canvas.BeginFill(turretColors[tur.color][0])
	canvas.DrawCircle(tur.domeRadius[0], tur.domeRadius[0], tur.domeRadius[0]*2)
	canvas.BeginFill(turretColors[tur.color][1])
	canvas.DrawCircle(tur.domeRadius[0], tur.domeRadius[0], tur.domeRadius[1]*2)
	canvas.EndFill()

	// todo: fix dome sprite dest size
	tur.domeSprite = display.NewSprite(clip)
	tur.domeSprite.W /= 2
	tur.domeSprite.H /= 2

	return canvas.Done()
}

func (tur *Turret) update(x, y float64) {
	tur.barrelSprite.X = x
	tur.barrelSprite.Y = y //+ tur.barrelSprite.TranslateY

	if tur.domeSprite != nil {
		tur.domeSprite.X = x
		tur.domeSprite.Y = y
	}
}
