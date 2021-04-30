// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dump

import (
	"fmt"
	"strings"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/debug/constlookup"
)

func DisplayMode(dm sdl.DisplayMode) string {
	return dump(displayModeDumper(dm))
}

func displayModeDumper(dm sdl.DisplayMode) *dumper {
	d := newDumper(dm)
	d.Const("Format", dm.Format, constlookup.PixelFormats.Lookup(dm.Format), "unknown pixel format")
	d.Number("W", dm.W)
	d.Number("H", dm.H)
	d.Number("RefreshRate", dm.RefreshRate)
	d.Number("DriverData", dm.DriverData)
	return d
}

func DisplayModes() string {
	vn, err := sdl.GetNumVideoDisplays()
	if err != nil {
		panic(err)
	}

	res := make([]string, 0, 20)
	for i := 0; i < vn; i++ {
		n, err := sdl.GetNumDisplayModes(i)
		if err != nil {
			panic(err)
		}

		for j := 0; j < n; j++ {
			dm, err := sdl.GetDisplayMode(i, j)
			if err != nil {
				continue
			}

			d := displayModeDumper(dm)
			d.Number("> VIDEO DISPLAY", i)
			d.Number("> DISPLAY MODE", j)
			d.Number("> RATIO", float32(dm.W)/float32(dm.H))
			res = append(res, dump(d))
		}
	}
	return strings.Join(res, "\n")
}

func WindowFlags(flags uint32) string {
	res := make([]string, 0, len(constlookup.WindowFlags))
	for flag, str := range constlookup.WindowFlags {
		if flags&flag != 0 {
			res = append(res, fmt.Sprintf("%s (%d)", str, flag))
		}
	}
	return strings.Join(res, "|")
}
