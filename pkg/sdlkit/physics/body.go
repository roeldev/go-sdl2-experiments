package physics

import (
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/geom"
)

type Body interface {
	geom.XY
	Collider

	Shape() geom.Shape
	Mass() float64

	// 	Shape
	// geom.Shape
	// Transform geom.Transform
	// Static    bool // when true, the Body cannot change position or transform
}

// NewBody creates a new Body from the provided components.
// Default MassComponent, ColliderComponent are added for known shapes when no
// custom component of that type is provided.
