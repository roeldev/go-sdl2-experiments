// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run ./internal/gen.go

package colors

import (
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
)

func RandColor(r *rand.Rand) sdl.Color { return AllColors[r.Intn(len(AllColors))] }

func RandBlue(r *rand.Rand) sdl.Color { return BlueColors[r.Intn(len(BlueColors))] }

func RandBrown(r *rand.Rand) sdl.Color { return BrownColors[r.Intn(len(BrownColors))] }

func RandGreen(r *rand.Rand) sdl.Color { return GreenColors[r.Intn(len(GreenColors))] }

func RandGrey(r *rand.Rand) sdl.Color { return GreyColors[r.Intn(len(GreyColors))] }

func RandOrange(r *rand.Rand) sdl.Color { return OrangeColors[r.Intn(len(OrangeColors))] }

func RandPink(r *rand.Rand) sdl.Color { return PinkColors[r.Intn(len(PinkColors))] }

func RandPurple(r *rand.Rand) sdl.Color { return PurpleColors[r.Intn(len(PurpleColors))] }

func RandRed(r *rand.Rand) sdl.Color { return RedColors[r.Intn(len(RedColors))] }

func RandWhite(r *rand.Rand) sdl.Color { return WhiteColors[r.Intn(len(WhiteColors))] }

func RandYellow(r *rand.Rand) sdl.Color { return YellowColors[r.Intn(len(YellowColors))] }
