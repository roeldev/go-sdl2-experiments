// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run -tags=static ./internal/gen.go

// Package event defines interfaces for sdl.Event handling.
package event

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/go-pogo/sdlkit"
)

type Manager struct {
	h handlers
}

func (m *Manager) Register(handler ...interface{}) { m.h.register(handler...) }

func (m *Manager) Handle(event sdl.Event) (err error) {
	if _, ok := event.(*sdl.QuitEvent); ok {
		return sdlkit.QUIT
	}

	return m.h.handle(event)
}

func (m *Manager) Loop() error {
	var event sdl.Event
	for {
		event = sdl.PollEvent()
		if event == nil {
			return nil
		}

		if err := m.Handle(event); err != nil {
			return err
		}
	}
}
