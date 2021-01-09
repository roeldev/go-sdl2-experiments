// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package geom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVector_Angle(t *testing.T) {
	v := Vector{10, 5}
	assert.Equal(t, 0.4636476090008061, float64(v.Angle()))
	assert.Equal(t, 11.180339887498949, v.Length())
}

func TestVector_Length(t *testing.T) {
	v := Vector{1, 1}
	assert.Equal(t, 1.4142135623730951, v.Length())

	v2 := Vector{2, 2}
	assert.Equal(t, v.Length()*2, v2.Length())
}

func TestVector_SetLength(t *testing.T) {
	var v Vector
	v.SetLength(10)
	assert.Equal(t, Vector{X: 10}, v)
}

func TestVector_Normalize(t *testing.T) {
	original := &Vector{10, 10}
	normalized := original.Normalize()

	assert.Equal(t, 1, int(normalized.Length()))
	assert.Equal(t, &Vector{10, 10}, original, "original Vector should not have changed")
}
