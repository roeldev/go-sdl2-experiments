package sdlkit

import (
	"github.com/go-pogo/errors"
	"github.com/veandco/go-sdl2/sdl"
)

type TextureAtlas struct {
	tx      *sdl.Texture
	sources []sdl.Rect
	names   map[string]int
	uniform bool
}

func NewTextureAtlas(tx *sdl.Texture, uniform bool) *TextureAtlas {
	return &TextureAtlas{
		tx:      tx,
		sources: make([]sdl.Rect, 0, 32),
		names:   make(map[string]int, 32),
		uniform: uniform,
	}
}

func NewTextureAtlasU(tx *sdl.Texture, cellW, cellH int32, total uint8) (*TextureAtlas, error) {
	if total < 1 {
		return nil, errors.Newf("sdlkit: a SpriteSheet needs at least 1 cell")
	}

	_, _, txW, txH, err := tx.Query()
	if err != nil {
		return nil, err
	}

	var x, y int32
	ta := &TextureAtlas{
		tx:      tx,
		sources: make([]sdl.Rect, 0, total),
		names:   make(map[string]int),
		uniform: true,
	}

Loop:
	for y = 0; y < txH; y += cellH {
		for x = 0; x < txW; x += cellW {
			ta.sources = append(ta.sources, sdl.Rect{X: x, Y: y, W: cellW, H: cellH})

			total--
			if total <= 0 {
				break Loop
			}
		}
	}

	return ta, nil
}

func (ta *TextureAtlas) Texture() *sdl.Texture { return ta.tx }

func (ta *TextureAtlas) Add(name string, src sdl.Rect) {
	if name != "" {
		if i, exists := ta.names[name]; exists {
			ta.sources[i] = src
			return
		}

		ta.names[name] = len(ta.sources)
	}

	ta.sources = append(ta.sources, src)
}

func (ta *TextureAtlas) Index(i int) (sdl.Rect, bool) {
	if len(ta.sources) <= i {
		return sdl.Rect{}, false
	}

	return ta.sources[i], true
}

func (ta *TextureAtlas) Name(name string) (sdl.Rect, bool) {
	if i, ok := ta.names[name]; ok {
		return ta.sources[i], true
	}

	return sdl.Rect{}, false
}
