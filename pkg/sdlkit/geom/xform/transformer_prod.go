// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build prod

package xform

// SetRotation sets the rotation transformation to the given amount of radians.
func (t *Transformer) SetRotation(radians float64) *Transformer {
	t.mutex.Lock()
	if !t.Constraints.Has(ConstraintRotation) {
		t.setRotation(radians)
	}
	t.mutex.Unlock()
	return t
}

func (t *Transformer) SetScaleX(scale float64) *Transformer {
	t.mutex.Lock()
	if !t.Constraints.Has(ConstraintScaleX) {
		t.setScaleX(scale)
	}
	t.mutex.Unlock()
	return t
}

func (t *Transformer) SetScaleY(scale float64) *Transformer {
	t.mutex.Lock()
	if !t.Constraints.Has(ConstraintScaleY) {
		t.setScaleY(scale)
	}
	t.mutex.Unlock()
	return t
}

func (t *Transformer) SetSkewX(skew float64) *Transformer {
	t.mutex.Lock()
	if !t.Constraints.Has(ConstraintSkewX) {
		t.setSkewX(skew)
	}
	t.mutex.Unlock()
	return t
}

func (t *Transformer) SetSkewY(skew float64) *Transformer {
	t.mutex.Lock()
	if !t.Constraints.Has(ConstraintSkewY) {
		t.setSkewY(skew)
	}
	t.mutex.Unlock()
	return t
}

func (t *Transformer) SetTranslateX(translate float64) *Transformer {
	t.mutex.Lock()
	if !t.Constraints.Has(ConstraintTranslateX) {
		t.setTranslateX(translate)
	}
	t.mutex.Unlock()
	return t
}

func (t *Transformer) SetTranslateY(translate float64) *Transformer {
	t.mutex.Lock()
	if !t.Constraints.Has(ConstraintTranslateY) {
		t.setTranslateY(translate)
	}
	t.mutex.Unlock()
	return t
}
