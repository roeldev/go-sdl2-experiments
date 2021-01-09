// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"time"

	"github.com/roeldev/go-x11colors"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/go-pogo/sdlkit"
)

type world struct {
	stage *sdlkit.Stage
	title string

	paused     bool
	pauseAfter time.Duration

	draw  []sdlkit.Drawable
	boxes []*box
}

func newWorld(stage *sdlkit.Stage) *world {
	return &world{
		stage: stage,
		title: stage.Window().GetTitle(),
	}
}

func (w *world) setup() {
	w.paused = false
	w.pauseAfter = time.Second * 10

	w.boxes = []*box{
		newBox(20, 20, 100, 100, 100, x11colors.Rand(sdlkit.RNG())),
		newBox(50, 50, 300, 200, -50, x11colors.Rand(sdlkit.RNG())),
	}

	w.draw = []sdlkit.Drawable{
		drawRect(24, 24, 338, 17),
		drawRect(54, 54, 198, 199),
	}

	for _, box := range w.boxes {
		w.draw = append(w.draw, box)
	}

	w.draw = append(w.draw, w.stage.Time().CreateDisplay(10, 10))
}

func (w *world) Pause(pause bool) {
	w.paused = pause

	if w.paused {
		w.stage.Window().SetTitle(w.title + " - " + w.stage.Time().String())
	} else {
		w.stage.Window().SetTitle(w.title)
	}
}

// Run starts and runs the custom game loop for this example world.
func (w *world) Run() error {
	w.setup()

	renderer := w.stage.Renderer()
	timer := w.stage.Time().Init()
	var timeCount float32

Loop:
	for {
		dt := timer.Tick()

		// process input/events
		for {
			event := sdl.PollEvent()
			if event == nil {
				break
			}

			switch event := event.(type) {
			case *sdl.QuitEvent:
				return nil

			case *sdl.KeyboardEvent:
				if event.Type != sdl.KEYUP {
					continue
				}

				switch event.Keysym.Scancode {
				case sdl.SCANCODE_SPACE:
					fallthrough
				case sdl.SCANCODE_KP_SPACE:
					w.Pause(!w.paused)

				case sdl.SCANCODE_RETURN:
					fallthrough
				case sdl.SCANCODE_RETURN2:
					fallthrough
				case sdl.SCANCODE_KP_ENTER:
					timer.LimitFps = !timer.LimitFps

				case sdl.SCANCODE_ESCAPE:
					if err := w.stage.ClearScreen(); err != nil {
						return err
					}

					w.stage.PresentScreen()
					w.setup()
					timeCount = 0
					goto Loop
				}
			}
		}

		if w.paused {
			// skip updates + drawing
			continue
		}

		// update game logic
		vp := *w.stage.Viewport()
		for _, box := range w.boxes {
			box.update(dt, vp)
		}

		// random delay to simulate heavy game logic calculations
		// time.Sleep(time.Millisecond * time.Duration(5+rand.Intn(15)))

		// render to window
		if err := w.stage.ClearScreen(); err != nil {
			return err
		}
		if err := sdlkit.Draw(renderer, w.draw...); err != nil {
			return err
		}

		w.stage.Renderer().Present()

		// automatically pause the world after X seconds
		if w.pauseAfter > 0 {
			timeCount += float32(time.Second) * dt
			if time.Duration(timeCount) >= w.pauseAfter {
				w.Pause(true)
				w.pauseAfter = -1
			}
		}
	}
}
