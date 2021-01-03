// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdlkit

func RunLoop(stage *Stage) error {
	sceneManager := stage.SceneManager()
	renderTarget := stage.RenderTarget()
	scene := stage.Scene()
	timer := stage.Time().Init()

	for {
		timer.Tick()

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
		scene.Update()

		// render to screen
		renderTarget.Clear()
		if err := scene.Render(renderTarget); err != nil {
			return err
		}
		if err := renderTarget.Err(); err != nil {
			return err
		}

		stage.PresentScreen()
	}
}
