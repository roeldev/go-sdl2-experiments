package physics

import (
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/geom"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/geom/xform"
)

type DynamicBody struct {
	geom.XY
	Collider

	// AngularVelocity is the DynamicBody's rotational velocity.
	AngularVelocity geom.Vector

	// LinearVelocity is the DynamicBody's linear velocity.
	LinearVelocity geom.Vector

	// MassDensity is the density of the DynamicBody's area and results in a
	// value for its Mass.
	MassDensity float64
	mass        float64

	transform       *xform.Transformer
	transformTarget geom.Transformable
}

// NewDynamicBody creates a new DynamicBody using the provided geom.Shape and
// Collider. When a nil value is provided for collider, a Collider is created
// using NewCollider and the provided geom.Shape.
func NewDynamicBody(shape geom.Shape, collider Collider) *DynamicBody {
	if collider == nil {
		collider = NewCollider(shape)
	}

	body := &DynamicBody{
		XY:          shape,
		Collider:    collider,
		MassDensity: 1,
	}

	if t, ok := shape.(geom.Transformable); ok {
		body.transform = xform.NewTransformer()
		body.transformTarget = t
	}

	// shape.UsePosition(&body.position)
	// if collider.Shape() != shape {
	// 	collider.Shape().UsePosition(&body.position)
	// }

	return body
}

func (b *DynamicBody) Transformer() *xform.Transformer { return b.transform }

func (b *DynamicBody) Mass() float64 {
	return b.Collider.Shape().Area() * b.MassDensity
}

func (b *DynamicBody) Update() {
	if b.transform.HasChanged() && b.transformTarget != nil {
		b.transformTarget.Transform(b.transform.Matrix())
	}
}
