package geom

type Shape interface {
	Position() *Point
	// HitTest(shape Shape)
	// HitTestPoint(point Point)
}

type Ellipse struct {
	X, Y             float64
	RadiusX, RadiusY float64
}

type Circle struct {
	X, Y, Radius float64
}

// The difference between a Rectangle and a Quad (Polygon) is that rectangle
// is a quadrilateral, having opposing sides parallel and four right angles,
// while a Quad is a four-sided polygon, having four angles and four straight
// sides. A Rectangle cannot be transformed using Transform but can change position and size.
type Rectangle struct {
	X, Y, W, H float64
}
