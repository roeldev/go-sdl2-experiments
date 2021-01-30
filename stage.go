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
	TargetFps: DefaultFps,
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
	TargetFps uint8 // todo: DisplayMode.RefreshRate
	LimitFps  bool
}

type Stage struct {
	BgColor        color.RGBA
	WindowTitleFps bool

	window   *sdl.Window
	renderer *sdl.Renderer
	scenes   *SceneManager
	time     *Time
	clock    *Clock

	ctx context.Context
	cfn context.CancelFunc

	initSize [2]int32
	prevSize [2]int32
	sizeRect sdl.Rect
	size     [2]float64
	fsMode   uint32
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
		scenes:   NewSceneManager(),
		time:     NewTime(opts.TargetFps, clock),
		clock:    clock,

		initSize: [2]int32{w, h},
		fsMode:   opts.FullscreenMode,
	}

	if err = stage.updateSize(w, h); err != nil {
		return nil, err
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

// Size returns the current logical size of the Stage.
func (s *Stage) Size() sdl.Rect { return s.sizeRect }

// Width returns the width of the logical size of the Stage.
func (s *Stage) Width() int32 { return s.sizeRect.W }

// Height returns the height of the logical size of the Stage.
func (s *Stage) Height() int32 { return s.sizeRect.H }

// Width returns the width of the logical size of the Stage as a float64.
func (s *Stage) FWidth() float64 { return s.size[0] }

// Height returns the height of the logical size of the Stage as a float64.
func (s *Stage) FHeight() float64 { return s.size[1] }

// Window returns the sdl.Window in which the Stage is set.
func (s *Stage) Window() *sdl.Window { return s.window }

// Renderer returns the sdl.Renderer that's attached to the window.
func (s *Stage) Renderer() *sdl.Renderer { return s.renderer }

// Time returns the Time that keeps track of time and framerate.
func (s *Stage) Time() *Time { return s.time }

func (s *Stage) Clock() *Clock { return s.clock }

// SceneManager returns the SceneManager instance that handles switching of
// scenes for the Stage.
func (s *Stage) SceneManager() *SceneManager { return s.scenes }

// Scene returns the current active scene from the SceneManager.
func (s *Stage) Scene() Scene {
	return s.scenes.Get(s.scenes.ActiveSceneName())
}

// AddScene adds a new Scene to the SceneManager of the Stage. It also activates
// the Scene when no other Scene is currently active.
func (s *Stage) AddScene(scene Scene) error {
	s.scenes.Add(scene)
	if s.scenes.ActiveSceneName() != "" {
		return nil
	}

	_, err := s.scenes.Activate(scene.SceneName())
	return err
}

// MustAddScene adds a Scene to the SceneManager, the same way AddScene does.
// Any errors are passed to FailOnErr.
func (s *Stage) MustAddScene(scene Scene, possibleErr error) {
	if possibleErr == nil {
		possibleErr = s.AddScene(scene)
	}
	FailOnErr(possibleErr)
}

func (s *Stage) HandleKeyUpEvent(e *sdl.KeyboardEvent) error {
	if e.Keysym.Scancode == sdl.SCANCODE_F11 {
		return s.ToggleFullscreen()
	}
	return nil
}

func (s *Stage) HandleWindowSizeChangedEvent(e *sdl.WindowEvent) error {
	return s.updateSize(e.Data1, e.Data2)
}

func (s *Stage) updateSize(w, h int32) error {
	if w == s.initSize[0] && h == s.initSize[1] {
		s.sizeRect.W = s.initSize[0]
		s.sizeRect.H = s.initSize[1]
	} else {
		s.sizeRect.W = s.initSize[0]
		s.sizeRect.H = int32(float32(h) / (float32(w) / float32(s.initSize[0])))
		if s.sizeRect.H < s.initSize[1] {
			s.sizeRect.W = int32(float32(w) / (float32(h) / float32(s.initSize[1])))
			s.sizeRect.H = s.initSize[1]
		}
	}

	s.size[0], s.size[1] = float64(s.sizeRect.W), float64(s.sizeRect.H)

	return s.renderer.SetLogicalSize(s.sizeRect.W, s.sizeRect.H)
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

func (s *Stage) Destroy() {
	s.cfn()

	// todo: send errors to log/stderr
	if s.scenes != nil {
		_ = s.scenes.Destroy()
	}

	_ = s.renderer.Destroy()
	_ = s.window.Destroy()
}
