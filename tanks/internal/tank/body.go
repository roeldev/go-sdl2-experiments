// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tank

import (
	"github.com/go-pogo/errors"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/display"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/geom"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/geom/xform"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/physics"
)

type Color uint8

const (
	Green Color = iota
	Blue
	Red
	Sand
)

func (c Color) BodyTextureClip(atlas *sdlkit.TextureAtlas) (sdlkit.TextureClip, error) {
	switch c {
	case Green:
		return atlas.GetFromName("tankBody_green_outline.png")
	case Blue:
		return atlas.GetFromName("tankBody_blue_outline.png")
	case Red:
		return atlas.GetFromName("tankBody_red_outline.png")
	case Sand:
		return atlas.GetFromName("tankBody_sand_outline.png")
	}

	return sdlkit.TextureClip{}, errors.Newf("tank: color `%d` is not defined", c)
}

func (c Color) BarrelTextureClip(atlas *sdlkit.TextureAtlas) (sdlkit.TextureClip, error) {
	switch c {
	case Green:
		return atlas.GetFromName("tankGreen_barrel2_outline.png")
	case Blue:
		return atlas.GetFromName("tankBlue_barrel2_outline.png")
	case Red:
		return atlas.GetFromName("tankRed_barrel2_outline.png")
	case Sand:
		return atlas.GetFromName("tankSand_barrel2_outline.png")
	}

	return sdlkit.TextureClip{}, errors.Newf("tank: color `%d` is not defined", c)
}

type body struct {
	physics.Collider

	width, height float64

	color     Color
	sprite    *display.Sprite
	shape     *geom.Polygon
	transform *xform.Transformer
}

func newBody(atlas *sdlkit.TextureAtlas, color Color) *body {
	sprite := display.MustNewSprite(color.BodyTextureClip(atlas))

	if color == Green || color == Blue {
		// flip the texture so the rear is used as front, in my opinion this
		// looks a little bit better
		sprite.Flip = sdl.FLIP_VERTICAL
	}

	// quad (rectangle like) polygon that can be transformed (rotated)
	w, h := sprite.Clip().Size()
	shape := geom.NewQuad(0, 0, w, h)

	return &body{
		Collider: physics.NewCollider(shape),

		width:  w,
		height: h,
		color:  color,
		sprite: sprite,
		shape:  shape,

		// tank body is only allowed to rotate, constraint everything else
		transform: xform.WithConstraints(
			xform.NewTransformer(),
			xform.ConstraintAll&^xform.ConstraintRotation,
		),
	}
}
