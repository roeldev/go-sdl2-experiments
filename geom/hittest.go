package geom

// InEllipse returns true when x/y is inside the ellipse, defined by the
// provided ex, ey, radx and rady values.
func InEllipse(x, y, ex, ey, radx, rady float64) bool {
	return false
}

// InCircle returns true when x/y is inside the circle, defined by the provided
// cx, cy and rad values. It calculates the squared distance between the x/y
// and cx/cy values and compares this with the squared rad value.
func InCircle(x, y, cx, cy, rad float64) bool {
	dx, dy := cx-x, cy-y
	return ((dx * dx) + (dy * dy)) < (rad * rad)
}

func InRect(x, y, rx, ry, rw, rh float64) bool {
	return (x >= rx) && (x < rx+rw) &&
		(y >= ry) && (y < ry+rh)
}

func InPolygon(x, y float64, edges []Point) bool {
	return InConvexPolygon(x, y, edges) || InConcavePolygon(x, y, edges)
}

// https://en.wikipedia.org/wiki/Convex_polygon
func InConvexPolygon(x, y float64, edges []Point) bool {
	return false
}

// https://en.wikipedia.org/wiki/Concave_polygon
func InConcavePolygon(x, y float64, edges []Point) bool {
	return false
}
