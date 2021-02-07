// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run -tags=static ./internal/gen.go

// Package event defines interfaces for sdl.Event handling.
package event

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/go-pogo/sdlkit"
)

type Manager struct {
	h handlers
}

func (m *Manager) Register(handler ...interface{}) {
	for _, h := range handler {
		m.h.register(h)
	}
}

func (m *Manager) MustRegister(handler ...interface{}) {
	for _, h := range handler {
		if m.h.register(h) == 0 {
			// todo: panic using log
			panic(fmt.Sprintf("sdlkit event.Manager:\n\tcannot register `%T` as it does not have any event handlers methods that match\n\tmake sure the handler is of the correct type (eg. is a pointer to the type)", h))
		}
	}
}

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
