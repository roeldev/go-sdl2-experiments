package geom

import (
	"math"
)

const (
	r2dPi = 180 / math.Pi
	d2rPi = math.Pi / 180
)

func RadToDeg(rad float64) float64 { return rad * r2dPi }

func DegToRad(deg float64) float64 { return deg * d2rPi }
