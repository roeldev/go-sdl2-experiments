// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package geom

type XYGetter interface {
	GetX() float64
	GetY() float64
}

type XYSetter interface {
	SetX(x float64)
	SetY(y float64)
}

type XY interface {
	XYGetter
	XYSetter
}
