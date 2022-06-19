package sdlkit

import (
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/geom"
)

// Convert world positions to screen positions
// https://gamedev.stackexchange.com/questions/121421/how-to-use-the-sdl-viewport-properly
// https://www.youtube.com/watch?v=D1dw7L0nC6s
type Camera struct {
	i32 [4]int32   // x,y,w,h
	f64 [4]float64 // x,y,w,h

	targetX  float64
	targetY  float64
	disabled bool
}

func NewCamera(x, y float64, w, h int32) *Camera {
	var cam Camera
	cam.SetX(x)
	cam.SetY(y)
	cam.Resize(w, h)
	return &cam
}

func (c *Camera) Enable()  { c.disabled = false }
func (c *Camera) Disable() { c.disabled = true }

func (c *Camera) IsEnabled() bool {
	return c != nil && !c.disabled
}

func (c *Camera) Width() int32  { return c.i32[2] }
func (c *Camera) Height() int32 { return c.i32[3] }

func (c *Camera) GetX() float64 { return c.f64[0] }
func (c *Camera) GetY() float64 { return c.f64[1] }

func (c *Camera) SetX(x float64) {
	if x < 0 {
		x = 0
	}
	c.i32[0], c.f64[0] = int32(x), x
}

func (c *Camera) SetY(y float64) {
	if y < 0 {
		y = 0
	}
	c.i32[1], c.f64[1] = int32(y), y
}

func (c *Camera) Follow(target geom.XYGetter) {
	tx, ty := target.GetX(), target.GetY()
	if tx-c.targetX > 0.5 || tx-c.targetX < -0.5 {
		c.SetX(tx - (c.f64[2] / 2))
		c.targetX = tx
	}
	if ty-c.targetY > 0.5 || ty-c.targetY < -0.5 {
		c.SetY(ty - (c.f64[3] / 2))
		c.targetY = ty
	}
}

func (c *Camera) Resize(w, h int32) {
	c.i32[2], c.f64[2] = w, float64(w)
	c.i32[3], c.f64[3] = h, float64(h)
}

func (c *Camera) TranslateX(x int32) int32 {
	return x - c.i32[0]
}

func (c *Camera) TranslateXF(x float64) int32 {
	return int32(x - c.f64[0])
}

func (c *Camera) TranslateY(y int32) int32 {
	return y - c.i32[1]
}

func (c *Camera) TranslateYF(y float64) int32 {
	return int32(y - c.f64[1])
}

// - camera + viewport toevoegen
// - ook ZOOM kunnen regelen
// - alle Draw() methods aanpassen zodat de nieuwe renderer gebruikt word
// - camera shake
