package physics

import (
	"math"

	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/geom"
)

// StaticBody is a Body that is not intended to move. It is ideal for
// implementing static entities in the environment, such as walls, trees or
// platforms.
type StaticBody struct {
	geom.XY
	Collider
}

// NewStaticBody creates a new StaticBody using the provided geom.Shape and
// Collider. When a nil value is provided for collider, a Collider is created
// using NewCollider and the provided geom.Shape.
func NewStaticBody(shape geom.Shape, collider Collider) *StaticBody {
	if collider == nil {
		collider = NewCollider(shape)
	}

	return &StaticBody{
		XY:       shape,
		Collider: collider,
	}
}

var infiniteMass = math.Inf(1)

func (sb StaticBody) Mass() float64 { return infiniteMass }
