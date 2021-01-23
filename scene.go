// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdlkit

import (
	"github.com/go-pogo/errors"
	"github.com/veandco/go-sdl2/sdl"
)

type Scene interface {
	Process() error
	Update(dt float64)
	Render(r *sdl.Renderer) error
}

//goland:noinspection SpellCheckingInspection
type SceneActivater interface {
	Scene
	Activate()
}

//goland:noinspection SpellCheckingInspection
type SceneDeactivater interface {
	Scene
	Deactivate()
}

type SceneDestroyer interface {
	Scene
	Destroy() error
}

type SceneManager struct {
	list     map[string]Scene
	active   string
	schedule string
}

func NewSceneManager() *SceneManager {
	return &SceneManager{
		list: make(map[string]Scene, 3),
	}
}

func (sm *SceneManager) ActiveSceneName() string { return sm.active }

func (sm *SceneManager) ActivationScheduled() bool { return sm.schedule != "" }

func (sm *SceneManager) Get(name string) Scene {
	scene, _ := sm.list[name]
	return scene
}

func (sm *SceneManager) Has(name string) bool {
	scene, exists := sm.list[name]
	return exists && scene != nil
}

func (sm *SceneManager) Add(name string, scene Scene, activate bool) {
	sm.list[name] = scene

	if activate {
		sm.activate(name, scene)
	}
}

func (sm *SceneManager) activate(name string, scene Scene) {
	if sm.active != "" {
		if s, ok := sm.list[sm.active].(SceneDeactivater); ok {
			s.Deactivate()
		}
	}

	sm.active = name
	if a, ok := scene.(SceneActivater); ok {
		a.Activate()
	}
}

func (sm *SceneManager) Activate(name string) (Scene, error) {
	scene, exists := sm.list[name]
	if !exists {
		return nil, errors.Newf("sdlkit.SceneManager: scene %s does not exist", name)
	}

	sm.activate(name, scene)
	return scene, nil
}

func (sm *SceneManager) ScheduleActivation(name string) error {
	if !sm.Has(name) {
		return errors.Newf("sdlkit.SceneManager: scene %s does not exist", name)
	}

	sm.schedule = name
	return nil
}

func (sm *SceneManager) UpdateActiveScene(scenePtr *Scene) bool {
	if sm.schedule == "" {
		return false
	}

	scene, err := sm.Activate(sm.schedule)
	if err != nil {
		return false
	}

	sm.schedule = ""
	*scenePtr = scene
	return true
}

func (sm *SceneManager) Remove(name string, destroy bool) (bool, error) {
	scene, exists := sm.list[name]
	if !exists {
		return false, nil
	}

	var err error
	if destroy {
		if d, ok := scene.(SceneDestroyer); ok {
			err = d.Destroy()
		}
	}

	sm.list[name] = nil
	return true, err
}

func (sm *SceneManager) Destroy() error {
	var err error
	for _, scene := range sm.list {
		if d, ok := scene.(SceneDestroyer); ok {
			errors.Append(&err, d.Destroy())
		}
	}
	return err
}
