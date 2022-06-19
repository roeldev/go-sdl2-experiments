// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdlkit

import (
	"context"
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	DefaultFps       uint8   = 60
	DefaultTimeScale float64 = 1.0
)

var fsec = float64(time.Second)

type Clock struct {
	time *Time

	// TimeScale affects the speed of time. Its default value is 1.0.
	// When TimeScale < 1, time slows down. Time speeds up when TimeScale > 1.
	TimeScale float64

	// Delta64 returns the current delta time value multiplied by TimeScale.
	Delta64 float64

	// Delta32 is a float32 version of Delta64.
	Delta32 float32
}

// todo: clock aanmaken via Time.CreateClock
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

	startTick uint32
	startTime time.Time
	prevTime  time.Time

	// delta decreases at higher fps
	// delta increases at lower fps
	// delta = 1 when fps = 1, meaning move 1 px per second
	delta float64

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
		startTick: sdl.GetTicks(),
		startTime: time.Now(),
	}

	t.prevTime = t.startTime
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

// ConvTicks coverts a ticks value from sdl.GetTicks to a time.Time value.
// The result may be a few microseconds off but is well below a millisecond.
func (t *Time) ConvTicks(ticks uint32) time.Time {
	return t.startTime.Add(time.Duration(ticks-t.startTick) * time.Millisecond)
}

// Fps returns the average FPS of the last 500 milliseconds.
func (t *Time) Fps() float32 { return t.avgPerSec.current }

// AvgFps returns the average FPS of the last 30 seconds.
func (t *Time) AvgFps() float32 { return t.avgPerMin.current }

func (t *Time) Elapsed() time.Duration { return t.elapsed }

func (t *Time) Init() *Time {
	t.prevTime = time.Now()
	return t
}

func (t *Time) Tick() float64 {
	now := time.Now()

	if t.LimitFps {
		elapsed := now.Sub(t.prevTime)
		if elapsed < t.targetFrameDuration {
			time.Sleep(t.targetFrameDuration - elapsed - time.Millisecond)
			now = time.Now()
		}
	}

	t.elapsed = now.Sub(t.prevTime)
	t.prevTime = now
	t.delta = float64(t.elapsed) / fsec

	for _, clock := range t.clocks {
		clock.Delta64 = t.delta * clock.TimeScale
		clock.Delta32 = float32(clock.Delta64)
	}

	t.avgPerSec.update(t.elapsed)
	t.avgPerMin.update(t.elapsed)

	return t.delta
}

func (t *Time) String() string {
	if t.avgPerMin.current == 0 {
		return t.avgPerSec.String()
	}
	return t.avgPerMin.String()
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

type windowTitleFps struct {
	ctx     context.Context
	cfn     context.CancelFunc
	stage   *Stage
	running bool
}

func newWindowTitleFps(stage *Stage) *windowTitleFps {
	return &windowTitleFps{stage: stage}
}

func (wt *windowTitleFps) run() {
	wt.ctx, wt.cfn = context.WithCancel(wt.stage.ctx)
	go func() {
		if wt.running {
			return
		}

		wt.running = true
		timer := time.NewTicker(time.Second / 2)
		title := wt.stage.window.GetTitle()

		for {
			wt.stage.window.SetTitle(title + " | " + wt.stage.time.avgPerSec.String())

			select {
			case <-timer.C: // wait before we update the title again
				continue

			case <-wt.ctx.Done():
				wt.running = false
				wt.stage.window.SetTitle(title)
				return
			}
		}
	}()
}

func (wt *windowTitleFps) stop() { wt.cfn() }
