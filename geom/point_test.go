package geom

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/veandco/go-sdl2/sdl"
)

func TestInCircle(t *testing.T) {
	tests := []struct {
		x, y, cx, cy, rad float32
		ok                bool
	}{
		{1, 1, 1, 1, 2, true},
		{3, 3, 0, 0, 1, false},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%.2fx%.2f", tc.x, tc.y), func(t *testing.T) {
			assert.Equal(t, tc.ok, InCircle(tc.x, tc.y, tc.cx, tc.cy, tc.rad))
			assert.Equal(t, tc.ok, (Point{X: tc.x, Y: tc.y}).InCircle(tc.cx, tc.cy, tc.rad))
		})
	}
}

func TestInRect(t *testing.T) {
	tests := []struct {
		x, y, rx, ry, rw, rh float32
		ok                   bool
	}{
		{1, 2, 0, 0, 22, 100, true},
		{10, 10, 0, 0, 10, 50, false},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%.2fx%.2f", tc.x, tc.y), func(t *testing.T) {
			assert.Equal(t, tc.ok, InRect(tc.x, tc.y, tc.rx, tc.ry, tc.rw, tc.rh))
			r := sdl.FRect{X: tc.rx, Y: tc.ry, W: tc.rw, H: tc.rh}
			assert.Equal(t, tc.ok, Point{X: tc.x, Y: tc.y}.InFRect(r))
		})
	}
}
