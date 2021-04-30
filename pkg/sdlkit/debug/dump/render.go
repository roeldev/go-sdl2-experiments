// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dump

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/debug/constlookup"
)

func TextureQuery(tx *sdl.Texture) string {
	if tx == nil {
		return "nil"
	}

	format, access, width, height, err := tx.Query()

	d := newDumper(nil)
	d.Const("format", format, constlookup.PixelFormats.Lookup(format), "unknown format")
	d.Number("access", access)
	d.Number("width", width)
	d.Number("height", height)

	if err != nil {
		d.String("err", err.Error())
	}
	return dump(d)
}
