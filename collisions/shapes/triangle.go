// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/colors"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/ecs"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/geom"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/physics"
)

type triangle struct {
	color sdl.Color
	shape geom.PolygonShape
	body  *physics.DynamicBody
}

func newTriangle() *triangle {
	shape := geom.NewRegularPolygon(800, 650, 40, 3)
	return &triangle{
		color: colors.RandColor(sdlkit.RNG()),
		shape: shape,
		body:  physics.NewDynamicBody(shape, nil),
	}
}

func (t *triangle) AddComponent(tag ecs.ComponentTag, component interface{}) ecs.Container {
	switch tag {
	case BodyComponent:
		t.body = component.(*physics.DynamicBody)
	// case TransformComponent:
	// 	return true
	case ColorComponent:
		t.color = component.(sdl.Color)
	}
	return ecs.Container(t)
}

func (t *triangle) HasComponent(tag ecs.ComponentTag) bool {
	switch tag {
	case BodyComponent:
		return true
	// case TransformComponent:
	// 	return true
	case ColorComponent:
		return true
	}

	return false
}

func (t *triangle) Component(tag ecs.ComponentTag) interface{} {
	switch tag {
	case BodyComponent:
		return t.body
	// case TransformComponent:
	// 	return geom.Transformer(t.body)
	case ColorComponent:
		return t.color
	}
	return nil
}

func (t *triangle) Components(tag ...ecs.ComponentTag) map[ecs.ComponentTag]interface{} {
	res := make(map[ecs.ComponentTag]interface{}, len(tag))
	for _, ct := range tag {
		res[ct] = t.Component(ct)
	}
	return res
}
