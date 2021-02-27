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

type Provider interface {
	RegisterEvents(em *Manager)
}

type Manager struct {
	h handlers
}

func (m *Manager) RegisterHandler(handler ...interface{}) { m.register(handler, false) }

func (m *Manager) MustRegisterHandler(handler ...interface{}) { m.register(handler, true) }

func (m *Manager) register(handlers []interface{}, mustRegister bool) {
	for _, handler := range handlers {
		if handler == nil {
			continue
		}

		if p, ok := handler.(Provider); ok {
			p.RegisterEvents(m)
		} else if m.h.register(handler) == 0 && mustRegister {
			// todo: panic using log
			panic(fmt.Sprintf("sdlkit event.Manager:\n\tcannot register `%T` as it does not have any event handlers methods that match\n\tmake sure the handler is of the correct type (eg. is a pointer to the type)", handler))
		}
	}
}

func (m *Manager) HandleEvent(event sdl.Event) (err error) {
	if _, ok := event.(*sdl.QuitEvent); ok {
		return sdlkit.QUIT
	}

	return m.h.handle(event)
}

func (m *Manager) Process() error {
	var event sdl.Event
	for {
		event = sdl.PollEvent()
		if event == nil {
			return nil
		}

		if err := m.HandleEvent(event); err != nil {
			return err
		}
	}
}
