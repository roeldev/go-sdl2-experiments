// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdlkit

import (
	"fmt"
	"time"

	"github.com/go-pogo/errors"
	sdlgfx "github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/go-pogo/sdlkit/geom"
)

const (
	DefaultFps       uint8   = 60
	DefaultTimeScale float32 = 1.0
)

var fsec = float32(time.Second)

type Clock struct {
	time *Time

	// DeltaTS returns the current delta time value adjusted with the current time
	// scale. This means if the time scale is set to half speed, the delta value is
	// also half it's original value.
	Delta float32

	// affects the speed of time
	// 1.0 means normal speed
	// 0.5 means twice as slow
	// 2.0 means twice as fast
	TimeScale float32
}

func NewClock() *Clock {
	return &Clock{TimeScale: DefaultTimeScale}
}

func (c *Clock) Time() *Time { return c.time }

func (c *Clock) Pause() {
	// todo: implement pause
}

func (c *Clock) Unpause() {
	// todo: implement unpause
}

type Time struct {
	targetFrameRate     uint8         // 60 fps
	targetFrameDuration time.Duration // max ticks per frame to reach targetFrameRate

	prevTick time.Time

	// delta decreases at higher fps
	// delta increases at lower fps
	// delta = 1 when fps = 1, meaning move 1 px per second
	delta float32

	// elapsed time since last tick
	elapsed time.Duration

	clocks []*Clock

	avgPerSec avgFps
	avgPerMin avgFps

	LimitFps bool
}

func NewTime(targetFps uint8, clock ...*Clock) *Time {
	t := &Time{
		avgPerSec: avgFps{after: time.Second / 2, current: float32(targetFps)},
		avgPerMin: avgFps{after: time.Second * 30},
		clocks:    clock,
		prevTick:  time.Now(),
	}

	t.SetTargetFps(targetFps)
	return t
}

func (t *Time) SetTargetFps(targetFps uint8) *Time {
	if targetFps < 1 {
		targetFps = DefaultFps
	}

	t.targetFrameRate = targetFps
	t.targetFrameDuration = time.Second / time.Duration(targetFps)
	return t
}

func (t *Time) RegisterClock(clock *Clock) {
	clock.time = t
	t.clocks = append(t.clocks, clock)
}

// Fps returns the average FPS of the last 500 milliseconds.
func (t *Time) Fps() float32 { return t.avgPerSec.current }

// AvgFps returns the average FPS of the last 30 seconds.
func (t *Time) AvgFps() float32 { return t.avgPerMin.current }

// Delta returns the current delta time value.
func (t *Time) Delta() float32 { return t.delta }

func (t *Time) Elapsed() time.Duration { return t.elapsed }

func (t *Time) Init() *Time {
	t.prevTick = time.Now()
	return t
}

func (t *Time) Tick() float32 {
	now := time.Now()

	if t.LimitFps {
		elapsed := now.Sub(t.prevTick)
		if elapsed < t.targetFrameDuration {
			time.Sleep(t.targetFrameDuration - elapsed - time.Millisecond)
			now = time.Now()
		}
	}

	t.elapsed = now.Sub(t.prevTick)
	t.prevTick = now
	t.delta = float32(t.elapsed) / fsec

	for _, clock := range t.clocks {
		clock.Delta = t.delta * clock.TimeScale
	}

	t.avgPerSec.update(t.elapsed)
	t.avgPerMin.update(t.elapsed)

	return t.delta
}

func (t *Time) CreateDisplay(x, y float32) *FpsDisplay {
	display := &FpsDisplay{
		time:        t,
		Scale:       2,
		TextColor:   sdl.Color{R: 255, G: 255, B: 255, A: 255},
		ShadowColor: sdl.Color{A: 100},
	}
	display.X = x
	display.Y = y
	return display
}

func (t *Time) String() string {
	if t.avgPerMin.current == 0 {
		return t.avgPerSec.String()
	}
	return t.avgPerMin.String()
}

type FpsDisplay struct {
	geom.Point
	time *Time

	Scale       float32
	TextColor   sdl.Color
	ShadowColor sdl.Color
}

func (d *FpsDisplay) Draw(r *sdl.Renderer) (err error) {
	var x, y int32
	var sx, sy float32

	if d.Scale > 1 {
		sx, sy = r.GetScale()
		errors.Append(&err, r.SetScale(d.Scale, d.Scale))
		x = int32(d.Point.X / d.Scale)
		y = int32(d.Point.Y / d.Scale)
	} else {
		x = int32(d.Point.X)
		y = int32(d.Point.Y)
	}

	fps := fmt.Sprintf("%.2f", d.time.Fps())
	sdlgfx.StringColor(r, x+1, x+1, fps, d.ShadowColor) // shadow
	sdlgfx.StringColor(r, x, y, fps, d.TextColor)

	if d.Scale > 1 {
		errors.Append(&err, r.SetScale(sx, sy))
	}
	return err
}

type avgFps struct {
	after   time.Duration
	elapsed time.Duration

	count   uint16
	current float32
	highest float32
	lowest  float32
}

func (f *avgFps) update(elapsed time.Duration) {
	f.count++
	f.elapsed += elapsed

	if f.elapsed < f.after {
		return
	}

	f.current = (float32(f.count) / float32(f.elapsed)) * float32(time.Second)
	if f.current > f.highest || f.highest == 0 {
		f.highest = f.current
	}
	if f.current < f.lowest || f.lowest == 0 {
		f.lowest = f.current
	}

	f.elapsed = 0
	f.count = 0
}

func (f *avgFps) String() string {
	return fmt.Sprintf("fps: %.2f, high: %.2f, low: %.2f", f.current, f.highest, f.lowest)
}
