package geom

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Pos struct{ sdl.FPoint }

func (p Pos) Point() sdl.Point {
	return sdl.Point{X: int32(p.X), Y: int32(p.Y)}
}

func (p *Pos) Center(x, y, w, h float32) {
	p.X = x + (w / 2)
	p.Y = y + (h / 2)
}

func (p *Pos) CenterRect(area sdl.Rect) {
	p.Center(
		float32(area.X),
		float32(area.Y),
		float32(area.W),
		float32(area.H),
	)
}

func (p *Pos) CenterFRect(area sdl.FRect) {
	p.Center(area.X, area.Y, area.W, area.H)
}
