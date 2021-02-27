// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xform

import (
	"math"

	"github.com/go-pogo/sdlkit/geom"
)

var NoLimits = Limits{
	Rotation:   Limit{-math.MaxFloat64, math.MaxFloat64},
	ScaleX:     Limit{0, math.MaxFloat64},
	ScaleY:     Limit{0, math.MaxFloat64},
	SkewX:      Limit{-math.MaxFloat64, math.MaxFloat64},
	SkewY:      Limit{-math.MaxFloat64, math.MaxFloat64},
	TranslateX: Limit{-math.MaxFloat64, math.MaxFloat64},
	TranslateY: Limit{-math.MaxFloat64, math.MaxFloat64},
}

func WithLimits(t *Transformer, limits Limits) *Transformer {
	t.mutex.Lock()
	t.Limits = limits
	t.mutex.Unlock()
	return t
}

type Limits struct {
	Rotation,
	ScaleX,
	ScaleY,
	SkewX,
	SkewY,
	TranslateX,
	TranslateY Limit
}

func (l Limits) Apply(target *geom.Transform) {
	l.Rotation.ClampPtr(&target.Rotation)
	l.ScaleX.ClampPtr(&target.ScaleX)
	l.ScaleY.ClampPtr(&target.ScaleY)
	l.SkewX.ClampPtr(&target.SkewX)
	l.SkewY.ClampPtr(&target.SkewY)
	l.TranslateX.ClampPtr(&target.TranslateX)
	l.TranslateY.ClampPtr(&target.TranslateY)
}

type Limit struct {
	Min, Max float64
}

func (l *Limit) Set(min, max float64) {
	l.Min = min
	l.Max = max
}

func (l Limit) Clamp(val float64) float64 {
	if val < l.Min {
		return l.Min
	}
	if val > l.Max {
		return l.Max
	}

	return val
}

func (l Limit) ClampPtr(val *float64) bool {
	if *val < l.Min {
		*val = l.Min
		return true
	} else if *val > l.Max {
		*val = l.Max
		return true
	}
	return false
}
