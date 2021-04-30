package main

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit"
)

type starField struct {
	texture *sdl.Texture

	density int32
	size    sdl.Rect
}

func newStarField(density int32) *starField {
	return &starField{
		density: density,
	}
}

func (sf *starField) Render2(ren *sdl.Renderer, size sdl.Rect) error {
	canvas := sdlkit.NewCanvas(ren)
	tx, err := canvas.CreateTexture(sdl.PIXELFORMAT_RGBA8888, 0, size.W, size.H)
	if err != nil {
		return err
	}

	canvas.BlendMode(sdl.BLENDMODE_BLEND)
	rng := sdlkit.RNG()

	var x, y int32 = 0, 0
	for y = 0; y < size.H; y += sf.density {
		for x = 0; x < size.W; {
			x += rng.Int31n(sf.density)
			canvas.BeginFillRGBA(255, 255, 255, uint8(55+rng.Intn(200)))
			canvas.DrawPixel(x, y+rng.Int31n(sf.density*2)-sf.density, rng.Int31n(4))
		}
	}

	if err = canvas.Done(); err != nil {
		return err
	}

	sf.texture = tx
	sf.size.W = size.W
	sf.size.H = size.H
	return nil
}

func (sf *starField) Render(ren *sdl.Renderer) error {
	if sf.texture == nil {
		return nil
	}

	return ren.Copy(sf.texture, nil, &sf.size)
}
