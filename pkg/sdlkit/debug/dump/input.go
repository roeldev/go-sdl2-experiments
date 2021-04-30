// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dump

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/debug/constlookup"
)

func Keysym(k sdl.Keysym) string {
	return dump(keysymDumper(k))
}

func keysymDumper(k sdl.Keysym) *dumper {
	d := newDumper(k)
	d.Const("Scancode", k.Scancode, constlookup.Scancodes.Lookup(k.Scancode), "unknown scancode")
	d.Const("Sym", k.Sym, constlookup.Keycodes.Lookup(k.Sym), "unknown keycode")
	d.Const("Mod", k.Mod, constlookup.Keymods.Lookup(k.Mod), "unknown keymod")
	return d
}
