// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdlkit

import (
	"context"
	"image/color"

	"github.com/go-pogo/errors"
	sdlimg "github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

var DefaultStageOpts = StageOpts{
	// window options
	PosX:        sdl.WINDOWPOS_UNDEFINED,
	PosY:        sdl.WINDOWPOS_UNDEFINED,
	WindowFlags: sdl.WINDOW_SHOWN | sdl.WINDOW_OPENGL,

	// renderer options
	RendererIndex: -1,
	RendererFlags: sdl.RENDERER_ACCELERATED,

	// timer options
	TargetFps: DefaultFps,
}

type StageOpts struct {
	Context context.Context

	// sdl.Window options
	// see https://wiki.libsdl.org/SDL_CreateWindow
	PosX, PosY  int32
	WindowFlags uint32

	// sdl.Renderer options
	// see https://wiki.libsdl.org/SDL_CreateRenderer
	RendererIndex int
	RendererFlags uint32

	Icon    []byte      // see https://wiki.libsdl.org/SDL_SetWindowIcon
	BgColor color.Color // see https://wiki.libsdl.org/SDL_RenderClear

	// timer
	TargetFps uint8
	LimitFps  bool
}

type Stage struct {
	ctx context.Context
	cfn context.CancelFunc

	window   *sdl.Window
	renderer *sdl.Renderer
	viewport *Viewport
	time     *Time
	clock    *Clock
	scenes   *SceneManager

	BgColor color.RGBA
}

// NewStage creates a new Stage by first creating a new sdl.Window and
// sdl.Renderer from the provided StageOpts.
func NewStage(title string, width, height int32, opts StageOpts) (*Stage, error) {
	window, err := sdl.CreateWindow(title, opts.PosX, opts.PosY, width, height, opts.WindowFlags)
	if err != nil {
		return nil, errors.Trace(err)
	}

	renderer, err := sdl.CreateRenderer(window, opts.RendererIndex, opts.RendererFlags)
	if err != nil {
		return nil, errors.Trace(err)
	}

	clock := NewClock()
	stage := &Stage{
		window:   window,
		renderer: renderer,
		viewport: &Viewport{W: width, H: height},
		time:     NewTime(opts.TargetFps, clock),
		clock:    clock,
		scenes:   NewSceneManager(),
	}

	if col, ok := opts.BgColor.(color.RGBA); ok {
		stage.BgColor = col
	} else {
		r, g, b, a := opts.BgColor.RGBA()
		stage.BgColor = color.RGBA{
			R: uint8(r),
			G: uint8(g),
			B: uint8(b),
			A: uint8(a),
		}
	}

	if opts.Context == nil {
		opts.Context = context.Background()
	}

	stage.ctx, stage.cfn = context.WithCancel(opts.Context)
	stage.time.LimitFps = opts.LimitFps
	// stage.renderer.SetLogicalSize(width, height)

	if opts.Icon != nil {
		err = stage.setIcon(opts.Icon)
	}

	return stage, err
}

// MustNewStage creates a new Stage using NewStage and returns it on success.
// Any returned errors from NewStage are passed to FailOnErr and result in a
// fatal exit of the program.
func MustNewStage(title string, width, height int32, opts StageOpts) *Stage {
	stage, err := NewStage(title, width, height, opts)
	if err != nil {
		FailOnErr(err)
	}
	return stage
}

func (s *Stage) setIcon(icon []byte) error {
	src, err := sdl.RWFromMem(icon)
	if err != nil {
		return err
	}

	surface, err := sdlimg.LoadRW(src, true)
	if err != nil {
		return err
	}

	s.window.SetIcon(surface)
	return nil
}

func (s *Stage) Context() context.Context { return s.ctx }

// Window returns the sdl.Window in which the stage is set.
func (s *Stage) Window() *sdl.Window { return s.window }

// Renderer returns the sdl.Renderer that's attached to the window.
func (s *Stage) Renderer() *sdl.Renderer { return s.renderer }

// Viewport returns the current visible area within the window.
func (s *Stage) Viewport() *Viewport { return s.viewport }

// Time returns the Time that keeps track of time and framerate.
func (s *Stage) Time() *Time { return s.time }

func (s *Stage) Clock() *Clock { return s.clock }

func (s *Stage) SceneManager() *SceneManager { return s.scenes }

// Scene returns the current active scene from the SceneManager.
func (s *Stage) Scene() Scene {
	return s.scenes.Get(s.scenes.ActiveSceneName())
}

func (s *Stage) AddScene(name string, scene Scene) {
	s.scenes.Add(name, scene, s.scenes.ActiveSceneName() == "")
}

func (s *Stage) ClearScreen() error {
	var err error
	errors.Append(&err,
		s.renderer.SetDrawColor(
			s.BgColor.R,
			s.BgColor.G,
			s.BgColor.B,
			0xFF,
		),
		s.renderer.Clear(),
	)
	return err
}

func (s *Stage) PresentScreen() { s.renderer.Present() }

func (s *Stage) HandleWindowEvent(e *sdl.WindowEvent) error {
	switch e.Event {
	case sdl.WINDOWEVENT_SIZE_CHANGED:
		s.viewport.W = e.Data1
		s.viewport.H = e.Data2
	}

	return nil
}

func (s *Stage) Destroy() {
	s.cfn()

	if s.scenes != nil {
		s.scenes.Destroy()
	}

	_ = s.renderer.Destroy()
	_ = s.window.Destroy()
}
