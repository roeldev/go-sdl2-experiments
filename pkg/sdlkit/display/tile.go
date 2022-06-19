package display

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit"
)

type StretchMode uint8

const (
	// Cover to fit the target's size.
	StretchFit StretchMode = iota
	// Tile inside the target's area, starting at the top left.
	StretchTile
)

// STRETCH_SCALE_ON_EXPAND = 0 --- Scale to fit the node's bounding rectangle, only if expand is true. Default stretch_mode, for backwards compatibility. Until you set expand to true, the texture will behave like STRETCH_KEEP.
// STRETCH_SCALE = 1 --- Scale to fit the node's bounding rectangle.
// STRETCH_TILE = 2 --- Tile inside the node's bounding rectangle.
// STRETCH_KEEP = 3 --- The texture keeps its original size and stays in the bounding rectangle's top-left corner.
// STRETCH_KEEP_CENTERED = 4 --- The texture keeps its original size and stays centered in the node's bounding rectangle.
// STRETCH_KEEP_ASPECT = 5 --- Scale the texture to fit the node's bounding rectangle, but maintain the texture's aspect ratio.
// STRETCH_KEEP_ASPECT_CENTERED = 6 --- Scale the texture to fit the node's bounding rectangle, center it and maintain its aspect ratio.
// STRETCH_KEEP_ASPECT_COVERED = 7 --- Scale the texture so that the shorter side fits the bounding rectangle. The other side clips to the node's limits.

type Tile struct {
	X, Y,
	W, H float64

	StretchMode StretchMode

	clip sdlkit.TextureClip
}

func NewTile(clip sdlkit.TextureClip) *Tile {
	return &Tile{
		W:    float64(clip.Location.W),
		H:    float64(clip.Location.H),
		clip: clip,
	}
}

func MustNewTile(clip sdlkit.TextureClip, possibleErr error) *Tile {
	if possibleErr != nil {
		sdlkit.FailOnErr(possibleErr)
	}

	return NewTile(clip)
}

func (s *Tile) GetX() float64  { return s.X }
func (s *Tile) GetY() float64  { return s.Y }
func (s *Tile) SetX(x float64) { s.X = x }
func (s *Tile) SetY(y float64) { s.Y = y }

func (s *Tile) Clip() sdlkit.TextureClip { return s.clip }

func (s *Tile) Draw(canvas *sdlkit.Canvas) {
	switch s.StretchMode {
	case StretchTile:
		drawStretchTile(canvas, s.clip, sdl.Rect{
			X: int32(s.X - (s.W / 2)),
			Y: int32(s.Y - (s.H / 2)),
			W: int32(s.W),
			H: int32(s.H),
		})

	case StretchFit:
		fallthrough
	default:
		canvas.DrawTextureClip(s.clip, sdl.Rect{
			X: int32(s.X - (s.W / 2)),
			Y: int32(s.Y - (s.H / 2)),
			W: int32(s.W),
			H: int32(s.H),
		})
	}
}

func drawStretchTile(canvas *sdlkit.Canvas, clip sdlkit.TextureClip, dest sdl.Rect) {
	nx := (dest.W / clip.Location.W) * clip.Location.W
	ny := (dest.H / clip.Location.H) * clip.Location.H

	// remaining
	rw := dest.W - nx
	rh := dest.H - ny

	var x, y int32
	for ; y < ny; y += clip.Location.H {
		for x = 0; x < nx; x += clip.Location.W {
			canvas.DrawTextureClip(clip, sdl.Rect{
				X: x + dest.X,
				Y: y + dest.Y,
				W: clip.Location.W,
				H: clip.Location.H,
			})
		}

		if rw != 0 {
			canvas.DrawTexture(clip.Texture,
				&sdl.Rect{
					X: clip.Location.X,
					Y: clip.Location.Y,
					W: rw,
					H: clip.Location.H,
				},
				sdl.Rect{
					X: x + dest.X,
					Y: y + dest.Y,
					W: rw,
					H: clip.Location.H,
				},
			)
		}
	}

	if rh != 0 {
		for x = 0; x < nx; x += clip.Location.W {
			canvas.DrawTexture(clip.Texture,
				&sdl.Rect{
					X: clip.Location.X,
					Y: clip.Location.Y,
					W: clip.Location.W,
					H: rh,
				},
				sdl.Rect{
					X: x + dest.X,
					Y: y + dest.Y,
					W: clip.Location.W,
					H: rh,
				},
			)
		}

		if rw != 0 {
			canvas.DrawTexture(clip.Texture,
				&sdl.Rect{
					X: clip.Location.X,
					Y: clip.Location.Y,
					W: rw,
					H: rh,
				},
				sdl.Rect{
					X: x + dest.X,
					Y: y + dest.Y,
					W: rw,
					H: rh,
				},
			)
		}
	}
}
