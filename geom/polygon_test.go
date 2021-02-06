// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package geom

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func TestPolygon_Area(t *testing.T) {
	tests := map[string]func() (*Polygon, float64){
		"regular trigon": func() (*Polygon, float64) {
			p := NewRegularPolygon(0, 0, 10, 3)
			// https://en.wikipedia.org/wiki/Heron%27s_formula
			a := Distance(p.actual[0].X, p.actual[0].Y, p.actual[1].X, p.actual[1].Y)
			b := Distance(p.actual[1].X, p.actual[1].Y, p.actual[2].X, p.actual[2].Y)
			c := Distance(p.actual[2].X, p.actual[2].Y, p.actual[0].X, p.actual[0].Y)
			return p, math.RoundToEven(0.25 * math.Sqrt((a+b+c)*(-a+b+c)*(a-b+c)*(a+b-c)))
		},

		"trigon": func() (*Polygon, float64) {
			w := float64(1 + rng.Int31n(100))
			h := float64(1 + rng.Int31n(100))

			return NewTrigon(0, 0, w, h), w * h / 2
		},
		"quad": func() (*Polygon, float64) {
			w := float64(1 + rng.Int31n(100))
			h := float64(1 + rng.Int31n(100))

			return NewQuad(0, 0, w, h), w * h
		},
	}

	for name, setup := range tests {
		poly, want := setup()
		for i := 0; i < 5; i++ {
			t.Run(name, func(t *testing.T) {
				rotation := rng.Float64() * math.Pi * 2
				if i != 0 {
					poly.Transform(RotationMatrix(rotation))
				}

				diff := poly.Area() - want
				assert.True(t, (diff >= 0 && diff <= 0.1) || (diff <= 0 && diff >= -0.1))
			})
		}
	}
}

func TestNewQuad(t *testing.T) {
	t.Run("size", func(t *testing.T) {
		want := [2]float64{100, 50}
		edges := NewQuad(0, 0, want[0], want[1]).Edges()

		assert.Equal(t, want, [2]float64{
			math.Abs(edges[0].X - edges[2].X),
			math.Abs(edges[0].Y - edges[2].Y),
		})
		assert.Equal(t, want, [2]float64{
			math.Abs(edges[1].X - edges[3].X),
			math.Abs(edges[1].Y - edges[3].Y),
		})
	})
}
