// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !prod

package xform

// SetRotation sets the rotation transformation to the given amount of radians.
func (t *Transformer) SetRotation(radians float64) *Transformer {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if t.Constraints.Has(ConstraintRotation) {
		panic("xform: unable to SetRotation due to constraints")
	}

	t.setRotation(radians)
	return t
}

func (t *Transformer) SetScaleX(scale float64) *Transformer {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if t.Constraints.Has(ConstraintScaleX) {
		panic("xform: unable to SetScaleX due to constraints")
	}

	t.setScaleX(scale)
	return t
}

func (t *Transformer) SetScaleY(scale float64) *Transformer {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if !t.Constraints.Has(ConstraintScaleY) {
		panic("xform: unable to SetScaleY due to constraints")
	}

	t.setScaleY(scale)
	return t
}

func (t *Transformer) SetSkewX(skew float64) *Transformer {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if !t.Constraints.Has(ConstraintSkewX) {
		panic("xform: unable to SetSkewX due to constraints")
	}

	t.setSkewY(skew)
	return t
}

func (t *Transformer) SetSkewY(skew float64) *Transformer {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if !t.Constraints.Has(ConstraintSkewY) {
		panic("xform: unable to SetSkewY due to constraints")
	}

	t.setSkewY(skew)
	return t
}

func (t *Transformer) SetTranslateX(translate float64) *Transformer {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if !t.Constraints.Has(ConstraintTranslateX) {
		panic("xform: unable to SetTranslateX due to constraints")
	}

	t.setTranslateX(translate)
	return t
}

func (t *Transformer) SetTranslateY(translate float64) *Transformer {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if !t.Constraints.Has(ConstraintTranslateY) {
		panic("xform: unable to SetTranslateY due to constraints")
	}

	t.setTranslateY(translate)
	return t
}
