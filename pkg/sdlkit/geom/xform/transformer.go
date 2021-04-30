// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package xform supplies additional transform related features.
package xform

import (
	"sync"

	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/geom"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/math"
)

type Transformer struct {
	Limits      Limits
	Constraints Constraints

	mutex      sync.RWMutex
	target     *geom.Transform
	hasChanged bool
	prevMatrix geom.Matrix
}

func New(transform *geom.Transform) *Transformer {
	return &Transformer{
		target:      transform,
		Limits:      NoLimits,
		Constraints: NoConstraints,
	}
}

func NewTransformer() *Transformer {
	return New(&geom.Transform{ScaleX: 1, ScaleY: 1})
}

func (t *Transformer) Locker() *sync.RWMutex { return &t.mutex }

func (t *Transformer) setRotation(radians float64) {
	radians = t.Limits.Rotation.Clamp(radians)
	if t.target.Rotation != radians {
		t.target.Rotation = radians
		t.hasChanged = true
	}
}

func (t *Transformer) SetRotationDeg(degrees float64) *Transformer {
	return t.SetRotation(degrees * math.D2R)
}

// Rotate adds the given radians to the rotation transformation.
func (t *Transformer) Rotate(radians float64) *Transformer {
	return t.SetRotation(t.target.Rotation + radians)
}

// RotateDeg adds the given degrees to the rotation transformation.
func (t *Transformer) RotateDeg(degrees float64) *Transformer {
	return t.SetRotation(t.target.Rotation + (degrees * math.D2R))
}

func (t *Transformer) setScaleX(scale float64) {
	scale = t.Limits.ScaleX.Clamp(scale)
	if t.target.ScaleX != scale {
		t.target.ScaleX = scale
		t.hasChanged = true
	}
}

func (t *Transformer) setScaleY(scale float64) {
	scale = t.Limits.ScaleY.Clamp(scale)
	if t.target.ScaleY != scale {
		t.target.ScaleY = scale
		t.hasChanged = true
	}
}

func (t *Transformer) ScaleX(add float64) *Transformer {
	return t.SetScaleX(t.target.ScaleX + add)
}

func (t *Transformer) ScaleY(add float64) *Transformer {
	return t.SetScaleY(t.target.ScaleY + add)
}

func (t *Transformer) setSkewX(skew float64) {
	skew = t.Limits.SkewX.Clamp(skew)
	if t.target.SkewX != skew {
		t.target.SkewX = skew
		t.hasChanged = true
	}
}

func (t *Transformer) setSkewY(skew float64) {
	skew = t.Limits.SkewY.Clamp(skew)
	if t.target.SkewY != skew {
		t.target.SkewY = skew
		t.hasChanged = true
	}
}

func (t *Transformer) SkewX(add float64) *Transformer {
	return t.SetSkewX(t.target.SkewX + add)
}

func (t *Transformer) SkewY(add float64) *Transformer {
	return t.SetSkewY(t.target.SkewY + add)
}

func (t *Transformer) setTranslateX(translate float64) {
	translate = t.Limits.TranslateX.Clamp(translate)
	if t.target.TranslateX != translate {
		t.target.TranslateX = translate
		t.hasChanged = true
	}
}

func (t *Transformer) setTranslateY(translate float64) {
	translate = t.Limits.TranslateY.Clamp(translate)
	if t.target.TranslateY != translate {
		t.target.TranslateY = translate
		t.hasChanged = true
	}
}

func (t *Transformer) TranslateX(add float64) *Transformer {
	return t.SetTranslateX(t.target.TranslateX + add)
}

func (t *Transformer) TranslateY(add float64) *Transformer {
	return t.SetTranslateY(t.target.TranslateY + add)
}

func (t *Transformer) HasChanged() bool {
	t.mutex.RLock()
	res := t.hasChanged
	t.mutex.RUnlock()
	return res
}

func (t *Transformer) Transform() geom.Transform {
	t.mutex.RLock()
	res := *t.target
	t.mutex.RUnlock()
	return res
}

func (t *Transformer) Reset() {
	t.mutex.Lock()
	t.target.Reset()
	t.hasChanged = true
	t.mutex.Unlock()
}

func (t *Transformer) Matrix() geom.Matrix {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if t.hasChanged {
		t.prevMatrix = t.target.Matrix()
		t.hasChanged = false
	}
	return t.prevMatrix
}
