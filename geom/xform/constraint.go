// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xform

//goland:noinspection GoUnusedConst
const (
	ConstraintRotation Constraints = 1 << iota
	ConstraintRotateCW             // todo: implement CW/CCW constraints
	ConstraintRotateCCW
	ConstraintScaleX
	ConstraintScaleY
	ConstraintSkewX
	ConstraintSkewY
	ConstraintTranslateX
	ConstraintTranslateY

	NoConstraints       Constraints = 0
	ConstraintScale                 = ConstraintScaleX | ConstraintScaleY
	ConstraintSkew                  = ConstraintSkewX | ConstraintSkewY
	ConstraintTranslate             = ConstraintTranslateX | ConstraintTranslateY
	ConstraintAll                   = ConstraintRotation | ConstraintScale | ConstraintSkew | ConstraintTranslate
)

func WithConstraints(t *Transformer, constraints Constraints) *Transformer {
	t.mutex.Lock()
	t.Constraints = constraints
	t.mutex.Unlock()
	return t
}

type Constraints uint16

func (tc Constraints) Has(c Constraints) bool { return tc&c != 0 }

func (tc *Constraints) Set(c Constraints) { *tc = c }

func (tc *Constraints) Add(c Constraints) { *tc |= c }

func (tc *Constraints) Remove(c Constraints) { *tc &^= c }

func (tc *Constraints) Clear() { *tc = 0 }
