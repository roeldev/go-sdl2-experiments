// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdlkit

func RunLoop(stage *Stage) error {
	sceneManager := stage.SceneManager()
	renderer := stage.Renderer()

	scene := stage.Scene()
	timer := stage.Time().Init()

	for {
		dt := timer.Tick()

		// handle events
		if err := scene.Process(); err != nil {
			return err
		}

		// returns true when a scene switch has happened
		// this means we should process new events before
		// updating and rendering
		if sceneManager.UpdateActiveScene(&scene) {
			continue
		}

		// update state of scene
		scene.Update(dt)

		// render to screen
		if err := stage.ClearScreen(); err != nil {
			return err
		}
		if err := scene.Render(renderer); err != nil {
			return err
		}

		stage.PresentScreen()
	}
}
