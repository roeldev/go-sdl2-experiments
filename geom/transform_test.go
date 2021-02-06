// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package geom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScaleMatrix(t *testing.T) {
	var x, y = rng.Float64(), rng.Float64()
	matrix := ScaleMatrix(x, y)
	assert.Equal(t, Matrix{x, 0, 0, 0, y, 0, 0, 0, 1}, matrix)
}
