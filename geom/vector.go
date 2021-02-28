// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package geom

import (
	"fmt"
	"math"

	math2 "github.com/go-pogo/sdlkit/math"
)

// var (
// 	vectorUp    = Vector{0, -1}
// 	vectorDown  = Vector{0, 1}
// 	vectorLeft  = Vector{-1, 0}
// 	vectorRight = Vector{1, 0}
// )

type Vector struct {
	X, Y float64
}

// VectorFromInt creates a new Vector from int32 values.
func VectorFromInt(x, y int32) *Vector {
	return &Vector{X: float64(x), Y: float64(y)}
}

// VectorFromXY creates a new Vector from an XYGetter.
func VectorFromXY(xy XYGetter) *Vector {
	return &Vector{X: xy.GetX(), Y: xy.GetY()}
}

func (v Vector) GetX() float64 { return v.X }
func (v Vector) GetY() float64 { return v.Y }

func (v *Vector) SetX(x float64) { v.X = x }
func (v *Vector) SetY(y float64) { v.Y = y }

// Point returns a Point with the same X and Y values as the Vector.
func (v Vector) Point() Point { return Point{X: v.X, Y: v.Y} }

// Clone returns a pointer to a new Vector with the same X and Y values as the
// source Vector.
func (v Vector) Clone() *Vector { return &Vector{X: v.X, Y: v.Y} }

// SetAngle in radians.
func (v *Vector) SetAngle(angle float64) *Vector {
	return v.FromPolar(angle, v.Length())
}

func (v *Vector) SetLength(l float64) *Vector {
	return v.FromPolar(v.Angle(), l)
}

// FromPolar sets the X and Y values according to the provided angle (in
// radians) and length.
func (v *Vector) FromPolar(angle, length float64) *Vector {
	v.X = math.Cos(angle) * length
	v.Y = math.Sin(angle) * length
	return v
}

// Zero sets this Vector to 0 values.
func (v *Vector) Zero() *Vector {
	v.X, v.Y = 0, 0
	return v
}

// Negate the X and Y values of this Vector, meaning negative numbers becoming
// positive and positive becoming negative.
func (v *Vector) Negate() *Vector {
	v.X = -v.X
	v.Y = -v.Y
	return v
}

// Scale this Vector by the given scale value, where 1 is equal to the Vector's
// current value.
func (v *Vector) Scale(scale float64) *Vector {
	v.X *= scale
	v.Y *= scale
	return v
}

// Limit the length of this Vector.
func (v *Vector) Limit(length float64) *Vector {
	if v.Length() <= length {
		return v
	}
	return v.SetLength(length)
}

// Rotate this Vector by an angle amount (in radians).
func (v *Vector) Rotate(angle float64) *Vector {
	x, y := v.X, v.Y
	cos, sin := math.Cos(angle), math.Sin(angle)

	v.X = (x * cos) - (y * sin)
	v.Y = (y * cos) - (x * sin)
	return v
}

func (v *Vector) Add(x, y float64) *Vector {
	v.X += x
	v.Y += y
	return v
}

// AddVec adds the given Vector to this Vector.
func (v *Vector) AddVec(add Vector) *Vector {
	v.X += add.X
	v.Y += add.Y
	return v
}

// AddXY the given XYGetter values to this Vector.
func (v *Vector) AddXY(add XYGetter) *Vector {
	v.X += add.GetX()
	v.Y += add.GetY()
	return v
}

// SubVec subtracts the given Vector from this Vector.
func (v *Vector) SubVec(sub Vector) *Vector {
	v.X -= sub.X
	v.Y -= sub.Y
	return v
}

// SubXY subtracts the given XYGetter values from this Vector.
func (v *Vector) SubXY(sub XYGetter) *Vector {
	v.X -= sub.GetX()
	v.Y -= sub.GetY()
	return v
}

// MulVec multiplies this Vector by the given Vector.
func (v *Vector) MulVec(mul Vector) *Vector {
	v.X *= mul.X
	v.Y *= mul.Y
	return v
}

// Div divides this Vector by the given Vector.
func (v *Vector) Div(div Vector) *Vector {
	v.X /= div.X
	v.Y /= div.Y
	return v
}

// Lerp linearly interpolates this Vector towards the target Vector. Value t is
// the interpolation percentage between 0 and 1.
func (v *Vector) Lerp(target Vector, t float64) *Vector {
	v.X += (target.X - v.X) * t
	v.Y += (target.Y - v.Y) * t
	return v
}

func (v *Vector) LerpRound(target Vector, t, e float64) *Vector {
	v.X = math2.LerpRound(v.X, target.X, t, e)
	v.Y = math2.LerpRound(v.Y, target.Y, t, e)
	return v
}

// Angle returns the angle in radians.
func (v Vector) Angle() float64 { return math.Atan2(v.Y, v.X) }

// Length returns the length (magnitude).
func (v Vector) Length() float64 { return math.Sqrt(v.LengthSq()) }

// Length returns the squared length (magnitude).
func (v Vector) LengthSq() float64 { return (v.X * v.X) + (v.Y * v.Y) }

// Equals compares the X and Y values of the Vector and the target Vector, and
// returns true when they are equal.
func (v Vector) Equals(target Vector) bool {
	return v.X == target.X && v.Y == target.Y
}

// Norm normalizes the Vector by making the length a magnitude of 1 in the same
// direction. It does not alter the source Vector.
func (v Vector) Normalize() Vector {
	l := v.LengthSq()
	if l == 1 {
		return v
	}
	v.Scale(1 / math.Sqrt(l))
	return v
}

// Dot calculates the dot product of the Vector and the target Vector.
func (v Vector) Dot(target Vector) float64 {
	return (v.X * target.X) + (v.Y * target.Y)
}

// Cross calculates the cross product of the Vector and the target Vector.
func (v Vector) Cross(target Vector) float64 {
	return (v.X * target.Y) + (v.Y * target.X)
}

func (v Vector) String() string {
	angle := v.Angle()
	return fmt.Sprintf("%T{%f, %f}, angle: %f (= %f°), length: %f",
		v, v.X, v.Y,
		angle, math2.RadToDeg(angle),
		v.Length(),
	)
}
