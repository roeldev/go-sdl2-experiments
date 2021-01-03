package geom

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Transform struct {
	Rotation       Degrees
	ScaleX, ScaleY float32
}

type Bounds struct {
	sdl.Rect
	origW, origH int32
}

func RectBounds(w, h int32) Bounds {
	return Bounds{
		Rect:  sdl.Rect{W: w, H: h},
		origW: w,
		origH: h,
	}
}

func (b *Bounds) Update(pos Pos) {
	b.X = int32(pos.X - (float32(b.origW) / 2))
	b.Y = int32(pos.Y - (float32(b.origH) / 2))
	b.W = b.origW
	b.H = b.origH
}

func (b *Bounds) Transform(pos Pos, t Transform) {
	var w, h float32
	if t.ScaleX != 0 {
		w = float32(b.origW) * t.ScaleX
		b.W = int32(w)
	} else {
		w = float32(b.origW)
		b.W = b.origW
	}
	if t.ScaleY != 0 {
		h = float32(b.origH) * t.ScaleY
		b.H = int32(h)
	} else {
		h = float32(b.origH)
		b.H = b.origH
	}

	b.X = int32(pos.X - (w / 2))
	b.Y = int32(pos.Y - (h / 2))
}
