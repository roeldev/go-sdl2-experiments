package display

import (
	"github.com/go-pogo/errors"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit"
)

type TileMap struct {
	atlas *sdlkit.TextureAtlas
	tiles []sdl.Rect
}

func NewTileMap(atlas *sdlkit.TextureAtlas) (*TileMap, error) {
	if !atlas.IsUniform() {
		return nil, errors.Newf("display: TileMap only works with a uniform TextureAtlas")
	}

	return &TileMap{
		atlas: atlas,
	}, nil
}

func (tm *TileMap) Render(ren *sdl.Renderer) error {
	texture := tm.atlas.Texture()
	clip, _ := tm.atlas.GetFomIndex(20)
	var x, y int32
	for ; y < 6; y++ {
		for x = 0; x < 11; x++ {
			ren.Copy(texture, &clip.Location, &sdl.Rect{
				X: clip.Location.W * x,
				Y: clip.Location.H * y,
				W: clip.Location.W,
				H: clip.Location.H,
			})
		}
	}
	return nil
}
