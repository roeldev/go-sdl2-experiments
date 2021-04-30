package main

import (
	sdlgfx "github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/geom"
)

type player struct {
	poly *geom.Polygon
}

func newPlayer(x, y float64) *player {
	p := &player{
		poly: geom.NewTrigon(x, y, 25, 20),
	}
	return p
}

func (p *player) Render(ren *sdl.Renderer) error {
	vx, vy := sdlkit.ConvPolygonPoints(p.poly.Vertices(), 0, 0)
	sdlgfx.FilledPolygonRGBA(ren, vx, vy, 0, 0, 0, 255)
	sdlgfx.PolygonRGBA(ren, vx, vy, 255, 255, 255, 255)
	return nil
}
