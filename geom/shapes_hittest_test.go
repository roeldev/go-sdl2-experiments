// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package geom

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHitTest(t *testing.T) {
	tests := []struct {
		x, y   float64
		target HitTester
		ok     bool
	}{
		{1, 1, Circle{1, 1, 2}, true},
		{3, 3, Circle{0, 0, 1}, false},
		{1, 2, Rect{0, 0, 22, 100}, true},
		{10, 10, Rect{0, 0, 10, 50}, false},
	}

	for _, tc := range tests {
		name := fmt.Sprintf("%T: %.2fx%.2f", tc.target, tc.x, tc.y)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.ok, tc.target.HitTest(tc.x, tc.y))
		})
	}
}

func barycentricInTriangle(x, y, p1X, p1Y, p2X, p2Y, p3X, p3Y float64) bool {
	s := p1Y*p3X - p1X*p3Y + (p3Y-p1Y)*x + (p1X-p3X)*y
	t := p1X*p2Y - p1Y*p2X + (p1Y-p2Y)*x + (p2X-p1X)*y

	if (s < 0) != (t < 0) {
		return false
	}

	A := -p2Y*p3X + p1Y*(p3X-p2X) + p1X*(p2Y-p3Y) + p2X*p3Y

	if A < 0 {
		return s <= 0 && s+t >= A
	} else {
		return s >= 0 && s+t <= A
	}
}

func BenchmarkInTriangle(b *testing.B) {
	barycentricInTriangle(0, 0, 0, 0, 0, 0, 0, 0)
}
