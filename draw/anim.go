package draw

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/go-pogo/sdlkit"
	"github.com/go-pogo/sdlkit/debug"
	"github.com/go-pogo/sdlkit/geom"
)

// https://docs.godotengine.org/en/stable/classes/class_spriteframes.html#class-spriteframes

type AnimSprite struct {
	geom.Point // center of sprite
	geom.Transform
	sdlkit.TextureDisplay
	Bounds geom.Bounds

	// vervangen voor spriteframes
	// spriteframes kunnen uit een spritesheet komen
	sheet        *SpriteSheet
	currentFrame int
	totalFrames  int

	rate time.Duration
	time time.Duration

	done   bool
	Repeat bool
}

func NewAnimSprite(sheet *SpriteSheet, duration time.Duration) *AnimSprite {
	a := &AnimSprite{
		Bounds: geom.RectBounds(sheet.indexes[0].W, sheet.indexes[0].H),
		sheet:  sheet,
		rate:   duration / time.Duration(sheet.Len()),
	}
	a.Alpha = 255
	return a
}

func (a *AnimSprite) Finished() bool { return a.done }

func (a *AnimSprite) Frame() (*sdl.Texture, sdl.Rect) {
	if a.done {
		return nil, sdl.Rect{}
	}

	return a.sheet.tx, a.sheet.indexes[a.currentFrame]
}

func (a *AnimSprite) Rewind() {
	a.currentFrame = 0
	a.done = false
}

func (a *AnimSprite) Update(dt float32) {
	a.Bounds.Transform(a.Point, a.Transform)

	if a.done {
		return
	}

	a.time += time.Duration(float32(time.Second) * dt)
	if a.time > a.rate {
		a.currentFrame++
		a.time = 0

		if a.currentFrame >= a.sheet.Len() {
			if a.Repeat {
				a.currentFrame = 0
			} else {
				a.done = true
			}
		}
	}
}

func (a *AnimSprite) Draw(r *sdl.Renderer) (err error) {
	tx, src := a.Frame()
	if tx != nil {
		// sdlkit.TextureDisplay{ niet standaard in anim? transform ook niet? }
		err = DrawTexture(r, tx, &src, &a.Bounds.Rect, a.Transform, a.TextureDisplay)
	}

	debug.DrawPos(r, a.Point)
	debug.DrawBounds(r, a.Bounds)

	return err
}
