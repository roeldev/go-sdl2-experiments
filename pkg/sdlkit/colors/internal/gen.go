// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"log"
	"path"
	"runtime"
	"text/template"

	"github.com/veandco/go-sdl2/sdl"
)

// https://simple.wikipedia.org/wiki/Web_color
var list = map[string]map[string]sdl.Color{
	"Blue": {
		"Aqua":            sdl.Color{G: 0xFF, B: 0xFF},
		"Cyan":            sdl.Color{G: 0xFF, B: 0xFF},
		"LightCyan":       sdl.Color{R: 0xE0, G: 0xFF, B: 0xFF},
		"PaleTurquoise":   sdl.Color{R: 0xAF, G: 0xEE, B: 0xEE},
		"Aquamarine":      sdl.Color{R: 0x7F, G: 0xFF, B: 0xD4},
		"Turquoise":       sdl.Color{R: 0x40, G: 0xE0, B: 0xD0},
		"MediumTurquoise": sdl.Color{R: 0x48, G: 0xD1, B: 0xCC},
		"DarkTurquoise":   sdl.Color{G: 0xCE, B: 0xD1},
		"CadetBlue":       sdl.Color{R: 0x5F, G: 0x9E, B: 0xA0},
		"SteelBlue":       sdl.Color{R: 0x46, G: 0x82, B: 0xB4},
		"LightSteelBlue":  sdl.Color{R: 0xB0, G: 0xC4, B: 0xDE},
		"PowderBlue":      sdl.Color{R: 0xB0, G: 0xE0, B: 0xE6},
		"LightBlue":       sdl.Color{R: 0xAD, G: 0xD8, B: 0xE6},
		"SkyBlue":         sdl.Color{R: 0x87, G: 0xCE, B: 0xEB},
		"LightSkyBlue":    sdl.Color{R: 0x87, G: 0xCE, B: 0xFA},
		"DeepSkyBlue":     sdl.Color{G: 0xBF, B: 0xFF},
		"DodgerBlue":      sdl.Color{R: 0x1E, G: 0x90, B: 0xFF},
		"CornflowerBlue":  sdl.Color{R: 0x64, G: 0x95, B: 0xED},
		"MediumSlateBlue": sdl.Color{R: 0x7B, G: 0x68, B: 0xEE},
		"RoyalBlue":       sdl.Color{R: 0x41, G: 0x69, B: 0xE1},
		"Blue":            sdl.Color{B: 0xFF},
		"MediumBlue":      sdl.Color{B: 0xCD},
		"DarkBlue":        sdl.Color{B: 0x8B},
		"Navy":            sdl.Color{B: 0x80},
		"MidnightBlue":    sdl.Color{R: 0x19, G: 0x19, B: 0x70},
	},
	"Brown": {
		"Cornsilk":       sdl.Color{R: 0xFF, G: 0xF8, B: 0xDC},
		"BlanchedAlmond": sdl.Color{R: 0xFF, G: 0xEB, B: 0xCD},
		"Bisque":         sdl.Color{R: 0xFF, G: 0xE4, B: 0xC4},
		"NavajoWhite":    sdl.Color{R: 0xFF, G: 0xDE, B: 0xAD},
		"Wheat":          sdl.Color{R: 0xF5, G: 0xDE, B: 0xB3},
		"BurlyWood":      sdl.Color{R: 0xDE, G: 0xB8, B: 0x87},
		"Tan":            sdl.Color{R: 0xD2, G: 0xB4, B: 0x8C},
		"RosyBrown":      sdl.Color{R: 0xBC, G: 0x8F, B: 0x8F},
		"SandyBrown":     sdl.Color{R: 0xF4, G: 0xA4, B: 0x60},
		"Goldenrod":      sdl.Color{R: 0xDA, G: 0xA5, B: 0x20},
		"DarkGoldenrod":  sdl.Color{R: 0xB8, G: 0x86, B: 0x0B},
		"Peru":           sdl.Color{R: 0xCD, G: 0x85, B: 0x3F},
		"Chocolate":      sdl.Color{R: 0xD2, G: 0x69, B: 0x1E},
		"SaddleBrown":    sdl.Color{R: 0x8B, G: 0x45, B: 0x13},
		"Sienna":         sdl.Color{R: 0xA0, G: 0x52, B: 0x2D},
		"Brown":          sdl.Color{R: 0xA5, G: 0x2A, B: 0x2A},
		"Maroon":         sdl.Color{R: 0x80},
	},
	"Green": {
		"GreenYellow":       sdl.Color{R: 0xAD, G: 0xFF, B: 0x2F},
		"Chartreuse":        sdl.Color{R: 0x7F, G: 0xFF},
		"LawnGreen":         sdl.Color{R: 0x7C, G: 0xFC},
		"Lime":              sdl.Color{G: 0xFF},
		"LimeGreen":         sdl.Color{R: 0x32, G: 0xCD, B: 0x32},
		"PaleGreen":         sdl.Color{R: 0x98, G: 0xFB, B: 0x98},
		"LightGreen":        sdl.Color{R: 0x90, G: 0xEE, B: 0x90},
		"MediumSpringGreen": sdl.Color{G: 0xFA, B: 0x9A},
		"SpringGreen":       sdl.Color{G: 0xFF, B: 0x7F},
		"MediumSeaGreen":    sdl.Color{R: 0x3C, G: 0xB3, B: 0x71},
		"SeaGreen":          sdl.Color{R: 0x2E, G: 0x8B, B: 0x57},
		"ForestGreen":       sdl.Color{R: 0x22, G: 0x8B, B: 0x22},
		"Green":             sdl.Color{G: 0x80},
		"DarkGreen":         sdl.Color{G: 0x64},
		"YellowGreen":       sdl.Color{R: 0x9A, G: 0xCD, B: 0x32},
		"OliveDrab":         sdl.Color{R: 0x6B, G: 0x8E, B: 0x23},
		"Olive":             sdl.Color{R: 0x80, G: 0x80},
		"DarkOliveGreen":    sdl.Color{R: 0x55, G: 0x6B, B: 0x2F},
		"MediumAquamarine":  sdl.Color{R: 0x66, G: 0xCD, B: 0xAA},
		"DarkSeaGreen":      sdl.Color{R: 0x8F, G: 0xBC, B: 0x8F},
		"LightSeaGreen":     sdl.Color{R: 0x20, G: 0xB2, B: 0xAA},
		"DarkCyan":          sdl.Color{G: 0x8B, B: 0x8B},
		"Teal":              sdl.Color{G: 0x80, B: 0x80},
	},
	"Grey": {
		"Gainsboro":      sdl.Color{R: 0xDC, G: 0xDC, B: 0xDC},
		"LightGrey":      sdl.Color{R: 0xD3, G: 0xD3, B: 0xD3},
		"Silver":         sdl.Color{R: 0xC0, G: 0xC0, B: 0xC0},
		"DarkGray":       sdl.Color{R: 0xA9, G: 0xA9, B: 0xA9},
		"Gray":           sdl.Color{R: 0x80, G: 0x80, B: 0x80},
		"DimGray":        sdl.Color{R: 0x69, G: 0x69, B: 0x69},
		"LightSlateGray": sdl.Color{R: 0x77, G: 0x88, B: 0x99},
		"SlateGray":      sdl.Color{R: 0x70, G: 0x80, B: 0x90},
		"DarkSlateGray":  sdl.Color{R: 0x2F, G: 0x4F, B: 0x4F},
		"Black":          sdl.Color{},
	},
	"Orange": {
		"Coral":      sdl.Color{R: 0xFF, G: 0x7F, B: 0x50},
		"Tomato":     sdl.Color{R: 0xFF, G: 0x63, B: 0x47},
		"OrangeRed":  sdl.Color{R: 0xFF, G: 0x45},
		"DarkOrange": sdl.Color{R: 0xFF, G: 0x8C},
		"Orange":     sdl.Color{R: 0xFF, G: 0xA5},
	},
	"Pink": {
		"Pink":            sdl.Color{R: 0xFF, G: 0xC0, B: 0xCB},
		"LightPink":       sdl.Color{R: 0xFF, G: 0xB6, B: 0xC1},
		"HotPink":         sdl.Color{R: 0xFF, G: 0x69, B: 0xB4},
		"DeepPink":        sdl.Color{R: 0xFF, G: 0x14, B: 0x93},
		"MediumVioletRed": sdl.Color{R: 0xC7, G: 0x15, B: 0x85},
		"PaleVioletRed":   sdl.Color{R: 0xDB, G: 0x70, B: 0x93},
	},
	"Purple": {
		"Lavender":        sdl.Color{R: 0xE6, G: 0xE6, B: 0xFA},
		"Thistle":         sdl.Color{R: 0xD8, G: 0xBF, B: 0xD8},
		"Plum":            sdl.Color{R: 0xDD, G: 0xA0, B: 0xDD},
		"Violet":          sdl.Color{R: 0xEE, G: 0x82, B: 0xEE},
		"Orchid":          sdl.Color{R: 0xDA, G: 0x70, B: 0xD6},
		"Fuchsia":         sdl.Color{R: 0xFF, B: 0xFF},
		"Magenta":         sdl.Color{R: 0xFF, B: 0xFF},
		"MediumOrchid":    sdl.Color{R: 0xBA, G: 0x55, B: 0xD3},
		"MediumPurple":    sdl.Color{R: 0x93, G: 0x70, B: 0xDB},
		"Amethyst":        sdl.Color{R: 0x99, G: 0x66, B: 0xCC},
		"BlueViolet":      sdl.Color{R: 0x8A, G: 0x2B, B: 0xE2},
		"DarkViolet":      sdl.Color{R: 0x94, B: 0xD3},
		"DarkOrchid":      sdl.Color{R: 0x99, G: 0x32, B: 0xCC},
		"DarkMagenta":     sdl.Color{R: 0x8B, B: 0x8B},
		"Purple":          sdl.Color{R: 0x80, B: 0x80},
		"Indigo":          sdl.Color{R: 0x4B, B: 0x82},
		"SlateBlue":       sdl.Color{R: 0x6A, G: 0x5A, B: 0xCD},
		"DarkSlateBlue":   sdl.Color{R: 0x48, G: 0x3D, B: 0x8B},
		"MediumSlateBlue": sdl.Color{R: 0x7B, G: 0x68, B: 0xEE},
	},
	"Red": {
		"IndianRed":   sdl.Color{R: 0xCD, G: 0x5C, B: 0x5C},
		"LightCoral":  sdl.Color{R: 0xF0, G: 0x80, B: 0x80},
		"Salmon":      sdl.Color{R: 0xFA, G: 0x80, B: 0x72},
		"DarkSalmon":  sdl.Color{R: 0xE9, G: 0x96, B: 0x7A},
		"LightSalmon": sdl.Color{R: 0xFF, G: 0xA0, B: 0x7A},
		"Crimson":     sdl.Color{R: 0xDC, G: 0x14, B: 0x3C},
		"Red":         sdl.Color{R: 0xFF},
		"FireBrick":   sdl.Color{R: 0xB2, G: 0x22, B: 0x22},
		"DarkRed":     sdl.Color{R: 0x8B},
	},
	"White": {
		"White":         sdl.Color{R: 0xFF, G: 0xFF, B: 0xFF},
		"Snow":          sdl.Color{R: 0xFF, G: 0xFA, B: 0xFA},
		"Honeydew":      sdl.Color{R: 0xF0, G: 0xFF, B: 0xF0},
		"MintCream":     sdl.Color{R: 0xF5, G: 0xFF, B: 0xFA},
		"Azure":         sdl.Color{R: 0xF0, G: 0xFF, B: 0xFF},
		"AliceBlue":     sdl.Color{R: 0xF0, G: 0xF8, B: 0xFF},
		"GhostWhite":    sdl.Color{R: 0xF8, G: 0xF8, B: 0xFF},
		"WhiteSmoke":    sdl.Color{R: 0xF5, G: 0xF5, B: 0xF5},
		"Seashell":      sdl.Color{R: 0xFF, G: 0xF5, B: 0xEE},
		"Beige":         sdl.Color{R: 0xF5, G: 0xF5, B: 0xDC},
		"OldLace":       sdl.Color{R: 0xFD, G: 0xF5, B: 0xE6},
		"FloralWhite":   sdl.Color{R: 0xFF, G: 0xFA, B: 0xF0},
		"Ivory":         sdl.Color{R: 0xFF, G: 0xFF, B: 0xF0},
		"AntiqueWhite":  sdl.Color{R: 0xFA, G: 0xEB, B: 0xD7},
		"Linen":         sdl.Color{R: 0xFA, G: 0xF0, B: 0xE6},
		"LavenderBlush": sdl.Color{R: 0xFF, G: 0xF0, B: 0xF5},
		"MistyRose":     sdl.Color{R: 0xFF, G: 0xE4, B: 0xE1},
	},
	"Yellow": {
		"Gold":                 sdl.Color{R: 0xFF, G: 0xD7},
		"Yellow":               sdl.Color{R: 0xFF, G: 0xFF},
		"LightYellow":          sdl.Color{R: 0xFF, G: 0xFF, B: 0xE0},
		"LemonChiffon":         sdl.Color{R: 0xFF, G: 0xFA, B: 0xCD},
		"LightGoldenrodYellow": sdl.Color{R: 0xFA, G: 0xFA, B: 0xD2},
		"PapayaWhip":           sdl.Color{R: 0xFF, G: 0xEF, B: 0xD5},
		"Moccasin":             sdl.Color{R: 0xFF, G: 0xE4, B: 0xB5},
		"PeachPuff":            sdl.Color{R: 0xFF, G: 0xDA, B: 0xB9},
		"PaleGoldenrod":        sdl.Color{R: 0xEE, G: 0xE8, B: 0xAA},
		"Khaki":                sdl.Color{R: 0xF0, G: 0xE6, B: 0x8C},
		"DarkKhaki":            sdl.Color{R: 0xBD, G: 0xB7, B: 0x6B},
	},
}

