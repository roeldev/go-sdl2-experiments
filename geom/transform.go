// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package geom

import (
	"math"
)

// A Transformable is any object that can be transformed using a transform Matrix.
type Transformable interface {
	Origin() *Point
	Transform(matrix Matrix)
}

type Transformer interface {
	Matrix() Matrix
	ApplyTransform(target Transformable)
}

// Matrix is a 3x3 matrix of float64 values.
type Matrix [9]float64
type TransformMatrixElement int

//goland:noinspection GoSnakeCaseUsage
const (
	ME_A  TransformMatrixElement = 0
	ME_B                         = 3
	ME_C                         = 1
	ME_D                         = 4
	ME_TX                        = 2
	ME_TY                        = 5
)

// TransformMatrix creates a Matrix with the following elements:
//   [ a, c, tx,
//     b, d, ty,
//     u, v, w ]
//
// The elements A and D affect the positioning of pixels along the x and y axis
// when scaling or rotating. B and C are the elements that affect the
// positioning of pixels along the x and y axis when rotating or skewing.
// TX and TY are the distances by which to translate each point along the x and
// y axis.
// The Matrix operates in 2D space so it always assumes that the (last three)
// elements u and v are 0.0, and w is 1.0.
// https://en.wikipedia.org/wiki/Transformation_matrix
// https://www.tutorialspoint.com/computer_graphics/2d_transformation.htm
func TransformMatrix(a, b, c, d, tx, ty float64) Matrix {
	return Matrix{a, c, tx, b, d, ty, 0, 0, 1}
}

func ScaleMatrix(x, y float64) Matrix {
	return Matrix{x, 0, 0, 0, y, 0, 0, 0, 1}
}

// RotationMatrix creates a Matrix which rotates the target by an angle,
// measured in radians.
// https://en.wikipedia.org/wiki/Rotation_matrix
func RotationMatrix(radians float64) Matrix {
	cos, sin := math.Cos(radians), math.Sin(radians)
	return TransformMatrix(cos, sin, -sin, cos, 0, 0)
}

// SkewMatrix creates a Matrix which progressively slides the target in a
// direction parallel to the x or y axis.
func SkewMatrix(x, y float64) Matrix {
	return Matrix{0, math.Tan(x), 0, math.Tan(y), 1, 0, 0, 0, 1}
}

// IdentityMatrix creates a Matrix with values that cause no transformation to
// the target.
func IdentityMatrix() Matrix {
	return Matrix{1, 0, 0, 0, 1, 0, 0, 0, 1}
}

// Transform calculates various transformations and merges them to a Matrix
// which can be applied to a Transformable.
type Transform struct {
	Rotation,
	ScaleX, ScaleY,
	// SkewX, SkewY,
	TranslateX, TranslateY float64
}

func NewTransform() *Transform {
	return &Transform{ScaleX: 1, ScaleY: 1}
}

func (t Transform) Matrix() Matrix {
	cos, sin := math.Cos(t.Rotation), math.Sin(t.Rotation)
	return TransformMatrix(
		cos*t.ScaleX,
		sin*t.ScaleY,
		-sin*t.ScaleX,
		cos*t.ScaleY,
		t.TranslateX,
		t.TranslateY,
	)
}

func (t Transform) ApplyTransform(target Transformable) {
	target.Transform(t.Matrix())
}

//goland:noinspection GoUnusedConst
const (
	ConstraintRotation = 1 << iota
	ConstraintScaleX
	ConstraintScaleY
	ConstraintSkewX
	ConstraintSkewY
	ConstraintTranslateX
	ConstraintTranslateY

	ConstraintScale     = ConstraintScaleX | ConstraintScaleY
	ConstraintSkew      = ConstraintSkewX | ConstraintSkewY
	ConstraintTranslate = ConstraintTranslateX | ConstraintTranslateY
	ConstraintAll       = ConstraintRotation | ConstraintScale | ConstraintSkew | ConstraintTranslate
)

type TransformConstraint uint8

func (tc TransformConstraint) Has(c TransformConstraint) bool { return tc&c != 0 }

func (tc *TransformConstraint) Set(c TransformConstraint) { *tc = c }

func (tc *TransformConstraint) Add(c TransformConstraint) { *tc |= c }

func (tc *TransformConstraint) Remove(c TransformConstraint) { *tc &^= c }

func (tc *TransformConstraint) Clear() { *tc = 0 }

type ConstraintTransform struct {
	transform   Transform
	constraints TransformConstraint
	limitRotation,
	limitScale,
	limitSkew,
	limitTranslate [2]float64
}

func (tc *ConstraintTransform) TransformConstraints() *TransformConstraint {
	return &tc.constraints
}

func (tc *ConstraintTransform) LimitRotation(min, max float64) {
	tc.limitRotation[0] = min
	tc.limitRotation[1] = max
}

// Rotation sets the rotation transformation to the given amount of radians.
func (tc *ConstraintTransform) Rotation(radians float64) {
	if !tc.constraints.Has(ConstraintRotation) {
		tc.transform.Rotation = radians
	}
}

// Rotate adds the given radians to the rotation transformation.
func (tc *ConstraintTransform) Rotate(radians float64) {
	if !tc.constraints.Has(ConstraintRotation) {
		tc.transform.Rotation += radians
	}
}

// RotateDeg adds the given degrees to the rotation transformation.
func (tc *ConstraintTransform) RotateDeg(degrees float64) {
	if !tc.constraints.Has(ConstraintRotation) {
		tc.transform.Rotation += degrees * D2R
	}
}

func (tc *ConstraintTransform) LimitScale(min, max float64) {
	tc.limitScale[0] = min
	tc.limitScale[1] = max
}

func (tc *ConstraintTransform) Scale(scale float64) {
	tc.ScaleX(scale)
	tc.ScaleY(scale)
}

func (tc *ConstraintTransform) ScaleX(scaleX float64) {
	if !tc.constraints.Has(ConstraintScaleX) {
		tc.transform.ScaleX = scaleX
	}
}

func (tc *ConstraintTransform) ScaleY(scaleY float64) {
	if !tc.constraints.Has(ConstraintScaleY) {
		tc.transform.ScaleY = scaleY
	}
}

func (tc *ConstraintTransform) LimitSkew(min, max float64) {
	tc.limitSkew[0] = min
	tc.limitSkew[1] = max
}

func (tc *ConstraintTransform) LimitTranslate(min, max float64) {
	tc.limitTranslate[0] = min
	tc.limitTranslate[1] = max
}

func (tc *ConstraintTransform) Translate(translate float64) {
	tc.TranslateX(translate)
	tc.TranslateY(translate)
}

func (tc *ConstraintTransform) TranslateX(translateX float64) {
	if !tc.constraints.Has(ConstraintTranslateX) {
		tc.transform.TranslateX = translateX
	}
}

func (tc *ConstraintTransform) TranslateY(translateY float64) {
	if !tc.constraints.Has(ConstraintTranslateY) {
		tc.transform.TranslateY = translateY
	}
}

func (tc *ConstraintTransform) Matrix() Matrix {
	return tc.transform.Matrix()
}

func (tc *ConstraintTransform) ApplyTransform(t Transformable) {
	t.Transform(tc.transform.Matrix())
}
