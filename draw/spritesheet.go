package draw

import (
	"github.com/go-pogo/errors"
	sdlimg "github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func LoadSprites(ren *sdl.Renderer, file string) {

}

func LoadUniformSprites(ren *sdl.Renderer, file string, cellW, cellH int32, total uint8) (*SpriteSheet, error) {
	tx, err := sdlimg.LoadTexture(ren, file)
	if err != nil {
		return nil, err
	}

	return parseUniformSprites(tx, cellW, cellH, total)
}

func LoadUniformSpritesRW(ren *sdl.Renderer, src *sdl.RWops, freeSrc bool, cellW, cellH int32, total uint8) (*SpriteSheet, error) {
	tx, err := sdlimg.LoadTextureRW(ren, src, freeSrc)
	if err != nil {
		return nil, err
	}

	return parseUniformSprites(tx, cellW, cellH, total)
}

func LoadUniformSpritesFromMem(ren *sdl.Renderer, data []byte, cellW, cellH int32, total uint8) (*SpriteSheet, error) {
	src, err := sdl.RWFromMem(data)
	if err != nil {
		return nil, err
	}

	return LoadUniformSpritesRW(ren, src, true, cellW, cellH, total)
}

// SpriteSheet word interface
// UniformSpriteSheet iedere cell vaste breedte/hoogte
// MixedSpriteSheet iedere cell verschillende dimensie

type SpriteSheet struct {
	tx      *sdl.Texture
	indexes []sdl.Rect
	names   map[string]sdl.Rect
}

// func x() {
// 	sheet := NewSpriteSheet("file.png")
// 	sheet.UniformCells("name_prefix", 128, 128, 40)
// 	sheet.LoadAtlasXML("file.xml")
// 	sheet.Parse()- > error
// 	sheet.MustParse()- > panic
// }

func parseUniformSprites(tx *sdl.Texture, cellW, cellH int32, total uint8) (*SpriteSheet, error) {
	_, _, txW, txH, err := tx.Query()
	if err != nil {
		return nil, err
	}

	var x, y int32
	cells := make([]sdl.Rect, 0, total)

	if total < 1 {
		return nil, errors.Newf("sdlkit: a SpriteSheet needs at least 1 cell")
	}

Loop:
	for y = 0; y < txH; y += cellH {
		for x = 0; x < txW; x += cellW {
			cells = append(cells, sdl.Rect{X: x, Y: y, W: cellW, H: cellH})

			total--
			if total <= 0 {
				break Loop
			}
		}
	}

	return &SpriteSheet{
		tx:      tx,
		indexes: cells,
	}, nil
}

// func NewSpriteSheetFromSurface(surface sdl.Surface) {
//
// }
//
// func NewSpriteSheetFromFile(filename string) {
//
// }
//
// func NewSpriteSheetFromMem(data []byte) {
//
// }

func (s *SpriteSheet) Texture() *sdl.Texture     { return s.tx }
func (s *SpriteSheet) Index(n int) sdl.Rect      { return s.indexes[n] }
func (s *SpriteSheet) Find(name string) sdl.Rect { return s.names[name] }
func (s *SpriteSheet) Len() int                  { return len(s.indexes) }

func (s *SpriteSheet) AllByIndex() []sdl.Rect         { return s.indexes }
func (s *SpriteSheet) AllByName() map[string]sdl.Rect { return s.names }