const all = "All"

type vars struct {
	Vars   map[string]sdl.Color
	Groups map[string]map[sdl.Color]string
}

func main() {
	_, dir, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("unable to resolve project dir location")
	}

	dir = path.Dir(path.Dir(dir))

	v := vars{
		Vars: make(map[string]sdl.Color, 150),
		Groups: map[string]map[sdl.Color]string{
			all: make(map[sdl.Color]string, 150),
		},
	}

	tmpl := template.Must(template.New("").Parse(groupFileSrc))

	for group, colors := range list {
		v.Groups[group] = make(map[sdl.Color]string, 20)

		for name, col := range colors {
			col.A = 0xFF
			v.Vars[name] = col
			v.Groups[group][col] = name
			v.Groups[all][col] = name
		}
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, v); err != nil {
		log.Fatal("error while executing template:", err)
	}

	data, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal("error while formatting code:", err)
	}

	filename := path.Join(dir, "colors.go")
	if err = ioutil.WriteFile(filename, data, 0644); err != nil {
		log.Fatal("error writing", filename+":", err)
	}
}

const groupFileSrc = `// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// generated by go generate; DO NOT EDIT.

package colors

import "github.com/veandco/go-sdl2/sdl"

var (
	{{ range $name, $color := .Vars }}
	{{ $name }} = {{ printf "%#v" $color }}{{ end }}
)

//goland:noinspection GoUnusedGlobalVariable
var (
	{{ range $name, $colors := .Groups }}
	{{ $name }}Colors = []sdl.Color{ {{ range $color, $name := $colors }}
		{{ $name }},{{ end }}
	}
	{{ end }}

	{{ range $name, $colors := .Groups }}
	{{ $name }}Names = []string{ {{ range $color, $name := $colors }}
		"{{ $name }}",{{ end }}
	}
	{{ end }}
)`
