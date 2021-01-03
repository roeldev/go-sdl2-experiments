// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package geom

import (
	"fmt"
	"math"
)

const (
	r2dPi = 180 / math.Pi
	d2rPi = math.Pi / 180
)

var (
	vectorUp    = Vector{0, -1}
	vectorDown  = Vector{0, 1}
	vectorLeft  = Vector{-1, 0}
	vectorRight = Vector{1, 0}
)

type Radians float64

func (r Radians) Degrees() float64 {
	return float64(r) * r2dPi
}

type Degrees float64

func (d Degrees) Radians() float64 {
	return float64(d) * d2rPi
}

type Vector struct {
	X float32
	Y float32
}

// Angle returns the angle in radians.
func (v *Vector) Angle() Radians {
	return Radians(math.Atan2(float64(v.Y), float64(v.X)))
}

// Length returns the length (magnitude) of the Vector.
func (v *Vector) Length() float64 {
	return math.Sqrt(v.LengthSq())
}

func (v *Vector) LengthSq() float64 {
	return float64((v.X * v.X) + (v.Y * v.Y))
}

func (v *Vector) Distance(vec Vector) float64 {
	return math.Sqrt(v.DistanceSq(vec))
}

func (v *Vector) DistanceSq(vec Vector) float64 {
	dx := vec.X - v.X
	dy := vec.Y - v.Y
	return float64((dx * dx) + (dy * dy))
}

// Equals compares the X and Y values of this Vector and the given Vector, and
// returns true when they are equal.
func (v *Vector) Equals(vec Vector) bool {
	return v.X == vec.X && v.Y == vec.Y
}

// Dot calculates the dot product of this Vector and the given Vector.
func (v *Vector) Dot(vec Vector) float32 {
	return (v.X * vec.X) + (v.Y * vec.Y)
}

// Cross calculates the cross product of this Vector and the given Vector.
func (v *Vector) Cross(vec Vector) float32 {
	return (v.X * vec.Y) + (v.Y * vec.X)
}

// set the X and Y values according to the provided angle in radians and length.
func (v *Vector) set(angle Radians, length float64) *Vector {
	v.X = float32(math.Cos(float64(angle)) * length)
	v.Y = float32(math.Sin(float64(angle)) * length)
	return v
}

func (v *Vector) SetAngle(angle Radians) *Vector { return v.set(angle, v.Length()) }

func (v *Vector) SetLength(l float64) *Vector { return v.set(v.Angle(), l) }

// Limit the length of this Vector.
func (v *Vector) Limit(length float64) *Vector {
	if v.Length() <= length {
		return v
	}

	return v.SetLength(length)
}

// Rotate this Vector by an angle amount.
func (v *Vector) Rotate(angle Radians) *Vector {
	cos := math.Cos(float64(angle))
	sin := math.Sin(float64(angle))

	x := float64(v.X)
	y := float64(v.Y)

	v.X = float32((cos * x) - (sin * y))
	v.Y = float32((sin * x) - (cos * y))
	return v
}

// Lerp linearly interpolates this Vector towards the given Vector. Value t is
// the interpolation percentage between 0 and 1.
func (v *Vector) Lerp(vec Vector, t float32) *Vector {
	v.X += (vec.X - v.X) * t
	v.Y += (vec.Y - v.Y) * t
	return v
}

// Scale the Vector by the given scale value, where 1 is equal to the Vector's
// current value.
func (v *Vector) Scale(scale float32) *Vector {
	v.X *= scale
	v.Y *= scale
	return v
}

func (v *Vector) String() string {
	angle := v.Angle()
	return fmt.Sprintf("%T{%f, %f}, angle: %f (= %f°), length: %f", v, v.X, v.Y, angle, angle*r2dPi, v.Length())
}
