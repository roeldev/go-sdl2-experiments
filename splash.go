// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdlkit

import (
	"github.com/go-pogo/errors"
	sdlimg "github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type SplashScreen struct {
	window *sdl.Window
}

func NewSplashScreen(w, h int32) (*SplashScreen, error) {
	window, err := sdl.CreateWindow("",
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		w, h,
		sdl.WINDOW_BORDERLESS|sdl.WINDOW_ALWAYS_ON_TOP|sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return &SplashScreen{window: window}, nil
}

func MustNewSplashScreen(w, h int32) *SplashScreen {
	splash, err := NewSplashScreen(w, h)
	if err != nil {
		FailOnErr(err)
	}
	return splash
}

// Window returns the sdl.Window where the SplashScreen is shown in.
func (s *SplashScreen) Window() *sdl.Window { return s.window }

func (s *SplashScreen) DisplayImage(data []byte) error {
	src, err := sdl.RWFromMem(data)
	if err != nil {
		return err
	}

	img, err := sdlimg.LoadRW(src, true)
	if err != nil {
		return err
	}

	win, err := s.window.GetSurface()
	if err != nil {
		return err
	}

	if err = win.FillRect(&win.ClipRect, sdl.MapRGB(win.Format, 0, 0, 0)); err != nil {
		return err
	}
	if err = img.Blit(nil, win, nil); err != nil {
		return err
	}

	return s.window.UpdateSurface()
}

func (s *SplashScreen) Destroy() error {
	return s.window.Destroy()
}
