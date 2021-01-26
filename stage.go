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

var DefaultOptions = Options{
	// window options
	PosX:           sdl.WINDOWPOS_CENTERED,
	PosY:           sdl.WINDOWPOS_CENTERED,
	WindowFlags:    sdl.WINDOW_SHOWN | sdl.WINDOW_INPUT_FOCUS,
	FullscreenMode: 1,

	// renderer options
	RendererIndex: -1,
	RendererFlags: sdl.RENDERER_ACCELERATED,
	BgColor:       color.RGBA{},

	// timer options
	TargetFps:      DefaultFps,
	WindowTitleFps: true,
}

type Options struct {
	Context context.Context

	// sdl.Window options
	// see https://wiki.libsdl.org/SDL_CreateWindow
	PosX, PosY     int32
	WindowFlags    uint32
	FullscreenMode uint32 // https://wiki.libsdl.org/SDL_SetWindowFullscreen

	// sdl.DisplayMode options
	// (https://wiki.libsdl.org/SDL_DisplayMode)
	DisplayMode sdl.DisplayMode

	// sdl.Renderer options
	// see https://wiki.libsdl.org/SDL_CreateRenderer
	RendererIndex int
	RendererFlags uint32

	BgColor color.Color // see https://wiki.libsdl.org/SDL_RenderClear

	// timer
	TargetFps      uint8 // todo: DisplayMode.RefreshRate
	LimitFps       bool
	WindowTitleFps bool
}

type Stage struct {
	ctx context.Context
	cfn context.CancelFunc

	minW, minH int32
	fsMode     uint32

	opts     Options
	window   *sdl.Window
	renderer *sdl.Renderer
	viewport *Viewport
	time     *Time
	clock    *Clock
	scenes   *SceneManager

	BgColor        color.RGBA
	WindowTitleFps bool
}

// NewStage creates a new Stage by first creating a new sdl.Window and
// sdl.Renderer. These are configured with the provided Options.
func NewStage(title string, w, h int32, opts Options) (*Stage, error) {
	window, err := sdl.CreateWindow(title, opts.PosX, opts.PosY, w, h, opts.WindowFlags)
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
		viewport: &Viewport{W: w, H: h},
		time:     NewTime(opts.TargetFps, clock),
		clock:    clock,
		scenes:   NewSceneManager(),

		minW:   w,
		minH:   h,
		fsMode: opts.FullscreenMode,

		WindowTitleFps: opts.WindowTitleFps,
	}

	index, err := window.GetDisplayIndex()
	if err == nil {
		var dm *sdl.DisplayMode
		opts.DisplayMode.W = w
		opts.DisplayMode.H = h

		dm, err = GetClosestDisplayModeRatio(index, opts.DisplayMode)
		if err == nil {
			_ = window.SetDisplayMode(dm)
		}
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

	return stage, err
}

// MustNewStage creates a new Stage using NewStage and returns it on success.
// Any returned errors from NewStage are passed to FailOnErr and result in a
// fatal exit of the program.
func MustNewStage(title string, width, height int32, opts Options) *Stage {
	stage, err := NewStage(title, width, height, opts)
	if err != nil {
		FailOnErr(err)
	}
	return stage
}

// https://wiki.libsdl.org/SDL_SetWindowIcon
func (s *Stage) SetWindowIcon(icon []byte) error {
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
		s.renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND),
	)
	return err
}

func (s *Stage) PresentScreen() { s.renderer.Present() }

func (s *Stage) ToggleFullscreen() (err error) {
	if s.window.GetFlags()&s.fsMode != 0 {
		return s.window.SetFullscreen(0)
	}

	return s.window.SetFullscreen(s.fsMode)
}

func (s *Stage) HandleKeyUpEvent(e *sdl.KeyboardEvent) error {
	if e.Keysym.Scancode == sdl.SCANCODE_F11 {
		return s.ToggleFullscreen()
	}
	return nil
}

func (s *Stage) HandleWindowSizeChangedEvent(e *sdl.WindowEvent) error {
	if e.Data1 == s.minW && e.Data2 == s.minH {
		s.viewport.W = s.minW
		s.viewport.H = s.minH
	} else {
		s.viewport.W = s.minW
		s.viewport.H = int32(float32(e.Data2) / (float32(e.Data1) / float32(s.minW)))
		if s.viewport.H < s.minH {
			s.viewport.W = int32(float32(e.Data1) / (float32(e.Data2) / float32(s.minH)))
			s.viewport.H = s.minH
		}
	}

	return s.renderer.SetLogicalSize(s.viewport.W, s.viewport.H)
}

func (s *Stage) Destroy() {
	s.cfn()

	// todo: send errors to log/stderr
	if s.scenes != nil {
		_ = s.scenes.Destroy()
	}

	_ = s.renderer.Destroy()
	_ = s.window.Destroy()
}
