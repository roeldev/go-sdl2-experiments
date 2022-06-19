package draw

import (
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/geom"
)

// func (c *Renderer) DrawShape(shape geom.Shape) {
// 	switch s := shape.(type) {
// 	case *geom.Ellipse:
// 		c.DrawEllipseShape(*s)
// 	case *geom.Circle:
// 		c.DrawCircleShape(*s)
// 	case *geom.Rect:
// 		c.DrawRectShape(*s)
// 	case geom.PolygonShape:
// 		c.DrawPolygonShape(s)
// 	}
// }

func Point(xy geom.XYGetter, size int32) sdlkit.DrawableFunc {
	return func(canvas *sdlkit.Canvas) {
		canvas.DrawPixel(int32(xy.GetX()), int32(xy.GetY()), size)
	}
}

func Circle(cs geom.Circle) sdlkit.DrawableFunc {
	return func(canvas *sdlkit.Canvas) {
		canvas.DrawEllipse(
			int32(cs.X),
			int32(cs.Y),
			int32(cs.Radius-0.5),
			int32(cs.Radius-0.5),
		)
	}
}

func Ellipse(es geom.Ellipse) sdlkit.DrawableFunc {
	return func(canvas *sdlkit.Canvas) {
		canvas.DrawEllipse(
			int32(es.X),
			int32(es.Y),
			int32(es.RadiusX-0.5),
			int32(es.RadiusY-0.5),
		)
	}
}

func Rect(rs geom.Rect) sdlkit.DrawableFunc {
	return func(canvas *sdlkit.Canvas) {
		canvas.DrawSdlRect(*rs.Rect())
	}
}

func Polygon(ps geom.PolygonShape) sdlkit.DrawableFunc {
	return func(canvas *sdlkit.Canvas) {
		var tx, ty int16

		camera := canvas.Camera()
		if camera.IsEnabled() {
			tx = int16(camera.TranslateX(0))
			ty = int16(camera.TranslateY(0))
		}

		vertices := ps.Vertices()
		vx := make([]int16, len(vertices))
		vy := make([]int16, len(vertices))

		for i, pt := range vertices {
			vx[i] = int16(pt.X-1) + tx
			vy[i] = int16(pt.Y-1) + ty
		}

		canvas.Camera().Disable()
		canvas.DrawPolygon(vx, vy)
		canvas.Camera().Enable()
	}
}
